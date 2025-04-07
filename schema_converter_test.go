/*
 * Copyright (C) 2025 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGCPCredentials(t *testing.T) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/path/to/credentials.json")
	defer os.Unsetenv("GOOGLE_APPLICATION_CREDENTIALS")

	t.Run("Credentials Set", func(t *testing.T) {
		if err := checkGCPCredentials(); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Credentials Not Set", func(t *testing.T) {
		os.Unsetenv("GOOGLE_APPLICATION_CREDENTIALS")
		if err := checkGCPCredentials(); err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestStmtExtractor(t *testing.T) {
	testCases := []struct {
		name             string
		content          string
		expectedStmts    []string
		expectError      bool
		expectedErrorMsg string
	}{
		{
			name: "Basic",
			content: "CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			expectedStmts: []string{
				"CREATE TABLE test (id INT PRIMARY KEY);",
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			},
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "No semicolon in the end of the last stmt",
			content: "CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"CREATE TABLE another_test (id INT PRIMARY KEY)",
			expectedStmts: []string{
				"CREATE TABLE test (id INT PRIMARY KEY);",
				"CREATE TABLE another_test (id INT PRIMARY KEY)",
			},
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Line comment at the beginning",
			content: "CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"// CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"-- CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			expectedStmts: []string{
				"CREATE TABLE test (id INT PRIMARY KEY);",
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			},
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Line comment at the end",
			content: "CREATE TABLE test (id INT PRIMARY KEY); // test 123\n" +
				"// CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"-- CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"CREATE TABLE another_test (id INT PRIMARY KEY); -- comment",
			expectedStmts: []string{
				"CREATE TABLE test (id INT PRIMARY KEY);",
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			},
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Block comment",
			content: "CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"/* \n" +
				"CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"CREATE TABLE test (id INT PRIMARY KEY);*/" +
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			expectedStmts: []string{
				"CREATE TABLE test (id INT PRIMARY KEY);",
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			},
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Unterminated Block comment",
			content: "CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"/* \n" +
				"CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			expectedStmts: []string{
				"CREATE TABLE test (id INT PRIMARY KEY);",
			},
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Inline block comment",
			content: "CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"/*CREATE TABLE test (id INT PRIMARY KEY);*/ CREATE TABLE test2 (id2 INT PRIMARY KEY);\n" +
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			expectedStmts: []string{
				"CREATE TABLE test (id INT PRIMARY KEY);",
				"CREATE TABLE test2 (id2 INT PRIMARY KEY);",
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			},
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Two block comments in the same line",
			content: "CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"/*CREATE TABLE test (id INT PRIMARY KEY);*/ CREATE TABLE test2 (id2 INT PRIMARY KEY); /* test */\n" +
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			expectedStmts: []string{
				"CREATE TABLE test (id INT PRIMARY KEY);",
				"CREATE TABLE test2 (id2 INT PRIMARY KEY);",
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			},
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Double-slash in block comment",
			content: "CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"/* test // */CREATE TABLE test2 (id INT PRIMARY KEY);\n" +
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			expectedStmts: []string{
				"CREATE TABLE test (id INT PRIMARY KEY);",
				"CREATE TABLE test2 (id INT PRIMARY KEY);",
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			},
			expectError:      false,
			expectedErrorMsg: "",
		},
		{
			name: "Double-slash after block comment",
			content: "CREATE TABLE test (id INT PRIMARY KEY);\n" +
				"/* test */ //CREATE TABLE test2 (id INT PRIMARY KEY);\n" +
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			expectedStmts: []string{
				"CREATE TABLE test (id INT PRIMARY KEY);",
				"CREATE TABLE another_test (id INT PRIMARY KEY);",
			},
			expectError:      false,
			expectedErrorMsg: "",
		},
	}

	// TODO: Use os.CreateTemp to create the test file.
	filepath := "test.cql"
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Write the content to a temporary file
			if err := os.WriteFile(filepath, []byte(tc.content), 0644); err != nil {
				t.Fatalf("Failed to write test file: %v", err)
			}
			defer os.Remove(filepath)

			stmts, err := parseCqlFile(filepath)
			if tc.expectError {
				assert.Equal(t, tc.expectedErrorMsg, err.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStmts, stmts)
		})
	}
}
