// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
	"github.com/cloudspannerecosystem/spanner-cassandra-schema-tool/translator"
)

// Flags holds the command-line flags for configuring the Cassandra to Spanner schema conversion.
type Flags struct {
	projectID  string
	instanceID string
	databaseID string
	cqlFile    string
	dryRun     bool
	// TODO: add a flag for creating the database if it does not exist.
}

// parseFlags parses the command-line flags and populates the Flags struct.
func (f *Flags) parseFlags() {
	flag.StringVar(&f.projectID, "project", "", "The Google Cloud project ID")
	flag.StringVar(&f.instanceID, "instance", "", "The Spanner instance ID")
	flag.StringVar(&f.databaseID, "database", "", "The Spanner database ID")
	flag.StringVar(&f.cqlFile, "cql", "", "The path of the CQL file containing DDL statements to be converted and executed on the Spanner database")
	flag.BoolVar(&f.dryRun, "dry-run", false, "Only output converted Spanner DDL statements without execution")

	flag.Parse()
}

// validateRequiredFlags checks if all required command-line flags are provided.
// It returns true if all required flags are present, and false otherwise.
// If false, it also prints usage instructions to the console.
func (f *Flags) validateRequiredFlags() bool {
	missingFlags := []string{}
	if f.projectID == "" {
		missingFlags = append(missingFlags, "-project")
	}
	if f.instanceID == "" {
		missingFlags = append(missingFlags, "-instance")
	}
	if f.databaseID == "" {
		missingFlags = append(missingFlags, "-database")
	}
	if f.cqlFile == "" {
		missingFlags = append(missingFlags, "-cql")
	}

	if len(missingFlags) > 0 {
		fmt.Println("Missing required flags:", missingFlags)
		flag.PrintDefaults()
		return false
	}
	return true
}

// checkGCPCredentials checks if the GCP credentials are set in the environment.
func checkGCPCredentials() error {
	credentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credentials == "" {
		return fmt.Errorf("GCP credentials are not found. Set GOOGLE_APPLICATION_CREDENTIALS environment variable")
	}
	return nil
}

// extractQueries parses a CQL file, splitting it into individual queries delimited by semicolons.
//
// TODO: Use the Antlr parser to extract the statements rather than using this function.
func extractQueries(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var queries []string
	var currentQuery strings.Builder
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "--") || strings.HasPrefix(line, "//") {
			continue
		}

		currentQuery.WriteString(line + " ")

		// If the line ends with a semicolon, treat it as a full query
		if strings.HasSuffix(line, ";") {
			queries = append(queries, currentQuery.String())
			currentQuery.Reset()
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return queries, nil
}

func main() {
	var flags Flags
	flags.parseFlags()
	if !flags.validateRequiredFlags() {
		os.Exit(1)
	}

	if err := checkGCPCredentials(); err != nil {
		log.Fatalf("Error: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create admin client: %v", err)
	}
	defer adminClient.Close()

	// Translates the Cassandra DDL stmts in the CQL file to corresponding Spanner Stmts.
	fmt.Printf("Starting Cassandra to Spanner conversion for '%s'\n\n\n", flags.cqlFile)
	var spannerCreateTableStmts []string
	queries, err := extractQueries(flags.cqlFile)
	if err != nil {
		log.Fatalf("Failed to read the file: %v\n", err)
	}
	for _, query := range queries {
		fmt.Printf("Converting statement: '%s'\n", query)
		spannerCreateTableStmt, err := translator.ToSpannerCreateTableStmt(query, flags.databaseID)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		fmt.Print("-> SUCCESS\n\n")
		spannerCreateTableStmts = append(spannerCreateTableStmts, spannerCreateTableStmt)
	}
	// TODO: Add a summary block.

	fmt.Print("------------ Generated Spanner DDL ------------\n\n")
	for _, value := range spannerCreateTableStmts {
		fmt.Printf("%s\n\n", value)
	}
	fmt.Println("------------ End of Spanner DDL ------------")

	if flags.dryRun {
		fmt.Println("\n*** DRY RUN: The DDL above shows the expected output but was NOT executed. ***")
		return
	}

	// Execute the translated Spanner DDL statements.
	fmt.Println("\nExecuting generated DDL on Spanner...")
	op, err := adminClient.UpdateDatabaseDdl(ctx, &databasepb.UpdateDatabaseDdlRequest{
		Database:   fmt.Sprintf("projects/%s/instances/%s/databases/%s", flags.projectID, flags.instanceID, flags.databaseID),
		Statements: spannerCreateTableStmts,
	})
	if err != nil {
		log.Fatalf("Failed to execute DDL statements on the Cloud Spanner database: %v", err)
	}
	if err := op.Wait(ctx); err != nil {
		log.Fatalf("Failed to execute DDL statements on the Cloud Spanner database: %v", err)
	}
	fmt.Println("Spanner schema update complete!")
}
