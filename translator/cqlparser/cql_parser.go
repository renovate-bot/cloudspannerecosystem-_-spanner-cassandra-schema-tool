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

// Code generated from CqlParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CqlParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CqlParser struct {
	*antlr.BaseParser
}

var CqlParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cqlparserParserInit() {
	staticData := &CqlParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'ADD'", "'AGGREGATE'", "'ALL'", "'ALLOW'", "'ALTER'", "'AND'",
		"'APPLY'", "'AS'", "'ASC'", "'ASCII'", "'AUTHORIZE'", "'BATCH'", "'BEGIN'",
		"'BIGINT'", "'BLOB'", "'BOOLEAN'", "'BY'", "'CALLED'", "'CLUSTERING'",
		"'COLUMNFAMILY'", "'COMPACT'", "'CONTAINS'", "'COUNT'", "'COUNTER'",
		"'CREATE'", "'CUSTOM'", "'DATE'", "'DECIMAL'", "'DELETE'", "'DESC'",
		"'DESCRIBE'", "'DISTINCT'", "'DOUBLE'", "'DROP'", "'ENTRIES'", "'EXECUTE'",
		"'EXISTS'", "'FILTERING'", "'FINALFUNC'", "'FLOAT'", "'FROM'", "'FROZEN'",
		"'FULL'", "'FUNCTION'", "'FUNCTIONS'", "'GRANT'", "'IF'", "'IN'", "'INDEX'",
		"'INET'", "'INFINITY'", "'INITCOND'", "'INPUT'", "'INSERT'", "'INT'",
		"'INTO'", "'JSON'", "'KEY'", "'KEYS'", "'KEYSPACE'", "'KEYSPACES'",
		"'LANGUAGE'", "'LIMIT'", "'LIST'", "'LOGIN'", "'MAP'", "'MODIFY'", "'NAN'",
		"'NOLOGIN'", "'NORECURSIVE'", "'NOSUPERUSER'", "'NOT'", "'NULL'", "'OF'",
		"'ON'", "'OPTIONS'", "'OR'", "'ORDER'", "'PASSWORD'", "'PERMISSION'",
		"'PERMISSIONS'", "'PRIMARY'", "'RENAME'", "'REPLACE'", "'RETURNS'",
		"'REVOKE'", "'ROLE'", "'ROLES'", "'SCHEMA'", "'SELECT'", "'SET'", "'SFUNC'",
		"'SMALLINT'", "'STATIC'", "'STORAGE'", "'STYPE'", "'SUPERUSER'", "'TABLE'",
		"'TEXT'", "'TIME'", "'TIMESTAMP'", "'TIMEUUID'", "'TINYINT'", "'TO'",
		"'TOKEN'", "'TRIGGER'", "'TRUNCATE'", "'TTL'", "'TUPLE'", "'TYPE'",
		"'UNLOGGED'", "'UPDATE'", "'USE'", "'USER'", "'USERS'", "'USING'", "'UUID'",
		"'VALUES'", "'VARCHAR'", "'VARINT'", "'WHERE'", "'WITH'", "'WRITETIME'",
		"';'", "'\"'", "'.'", "','", "'('", "')'", "'<'", "'>'",
	}
	staticData.SymbolicNames = []string{
		"", "K_ADD", "K_AGGREGATE", "K_ALL", "K_ALLOW", "K_ALTER", "K_AND",
		"K_APPLY", "K_AS", "K_ASC", "K_ASCII", "K_AUTHORIZE", "K_BATCH", "K_BEGIN",
		"K_BIGINT", "K_BLOB", "K_BOOLEAN", "K_BY", "K_CALLED", "K_CLUSTERING",
		"K_COLUMNFAMILY", "K_COMPACT", "K_CONTAINS", "K_COUNT", "K_COUNTER",
		"K_CREATE", "K_CUSTOM", "K_DATE", "K_DECIMAL", "K_DELETE", "K_DESC",
		"K_DESCRIBE", "K_DISTINCT", "K_DOUBLE", "K_DROP", "K_ENTRIES", "K_EXECUTE",
		"K_EXISTS", "K_FILTERING", "K_FINALFUNC", "K_FLOAT", "K_FROM", "K_FROZEN",
		"K_FULL", "K_FUNCTION", "K_FUNCTIONS", "K_GRANT", "K_IF", "K_IN", "K_INDEX",
		"K_INET", "K_INFINITY", "K_INITCOND", "K_INPUT", "K_INSERT", "K_INT",
		"K_INTO", "K_JSON", "K_KEY", "K_KEYS", "K_KEYSPACE", "K_KEYSPACES",
		"K_LANGUAGE", "K_LIMIT", "K_LIST", "K_LOGIN", "K_MAP", "K_MODIFY", "K_NAN",
		"K_NOLOGIN", "K_NORECURSIVE", "K_NOSUPERUSER", "K_NOT", "K_NULL", "K_OF",
		"K_ON", "K_OPTIONS", "K_OR", "K_ORDER", "K_PASSWORD", "K_PERMISSION",
		"K_PERMISSIONS", "K_PRIMARY", "K_RENAME", "K_REPLACE", "K_RETURNS",
		"K_REVOKE", "K_ROLE", "K_ROLES", "K_SCHEMA", "K_SELECT", "K_SET", "K_SFUNC",
		"K_SMALLINT", "K_STATIC", "K_STORAGE", "K_STYPE", "K_SUPERUSER", "K_TABLE",
		"K_TEXT", "K_TIME", "K_TIMESTAMP", "K_TIMEUUID", "K_TINYINT", "K_TO",
		"K_TOKEN", "K_TRIGGER", "K_TRUNCATE", "K_TTL", "K_TUPLE", "K_TYPE",
		"K_UNLOGGED", "K_UPDATE", "K_USE", "K_USER", "K_USERS", "K_USING", "K_UUID",
		"K_VALUES", "K_VARCHAR", "K_VARINT", "K_WHERE", "K_WITH", "K_WRITETIME",
		"SEMICOLON", "DQUOTE", "DOT", "COMMA", "L_PAREN", "R_PAREN", "L_ANGLE_BRACKET",
		"R_ANGLE_BRACKET", "IDENTIFIER", "IDENTIFIER_WITH_HYPHEN", "WS", "COMMENT",
		"MULTILINE_COMMENT", "UNKNOWN",
	}
	staticData.RuleNames = []string{
		"root", "cqlStatement", "createTable", "columnDefinitionList", "columnDefinition",
		"primaryKeyClause", "primaryKey", "partitionKey", "clusteringColumns",
		"columnName", "tableName", "keyspaceName", "generalIdentifier", "identifier",
		"cqlType", "primaryKeyKeywords", "ifNotExist", "cqlNativeType", "cqlCollectionType",
		"wihtTableOptions", "nonSemicolonToken", "keyword", "reservedKeyword",
		"nonReservedKeyword",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 137, 192, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 3, 0, 51, 8, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 3, 2, 60, 8, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 3, 2, 67, 8, 2, 1, 3, 1, 3, 1, 3, 5, 3, 72, 8, 3, 10, 3, 12, 3,
		75, 9, 3, 1, 3, 1, 3, 3, 3, 79, 8, 3, 1, 4, 1, 4, 1, 4, 3, 4, 84, 8, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 3, 6, 94, 8, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 5, 7, 101, 8, 7, 10, 7, 12, 7, 104, 9, 7, 1, 7, 1,
		7, 3, 7, 108, 8, 7, 1, 8, 1, 8, 1, 8, 5, 8, 113, 8, 8, 10, 8, 12, 8, 116,
		9, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10, 123, 8, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 12, 1, 12, 3, 12, 131, 8, 12, 1, 13, 1, 13, 3, 13, 135,
		8, 13, 1, 14, 1, 14, 3, 14, 139, 8, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 162, 8, 18, 1, 19, 1,
		19, 5, 19, 166, 8, 19, 10, 19, 12, 19, 169, 9, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 182, 8,
		20, 1, 21, 1, 21, 3, 21, 186, 8, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		0, 0, 24, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 0, 4, 12, 0, 10, 10, 14, 16, 24, 24, 27, 28,
		33, 33, 40, 40, 50, 50, 55, 55, 93, 93, 99, 103, 117, 117, 119, 120, 2,
		0, 64, 64, 91, 91, 30, 0, 1, 1, 4, 7, 9, 9, 11, 13, 17, 17, 20, 20, 25,
		25, 29, 31, 34, 36, 41, 41, 43, 43, 46, 49, 51, 51, 54, 54, 56, 56, 60,
		60, 63, 63, 67, 68, 70, 70, 72, 75, 77, 78, 82, 84, 86, 86, 89, 91, 98,
		98, 104, 105, 107, 107, 111, 113, 116, 116, 121, 122, 30, 0, 2, 3, 8, 8,
		10, 10, 14, 16, 18, 19, 21, 24, 26, 28, 32, 33, 37, 40, 42, 42, 44, 45,
		50, 50, 52, 53, 55, 55, 57, 59, 61, 62, 64, 66, 69, 69, 71, 71, 76, 76,
		79, 81, 85, 85, 87, 88, 92, 97, 99, 103, 106, 106, 108, 110, 114, 115,
		117, 120, 123, 123, 194, 0, 48, 1, 0, 0, 0, 2, 54, 1, 0, 0, 0, 4, 56, 1,
		0, 0, 0, 6, 68, 1, 0, 0, 0, 8, 80, 1, 0, 0, 0, 10, 85, 1, 0, 0, 0, 12,
		90, 1, 0, 0, 0, 14, 107, 1, 0, 0, 0, 16, 109, 1, 0, 0, 0, 18, 117, 1, 0,
		0, 0, 20, 122, 1, 0, 0, 0, 22, 126, 1, 0, 0, 0, 24, 130, 1, 0, 0, 0, 26,
		134, 1, 0, 0, 0, 28, 138, 1, 0, 0, 0, 30, 140, 1, 0, 0, 0, 32, 143, 1,
		0, 0, 0, 34, 147, 1, 0, 0, 0, 36, 161, 1, 0, 0, 0, 38, 163, 1, 0, 0, 0,
		40, 181, 1, 0, 0, 0, 42, 185, 1, 0, 0, 0, 44, 187, 1, 0, 0, 0, 46, 189,
		1, 0, 0, 0, 48, 50, 3, 2, 1, 0, 49, 51, 5, 124, 0, 0, 50, 49, 1, 0, 0,
		0, 50, 51, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 5, 0, 0, 1, 53, 1, 1,
		0, 0, 0, 54, 55, 3, 4, 2, 0, 55, 3, 1, 0, 0, 0, 56, 57, 5, 25, 0, 0, 57,
		59, 5, 98, 0, 0, 58, 60, 3, 32, 16, 0, 59, 58, 1, 0, 0, 0, 59, 60, 1, 0,
		0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 3, 20, 10, 0, 62, 63, 5, 128, 0, 0, 63,
		64, 3, 6, 3, 0, 64, 66, 5, 129, 0, 0, 65, 67, 3, 38, 19, 0, 66, 65, 1,
		0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 5, 1, 0, 0, 0, 68, 73, 3, 8, 4, 0, 69,
		70, 5, 127, 0, 0, 70, 72, 3, 8, 4, 0, 71, 69, 1, 0, 0, 0, 72, 75, 1, 0,
		0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 78, 1, 0, 0, 0, 75, 73,
		1, 0, 0, 0, 76, 77, 5, 127, 0, 0, 77, 79, 3, 10, 5, 0, 78, 76, 1, 0, 0,
		0, 78, 79, 1, 0, 0, 0, 79, 7, 1, 0, 0, 0, 80, 81, 3, 18, 9, 0, 81, 83,
		3, 28, 14, 0, 82, 84, 3, 30, 15, 0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0, 0,
		0, 84, 9, 1, 0, 0, 0, 85, 86, 3, 30, 15, 0, 86, 87, 5, 128, 0, 0, 87, 88,
		3, 12, 6, 0, 88, 89, 5, 129, 0, 0, 89, 11, 1, 0, 0, 0, 90, 93, 3, 14, 7,
		0, 91, 92, 5, 127, 0, 0, 92, 94, 3, 16, 8, 0, 93, 91, 1, 0, 0, 0, 93, 94,
		1, 0, 0, 0, 94, 13, 1, 0, 0, 0, 95, 108, 3, 18, 9, 0, 96, 97, 5, 128, 0,
		0, 97, 102, 3, 18, 9, 0, 98, 99, 5, 127, 0, 0, 99, 101, 3, 18, 9, 0, 100,
		98, 1, 0, 0, 0, 101, 104, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1,
		0, 0, 0, 103, 105, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 105, 106, 5, 129,
		0, 0, 106, 108, 1, 0, 0, 0, 107, 95, 1, 0, 0, 0, 107, 96, 1, 0, 0, 0, 108,
		15, 1, 0, 0, 0, 109, 114, 3, 18, 9, 0, 110, 111, 5, 127, 0, 0, 111, 113,
		3, 18, 9, 0, 112, 110, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0,
		0, 0, 114, 115, 1, 0, 0, 0, 115, 17, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0,
		117, 118, 3, 26, 13, 0, 118, 19, 1, 0, 0, 0, 119, 120, 3, 22, 11, 0, 120,
		121, 5, 126, 0, 0, 121, 123, 1, 0, 0, 0, 122, 119, 1, 0, 0, 0, 122, 123,
		1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 3, 26, 13, 0, 125, 21, 1, 0,
		0, 0, 126, 127, 3, 24, 12, 0, 127, 23, 1, 0, 0, 0, 128, 131, 5, 133, 0,
		0, 129, 131, 3, 26, 13, 0, 130, 128, 1, 0, 0, 0, 130, 129, 1, 0, 0, 0,
		131, 25, 1, 0, 0, 0, 132, 135, 3, 46, 23, 0, 133, 135, 5, 132, 0, 0, 134,
		132, 1, 0, 0, 0, 134, 133, 1, 0, 0, 0, 135, 27, 1, 0, 0, 0, 136, 139, 3,
		34, 17, 0, 137, 139, 3, 36, 18, 0, 138, 136, 1, 0, 0, 0, 138, 137, 1, 0,
		0, 0, 139, 29, 1, 0, 0, 0, 140, 141, 5, 82, 0, 0, 141, 142, 5, 58, 0, 0,
		142, 31, 1, 0, 0, 0, 143, 144, 5, 47, 0, 0, 144, 145, 5, 72, 0, 0, 145,
		146, 5, 37, 0, 0, 146, 33, 1, 0, 0, 0, 147, 148, 7, 0, 0, 0, 148, 35, 1,
		0, 0, 0, 149, 150, 5, 66, 0, 0, 150, 151, 5, 130, 0, 0, 151, 152, 3, 34,
		17, 0, 152, 153, 5, 127, 0, 0, 153, 154, 3, 34, 17, 0, 154, 155, 5, 131,
		0, 0, 155, 162, 1, 0, 0, 0, 156, 157, 7, 1, 0, 0, 157, 158, 5, 130, 0,
		0, 158, 159, 3, 34, 17, 0, 159, 160, 5, 131, 0, 0, 160, 162, 1, 0, 0, 0,
		161, 149, 1, 0, 0, 0, 161, 156, 1, 0, 0, 0, 162, 37, 1, 0, 0, 0, 163, 167,
		5, 122, 0, 0, 164, 166, 3, 40, 20, 0, 165, 164, 1, 0, 0, 0, 166, 169, 1,
		0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 39, 1, 0, 0,
		0, 169, 167, 1, 0, 0, 0, 170, 182, 3, 42, 21, 0, 171, 182, 5, 125, 0, 0,
		172, 182, 5, 126, 0, 0, 173, 182, 5, 127, 0, 0, 174, 182, 5, 128, 0, 0,
		175, 182, 5, 129, 0, 0, 176, 182, 5, 130, 0, 0, 177, 182, 5, 131, 0, 0,
		178, 182, 5, 132, 0, 0, 179, 182, 5, 133, 0, 0, 180, 182, 5, 137, 0, 0,
		181, 170, 1, 0, 0, 0, 181, 171, 1, 0, 0, 0, 181, 172, 1, 0, 0, 0, 181,
		173, 1, 0, 0, 0, 181, 174, 1, 0, 0, 0, 181, 175, 1, 0, 0, 0, 181, 176,
		1, 0, 0, 0, 181, 177, 1, 0, 0, 0, 181, 178, 1, 0, 0, 0, 181, 179, 1, 0,
		0, 0, 181, 180, 1, 0, 0, 0, 182, 41, 1, 0, 0, 0, 183, 186, 3, 44, 22, 0,
		184, 186, 3, 46, 23, 0, 185, 183, 1, 0, 0, 0, 185, 184, 1, 0, 0, 0, 186,
		43, 1, 0, 0, 0, 187, 188, 7, 2, 0, 0, 188, 45, 1, 0, 0, 0, 189, 190, 7,
		3, 0, 0, 190, 47, 1, 0, 0, 0, 18, 50, 59, 66, 73, 78, 83, 93, 102, 107,
		114, 122, 130, 134, 138, 161, 167, 181, 185,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CqlParserInit initializes any static state used to implement CqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CqlParserInit() {
	staticData := &CqlParserParserStaticData
	staticData.once.Do(cqlparserParserInit)
}

// NewCqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCqlParser(input antlr.TokenStream) *CqlParser {
	CqlParserInit()
	this := new(CqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CqlParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "CqlParser.g4"

	return this
}

// CqlParser tokens.
const (
	CqlParserEOF                    = antlr.TokenEOF
	CqlParserK_ADD                  = 1
	CqlParserK_AGGREGATE            = 2
	CqlParserK_ALL                  = 3
	CqlParserK_ALLOW                = 4
	CqlParserK_ALTER                = 5
	CqlParserK_AND                  = 6
	CqlParserK_APPLY                = 7
	CqlParserK_AS                   = 8
	CqlParserK_ASC                  = 9
	CqlParserK_ASCII                = 10
	CqlParserK_AUTHORIZE            = 11
	CqlParserK_BATCH                = 12
	CqlParserK_BEGIN                = 13
	CqlParserK_BIGINT               = 14
	CqlParserK_BLOB                 = 15
	CqlParserK_BOOLEAN              = 16
	CqlParserK_BY                   = 17
	CqlParserK_CALLED               = 18
	CqlParserK_CLUSTERING           = 19
	CqlParserK_COLUMNFAMILY         = 20
	CqlParserK_COMPACT              = 21
	CqlParserK_CONTAINS             = 22
	CqlParserK_COUNT                = 23
	CqlParserK_COUNTER              = 24
	CqlParserK_CREATE               = 25
	CqlParserK_CUSTOM               = 26
	CqlParserK_DATE                 = 27
	CqlParserK_DECIMAL              = 28
	CqlParserK_DELETE               = 29
	CqlParserK_DESC                 = 30
	CqlParserK_DESCRIBE             = 31
	CqlParserK_DISTINCT             = 32
	CqlParserK_DOUBLE               = 33
	CqlParserK_DROP                 = 34
	CqlParserK_ENTRIES              = 35
	CqlParserK_EXECUTE              = 36
	CqlParserK_EXISTS               = 37
	CqlParserK_FILTERING            = 38
	CqlParserK_FINALFUNC            = 39
	CqlParserK_FLOAT                = 40
	CqlParserK_FROM                 = 41
	CqlParserK_FROZEN               = 42
	CqlParserK_FULL                 = 43
	CqlParserK_FUNCTION             = 44
	CqlParserK_FUNCTIONS            = 45
	CqlParserK_GRANT                = 46
	CqlParserK_IF                   = 47
	CqlParserK_IN                   = 48
	CqlParserK_INDEX                = 49
	CqlParserK_INET                 = 50
	CqlParserK_INFINITY             = 51
	CqlParserK_INITCOND             = 52
	CqlParserK_INPUT                = 53
	CqlParserK_INSERT               = 54
	CqlParserK_INT                  = 55
	CqlParserK_INTO                 = 56
	CqlParserK_JSON                 = 57
	CqlParserK_KEY                  = 58
	CqlParserK_KEYS                 = 59
	CqlParserK_KEYSPACE             = 60
	CqlParserK_KEYSPACES            = 61
	CqlParserK_LANGUAGE             = 62
	CqlParserK_LIMIT                = 63
	CqlParserK_LIST                 = 64
	CqlParserK_LOGIN                = 65
	CqlParserK_MAP                  = 66
	CqlParserK_MODIFY               = 67
	CqlParserK_NAN                  = 68
	CqlParserK_NOLOGIN              = 69
	CqlParserK_NORECURSIVE          = 70
	CqlParserK_NOSUPERUSER          = 71
	CqlParserK_NOT                  = 72
	CqlParserK_NULL                 = 73
	CqlParserK_OF                   = 74
	CqlParserK_ON                   = 75
	CqlParserK_OPTIONS              = 76
	CqlParserK_OR                   = 77
	CqlParserK_ORDER                = 78
	CqlParserK_PASSWORD             = 79
	CqlParserK_PERMISSION           = 80
	CqlParserK_PERMISSIONS          = 81
	CqlParserK_PRIMARY              = 82
	CqlParserK_RENAME               = 83
	CqlParserK_REPLACE              = 84
	CqlParserK_RETURNS              = 85
	CqlParserK_REVOKE               = 86
	CqlParserK_ROLE                 = 87
	CqlParserK_ROLES                = 88
	CqlParserK_SCHEMA               = 89
	CqlParserK_SELECT               = 90
	CqlParserK_SET                  = 91
	CqlParserK_SFUNC                = 92
	CqlParserK_SMALLINT             = 93
	CqlParserK_STATIC               = 94
	CqlParserK_STORAGE              = 95
	CqlParserK_STYPE                = 96
	CqlParserK_SUPERUSER            = 97
	CqlParserK_TABLE                = 98
	CqlParserK_TEXT                 = 99
	CqlParserK_TIME                 = 100
	CqlParserK_TIMESTAMP            = 101
	CqlParserK_TIMEUUID             = 102
	CqlParserK_TINYINT              = 103
	CqlParserK_TO                   = 104
	CqlParserK_TOKEN                = 105
	CqlParserK_TRIGGER              = 106
	CqlParserK_TRUNCATE             = 107
	CqlParserK_TTL                  = 108
	CqlParserK_TUPLE                = 109
	CqlParserK_TYPE                 = 110
	CqlParserK_UNLOGGED             = 111
	CqlParserK_UPDATE               = 112
	CqlParserK_USE                  = 113
	CqlParserK_USER                 = 114
	CqlParserK_USERS                = 115
	CqlParserK_USING                = 116
	CqlParserK_UUID                 = 117
	CqlParserK_VALUES               = 118
	CqlParserK_VARCHAR              = 119
	CqlParserK_VARINT               = 120
	CqlParserK_WHERE                = 121
	CqlParserK_WITH                 = 122
	CqlParserK_WRITETIME            = 123
	CqlParserSEMICOLON              = 124
	CqlParserDQUOTE                 = 125
	CqlParserDOT                    = 126
	CqlParserCOMMA                  = 127
	CqlParserL_PAREN                = 128
	CqlParserR_PAREN                = 129
	CqlParserL_ANGLE_BRACKET        = 130
	CqlParserR_ANGLE_BRACKET        = 131
	CqlParserIDENTIFIER             = 132
	CqlParserIDENTIFIER_WITH_HYPHEN = 133
	CqlParserWS                     = 134
	CqlParserCOMMENT                = 135
	CqlParserMULTILINE_COMMENT      = 136
	CqlParserUNKNOWN                = 137
)

// CqlParser rules.
const (
	CqlParserRULE_root                 = 0
	CqlParserRULE_cqlStatement         = 1
	CqlParserRULE_createTable          = 2
	CqlParserRULE_columnDefinitionList = 3
	CqlParserRULE_columnDefinition     = 4
	CqlParserRULE_primaryKeyClause     = 5
	CqlParserRULE_primaryKey           = 6
	CqlParserRULE_partitionKey         = 7
	CqlParserRULE_clusteringColumns    = 8
	CqlParserRULE_columnName           = 9
	CqlParserRULE_tableName            = 10
	CqlParserRULE_keyspaceName         = 11
	CqlParserRULE_generalIdentifier    = 12
	CqlParserRULE_identifier           = 13
	CqlParserRULE_cqlType              = 14
	CqlParserRULE_primaryKeyKeywords   = 15
	CqlParserRULE_ifNotExist           = 16
	CqlParserRULE_cqlNativeType        = 17
	CqlParserRULE_cqlCollectionType    = 18
	CqlParserRULE_wihtTableOptions     = 19
	CqlParserRULE_nonSemicolonToken    = 20
	CqlParserRULE_keyword              = 21
	CqlParserRULE_reservedKeyword      = 22
	CqlParserRULE_nonReservedKeyword   = 23
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CqlStatement() ICqlStatementContext
	EOF() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) CqlStatement() ICqlStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICqlStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICqlStatementContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(CqlParserEOF, 0)
}

func (s *RootContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(CqlParserSEMICOLON, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *CqlParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CqlParserRULE_root)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.CqlStatement()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserSEMICOLON {
		{
			p.SetState(49)
			p.Match(CqlParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(52)
		p.Match(CqlParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICqlStatementContext is an interface to support dynamic dispatch.
type ICqlStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CreateTable() ICreateTableContext

	// IsCqlStatementContext differentiates from other interfaces.
	IsCqlStatementContext()
}

type CqlStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCqlStatementContext() *CqlStatementContext {
	var p = new(CqlStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlStatement
	return p
}

func InitEmptyCqlStatementContext(p *CqlStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlStatement
}

func (*CqlStatementContext) IsCqlStatementContext() {}

func NewCqlStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlStatementContext {
	var p = new(CqlStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_cqlStatement

	return p
}

func (s *CqlStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlStatementContext) CreateTable() ICreateTableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateTableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateTableContext)
}

func (s *CqlStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterCqlStatement(s)
	}
}

func (s *CqlStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitCqlStatement(s)
	}
}

func (p *CqlParser) CqlStatement() (localctx ICqlStatementContext) {
	localctx = NewCqlStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CqlParserRULE_cqlStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.CreateTable()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreateTableContext is an interface to support dynamic dispatch.
type ICreateTableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_CREATE() antlr.TerminalNode
	K_TABLE() antlr.TerminalNode
	TableName() ITableNameContext
	L_PAREN() antlr.TerminalNode
	ColumnDefinitionList() IColumnDefinitionListContext
	R_PAREN() antlr.TerminalNode
	IfNotExist() IIfNotExistContext
	WihtTableOptions() IWihtTableOptionsContext

	// IsCreateTableContext differentiates from other interfaces.
	IsCreateTableContext()
}

type CreateTableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateTableContext() *CreateTableContext {
	var p = new(CreateTableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_createTable
	return p
}

func InitEmptyCreateTableContext(p *CreateTableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_createTable
}

func (*CreateTableContext) IsCreateTableContext() {}

func NewCreateTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateTableContext {
	var p = new(CreateTableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_createTable

	return p
}

func (s *CreateTableContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateTableContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CREATE, 0)
}

func (s *CreateTableContext) K_TABLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TABLE, 0)
}

func (s *CreateTableContext) TableName() ITableNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *CreateTableContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserL_PAREN, 0)
}

func (s *CreateTableContext) ColumnDefinitionList() IColumnDefinitionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnDefinitionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnDefinitionListContext)
}

func (s *CreateTableContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserR_PAREN, 0)
}

func (s *CreateTableContext) IfNotExist() IIfNotExistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfNotExistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfNotExistContext)
}

func (s *CreateTableContext) WihtTableOptions() IWihtTableOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWihtTableOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWihtTableOptionsContext)
}

func (s *CreateTableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateTableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateTableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterCreateTable(s)
	}
}

func (s *CreateTableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitCreateTable(s)
	}
}

func (p *CqlParser) CreateTable() (localctx ICreateTableContext) {
	localctx = NewCreateTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CqlParserRULE_createTable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(CqlParserK_CREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Match(CqlParserK_TABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_IF {
		{
			p.SetState(58)
			p.IfNotExist()
		}

	}
	{
		p.SetState(61)
		p.TableName()
	}
	{
		p.SetState(62)
		p.Match(CqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.ColumnDefinitionList()
	}
	{
		p.SetState(64)
		p.Match(CqlParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_WITH {
		{
			p.SetState(65)
			p.WihtTableOptions()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnDefinitionListContext is an interface to support dynamic dispatch.
type IColumnDefinitionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumnDefinition() []IColumnDefinitionContext
	ColumnDefinition(i int) IColumnDefinitionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	PrimaryKeyClause() IPrimaryKeyClauseContext

	// IsColumnDefinitionListContext differentiates from other interfaces.
	IsColumnDefinitionListContext()
}

type ColumnDefinitionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnDefinitionListContext() *ColumnDefinitionListContext {
	var p = new(ColumnDefinitionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_columnDefinitionList
	return p
}

func InitEmptyColumnDefinitionListContext(p *ColumnDefinitionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_columnDefinitionList
}

func (*ColumnDefinitionListContext) IsColumnDefinitionListContext() {}

func NewColumnDefinitionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnDefinitionListContext {
	var p = new(ColumnDefinitionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_columnDefinitionList

	return p
}

func (s *ColumnDefinitionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnDefinitionListContext) AllColumnDefinition() []IColumnDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IColumnDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnDefinitionContext); ok {
			tst[i] = t.(IColumnDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *ColumnDefinitionListContext) ColumnDefinition(i int) IColumnDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnDefinitionContext)
}

func (s *ColumnDefinitionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CqlParserCOMMA)
}

func (s *ColumnDefinitionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CqlParserCOMMA, i)
}

func (s *ColumnDefinitionListContext) PrimaryKeyClause() IPrimaryKeyClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryKeyClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryKeyClauseContext)
}

func (s *ColumnDefinitionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnDefinitionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnDefinitionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterColumnDefinitionList(s)
	}
}

func (s *ColumnDefinitionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitColumnDefinitionList(s)
	}
}

func (p *CqlParser) ColumnDefinitionList() (localctx IColumnDefinitionListContext) {
	localctx = NewColumnDefinitionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CqlParserRULE_columnDefinitionList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.ColumnDefinition()
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(69)
				p.Match(CqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(70)
				p.ColumnDefinition()
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserCOMMA {
		{
			p.SetState(76)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.PrimaryKeyClause()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnDefinitionContext is an interface to support dynamic dispatch.
type IColumnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ColumnName() IColumnNameContext
	CqlType() ICqlTypeContext
	PrimaryKeyKeywords() IPrimaryKeyKeywordsContext

	// IsColumnDefinitionContext differentiates from other interfaces.
	IsColumnDefinitionContext()
}

type ColumnDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnDefinitionContext() *ColumnDefinitionContext {
	var p = new(ColumnDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_columnDefinition
	return p
}

func InitEmptyColumnDefinitionContext(p *ColumnDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_columnDefinition
}

func (*ColumnDefinitionContext) IsColumnDefinitionContext() {}

func NewColumnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnDefinitionContext {
	var p = new(ColumnDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_columnDefinition

	return p
}

func (s *ColumnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnDefinitionContext) ColumnName() IColumnNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ColumnDefinitionContext) CqlType() ICqlTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICqlTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICqlTypeContext)
}

func (s *ColumnDefinitionContext) PrimaryKeyKeywords() IPrimaryKeyKeywordsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryKeyKeywordsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryKeyKeywordsContext)
}

func (s *ColumnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterColumnDefinition(s)
	}
}

func (s *ColumnDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitColumnDefinition(s)
	}
}

func (p *CqlParser) ColumnDefinition() (localctx IColumnDefinitionContext) {
	localctx = NewColumnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CqlParserRULE_columnDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.ColumnName()
	}
	{
		p.SetState(81)
		p.CqlType()
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_PRIMARY {
		{
			p.SetState(82)
			p.PrimaryKeyKeywords()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryKeyClauseContext is an interface to support dynamic dispatch.
type IPrimaryKeyClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryKeyKeywords() IPrimaryKeyKeywordsContext
	L_PAREN() antlr.TerminalNode
	PrimaryKey() IPrimaryKeyContext
	R_PAREN() antlr.TerminalNode

	// IsPrimaryKeyClauseContext differentiates from other interfaces.
	IsPrimaryKeyClauseContext()
}

type PrimaryKeyClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryKeyClauseContext() *PrimaryKeyClauseContext {
	var p = new(PrimaryKeyClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_primaryKeyClause
	return p
}

func InitEmptyPrimaryKeyClauseContext(p *PrimaryKeyClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_primaryKeyClause
}

func (*PrimaryKeyClauseContext) IsPrimaryKeyClauseContext() {}

func NewPrimaryKeyClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryKeyClauseContext {
	var p = new(PrimaryKeyClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_primaryKeyClause

	return p
}

func (s *PrimaryKeyClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryKeyClauseContext) PrimaryKeyKeywords() IPrimaryKeyKeywordsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryKeyKeywordsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryKeyKeywordsContext)
}

func (s *PrimaryKeyClauseContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserL_PAREN, 0)
}

func (s *PrimaryKeyClauseContext) PrimaryKey() IPrimaryKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryKeyContext)
}

func (s *PrimaryKeyClauseContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserR_PAREN, 0)
}

func (s *PrimaryKeyClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryKeyClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryKeyClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterPrimaryKeyClause(s)
	}
}

func (s *PrimaryKeyClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitPrimaryKeyClause(s)
	}
}

func (p *CqlParser) PrimaryKeyClause() (localctx IPrimaryKeyClauseContext) {
	localctx = NewPrimaryKeyClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CqlParserRULE_primaryKeyClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.PrimaryKeyKeywords()
	}
	{
		p.SetState(86)
		p.Match(CqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.PrimaryKey()
	}
	{
		p.SetState(88)
		p.Match(CqlParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryKeyContext is an interface to support dynamic dispatch.
type IPrimaryKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PartitionKey() IPartitionKeyContext
	COMMA() antlr.TerminalNode
	ClusteringColumns() IClusteringColumnsContext

	// IsPrimaryKeyContext differentiates from other interfaces.
	IsPrimaryKeyContext()
}

type PrimaryKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryKeyContext() *PrimaryKeyContext {
	var p = new(PrimaryKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_primaryKey
	return p
}

func InitEmptyPrimaryKeyContext(p *PrimaryKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_primaryKey
}

func (*PrimaryKeyContext) IsPrimaryKeyContext() {}

func NewPrimaryKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryKeyContext {
	var p = new(PrimaryKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_primaryKey

	return p
}

func (s *PrimaryKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryKeyContext) PartitionKey() IPartitionKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPartitionKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPartitionKeyContext)
}

func (s *PrimaryKeyContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CqlParserCOMMA, 0)
}

func (s *PrimaryKeyContext) ClusteringColumns() IClusteringColumnsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClusteringColumnsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClusteringColumnsContext)
}

func (s *PrimaryKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterPrimaryKey(s)
	}
}

func (s *PrimaryKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitPrimaryKey(s)
	}
}

func (p *CqlParser) PrimaryKey() (localctx IPrimaryKeyContext) {
	localctx = NewPrimaryKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CqlParserRULE_primaryKey)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.PartitionKey()
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserCOMMA {
		{
			p.SetState(91)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.ClusteringColumns()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPartitionKeyContext is an interface to support dynamic dispatch.
type IPartitionKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumnName() []IColumnNameContext
	ColumnName(i int) IColumnNameContext
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPartitionKeyContext differentiates from other interfaces.
	IsPartitionKeyContext()
}

type PartitionKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPartitionKeyContext() *PartitionKeyContext {
	var p = new(PartitionKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_partitionKey
	return p
}

func InitEmptyPartitionKeyContext(p *PartitionKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_partitionKey
}

func (*PartitionKeyContext) IsPartitionKeyContext() {}

func NewPartitionKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PartitionKeyContext {
	var p = new(PartitionKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_partitionKey

	return p
}

func (s *PartitionKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *PartitionKeyContext) AllColumnName() []IColumnNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnNameContext); ok {
			len++
		}
	}

	tst := make([]IColumnNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnNameContext); ok {
			tst[i] = t.(IColumnNameContext)
			i++
		}
	}

	return tst
}

func (s *PartitionKeyContext) ColumnName(i int) IColumnNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *PartitionKeyContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserL_PAREN, 0)
}

func (s *PartitionKeyContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserR_PAREN, 0)
}

func (s *PartitionKeyContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CqlParserCOMMA)
}

func (s *PartitionKeyContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CqlParserCOMMA, i)
}

func (s *PartitionKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PartitionKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PartitionKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterPartitionKey(s)
	}
}

func (s *PartitionKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitPartitionKey(s)
	}
}

func (p *CqlParser) PartitionKey() (localctx IPartitionKeyContext) {
	localctx = NewPartitionKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CqlParserRULE_partitionKey)
	var _la int

	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_AGGREGATE, CqlParserK_ALL, CqlParserK_AS, CqlParserK_ASCII, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_CALLED, CqlParserK_CLUSTERING, CqlParserK_COMPACT, CqlParserK_CONTAINS, CqlParserK_COUNT, CqlParserK_COUNTER, CqlParserK_CUSTOM, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DISTINCT, CqlParserK_DOUBLE, CqlParserK_EXISTS, CqlParserK_FILTERING, CqlParserK_FINALFUNC, CqlParserK_FLOAT, CqlParserK_FROZEN, CqlParserK_FUNCTION, CqlParserK_FUNCTIONS, CqlParserK_INET, CqlParserK_INITCOND, CqlParserK_INPUT, CqlParserK_INT, CqlParserK_JSON, CqlParserK_KEY, CqlParserK_KEYS, CqlParserK_KEYSPACES, CqlParserK_LANGUAGE, CqlParserK_LIST, CqlParserK_LOGIN, CqlParserK_MAP, CqlParserK_NOLOGIN, CqlParserK_NOSUPERUSER, CqlParserK_OPTIONS, CqlParserK_PASSWORD, CqlParserK_PERMISSION, CqlParserK_PERMISSIONS, CqlParserK_RETURNS, CqlParserK_ROLE, CqlParserK_ROLES, CqlParserK_SFUNC, CqlParserK_SMALLINT, CqlParserK_STATIC, CqlParserK_STORAGE, CqlParserK_STYPE, CqlParserK_SUPERUSER, CqlParserK_TEXT, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_TRIGGER, CqlParserK_TTL, CqlParserK_TUPLE, CqlParserK_TYPE, CqlParserK_USER, CqlParserK_USERS, CqlParserK_UUID, CqlParserK_VALUES, CqlParserK_VARCHAR, CqlParserK_VARINT, CqlParserK_WRITETIME, CqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.ColumnName()
		}

	case CqlParserL_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.Match(CqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.ColumnName()
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CqlParserCOMMA {
			{
				p.SetState(98)
				p.Match(CqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(99)
				p.ColumnName()
			}

			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(105)
			p.Match(CqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClusteringColumnsContext is an interface to support dynamic dispatch.
type IClusteringColumnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumnName() []IColumnNameContext
	ColumnName(i int) IColumnNameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsClusteringColumnsContext differentiates from other interfaces.
	IsClusteringColumnsContext()
}

type ClusteringColumnsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClusteringColumnsContext() *ClusteringColumnsContext {
	var p = new(ClusteringColumnsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_clusteringColumns
	return p
}

func InitEmptyClusteringColumnsContext(p *ClusteringColumnsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_clusteringColumns
}

func (*ClusteringColumnsContext) IsClusteringColumnsContext() {}

func NewClusteringColumnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClusteringColumnsContext {
	var p = new(ClusteringColumnsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_clusteringColumns

	return p
}

func (s *ClusteringColumnsContext) GetParser() antlr.Parser { return s.parser }

func (s *ClusteringColumnsContext) AllColumnName() []IColumnNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnNameContext); ok {
			len++
		}
	}

	tst := make([]IColumnNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnNameContext); ok {
			tst[i] = t.(IColumnNameContext)
			i++
		}
	}

	return tst
}

func (s *ClusteringColumnsContext) ColumnName(i int) IColumnNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ClusteringColumnsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CqlParserCOMMA)
}

func (s *ClusteringColumnsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CqlParserCOMMA, i)
}

func (s *ClusteringColumnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClusteringColumnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClusteringColumnsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterClusteringColumns(s)
	}
}

func (s *ClusteringColumnsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitClusteringColumns(s)
	}
}

func (p *CqlParser) ClusteringColumns() (localctx IClusteringColumnsContext) {
	localctx = NewClusteringColumnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CqlParserRULE_clusteringColumns)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.ColumnName()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CqlParserCOMMA {
		{
			p.SetState(110)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.ColumnName()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnNameContext is an interface to support dynamic dispatch.
type IColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

	// IsColumnNameContext differentiates from other interfaces.
	IsColumnNameContext()
}

type ColumnNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnNameContext() *ColumnNameContext {
	var p = new(ColumnNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_columnName
	return p
}

func InitEmptyColumnNameContext(p *ColumnNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_columnName
}

func (*ColumnNameContext) IsColumnNameContext() {}

func NewColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNameContext {
	var p = new(ColumnNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_columnName

	return p
}

func (s *ColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnNameContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterColumnName(s)
	}
}

func (s *ColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitColumnName(s)
	}
}

func (p *CqlParser) ColumnName() (localctx IColumnNameContext) {
	localctx = NewColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CqlParserRULE_columnName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Identifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	KeyspaceName() IKeyspaceNameContext
	DOT() antlr.TerminalNode

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_tableName
	return p
}

func InitEmptyTableNameContext(p *TableNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_tableName
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *TableNameContext) KeyspaceName() IKeyspaceNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyspaceNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyspaceNameContext)
}

func (s *TableNameContext) DOT() antlr.TerminalNode {
	return s.GetToken(CqlParserDOT, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (p *CqlParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CqlParserRULE_tableName)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(122)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(119)
			p.KeyspaceName()
		}
		{
			p.SetState(120)
			p.Match(CqlParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(124)
		p.Identifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeyspaceNameContext is an interface to support dynamic dispatch.
type IKeyspaceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GeneralIdentifier() IGeneralIdentifierContext

	// IsKeyspaceNameContext differentiates from other interfaces.
	IsKeyspaceNameContext()
}

type KeyspaceNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyspaceNameContext() *KeyspaceNameContext {
	var p = new(KeyspaceNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_keyspaceName
	return p
}

func InitEmptyKeyspaceNameContext(p *KeyspaceNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_keyspaceName
}

func (*KeyspaceNameContext) IsKeyspaceNameContext() {}

func NewKeyspaceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyspaceNameContext {
	var p = new(KeyspaceNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_keyspaceName

	return p
}

func (s *KeyspaceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyspaceNameContext) GeneralIdentifier() IGeneralIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeneralIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeneralIdentifierContext)
}

func (s *KeyspaceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyspaceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyspaceNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterKeyspaceName(s)
	}
}

func (s *KeyspaceNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitKeyspaceName(s)
	}
}

func (p *CqlParser) KeyspaceName() (localctx IKeyspaceNameContext) {
	localctx = NewKeyspaceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CqlParserRULE_keyspaceName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.GeneralIdentifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGeneralIdentifierContext is an interface to support dynamic dispatch.
type IGeneralIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER_WITH_HYPHEN() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsGeneralIdentifierContext differentiates from other interfaces.
	IsGeneralIdentifierContext()
}

type GeneralIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeneralIdentifierContext() *GeneralIdentifierContext {
	var p = new(GeneralIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_generalIdentifier
	return p
}

func InitEmptyGeneralIdentifierContext(p *GeneralIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_generalIdentifier
}

func (*GeneralIdentifierContext) IsGeneralIdentifierContext() {}

func NewGeneralIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeneralIdentifierContext {
	var p = new(GeneralIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_generalIdentifier

	return p
}

func (s *GeneralIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *GeneralIdentifierContext) IDENTIFIER_WITH_HYPHEN() antlr.TerminalNode {
	return s.GetToken(CqlParserIDENTIFIER_WITH_HYPHEN, 0)
}

func (s *GeneralIdentifierContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *GeneralIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeneralIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeneralIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterGeneralIdentifier(s)
	}
}

func (s *GeneralIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitGeneralIdentifier(s)
	}
}

func (p *CqlParser) GeneralIdentifier() (localctx IGeneralIdentifierContext) {
	localctx = NewGeneralIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CqlParserRULE_generalIdentifier)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserIDENTIFIER_WITH_HYPHEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Match(CqlParserIDENTIFIER_WITH_HYPHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserK_AGGREGATE, CqlParserK_ALL, CqlParserK_AS, CqlParserK_ASCII, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_CALLED, CqlParserK_CLUSTERING, CqlParserK_COMPACT, CqlParserK_CONTAINS, CqlParserK_COUNT, CqlParserK_COUNTER, CqlParserK_CUSTOM, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DISTINCT, CqlParserK_DOUBLE, CqlParserK_EXISTS, CqlParserK_FILTERING, CqlParserK_FINALFUNC, CqlParserK_FLOAT, CqlParserK_FROZEN, CqlParserK_FUNCTION, CqlParserK_FUNCTIONS, CqlParserK_INET, CqlParserK_INITCOND, CqlParserK_INPUT, CqlParserK_INT, CqlParserK_JSON, CqlParserK_KEY, CqlParserK_KEYS, CqlParserK_KEYSPACES, CqlParserK_LANGUAGE, CqlParserK_LIST, CqlParserK_LOGIN, CqlParserK_MAP, CqlParserK_NOLOGIN, CqlParserK_NOSUPERUSER, CqlParserK_OPTIONS, CqlParserK_PASSWORD, CqlParserK_PERMISSION, CqlParserK_PERMISSIONS, CqlParserK_RETURNS, CqlParserK_ROLE, CqlParserK_ROLES, CqlParserK_SFUNC, CqlParserK_SMALLINT, CqlParserK_STATIC, CqlParserK_STORAGE, CqlParserK_STYPE, CqlParserK_SUPERUSER, CqlParserK_TEXT, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_TRIGGER, CqlParserK_TTL, CqlParserK_TUPLE, CqlParserK_TYPE, CqlParserK_USER, CqlParserK_USERS, CqlParserK_UUID, CqlParserK_VALUES, CqlParserK_VARCHAR, CqlParserK_VARINT, CqlParserK_WRITETIME, CqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Identifier()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NonReservedKeyword() INonReservedKeywordContext
	IDENTIFIER() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) NonReservedKeyword() INonReservedKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INonReservedKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INonReservedKeywordContext)
}

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CqlParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *CqlParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CqlParserRULE_identifier)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_AGGREGATE, CqlParserK_ALL, CqlParserK_AS, CqlParserK_ASCII, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_CALLED, CqlParserK_CLUSTERING, CqlParserK_COMPACT, CqlParserK_CONTAINS, CqlParserK_COUNT, CqlParserK_COUNTER, CqlParserK_CUSTOM, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DISTINCT, CqlParserK_DOUBLE, CqlParserK_EXISTS, CqlParserK_FILTERING, CqlParserK_FINALFUNC, CqlParserK_FLOAT, CqlParserK_FROZEN, CqlParserK_FUNCTION, CqlParserK_FUNCTIONS, CqlParserK_INET, CqlParserK_INITCOND, CqlParserK_INPUT, CqlParserK_INT, CqlParserK_JSON, CqlParserK_KEY, CqlParserK_KEYS, CqlParserK_KEYSPACES, CqlParserK_LANGUAGE, CqlParserK_LIST, CqlParserK_LOGIN, CqlParserK_MAP, CqlParserK_NOLOGIN, CqlParserK_NOSUPERUSER, CqlParserK_OPTIONS, CqlParserK_PASSWORD, CqlParserK_PERMISSION, CqlParserK_PERMISSIONS, CqlParserK_RETURNS, CqlParserK_ROLE, CqlParserK_ROLES, CqlParserK_SFUNC, CqlParserK_SMALLINT, CqlParserK_STATIC, CqlParserK_STORAGE, CqlParserK_STYPE, CqlParserK_SUPERUSER, CqlParserK_TEXT, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_TRIGGER, CqlParserK_TTL, CqlParserK_TUPLE, CqlParserK_TYPE, CqlParserK_USER, CqlParserK_USERS, CqlParserK_UUID, CqlParserK_VALUES, CqlParserK_VARCHAR, CqlParserK_VARINT, CqlParserK_WRITETIME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.NonReservedKeyword()
		}

	case CqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.Match(CqlParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICqlTypeContext is an interface to support dynamic dispatch.
type ICqlTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CqlNativeType() ICqlNativeTypeContext
	CqlCollectionType() ICqlCollectionTypeContext

	// IsCqlTypeContext differentiates from other interfaces.
	IsCqlTypeContext()
}

type CqlTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCqlTypeContext() *CqlTypeContext {
	var p = new(CqlTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlType
	return p
}

func InitEmptyCqlTypeContext(p *CqlTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlType
}

func (*CqlTypeContext) IsCqlTypeContext() {}

func NewCqlTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlTypeContext {
	var p = new(CqlTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_cqlType

	return p
}

func (s *CqlTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlTypeContext) CqlNativeType() ICqlNativeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICqlNativeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICqlNativeTypeContext)
}

func (s *CqlTypeContext) CqlCollectionType() ICqlCollectionTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICqlCollectionTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICqlCollectionTypeContext)
}

func (s *CqlTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterCqlType(s)
	}
}

func (s *CqlTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitCqlType(s)
	}
}

func (p *CqlParser) CqlType() (localctx ICqlTypeContext) {
	localctx = NewCqlTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CqlParserRULE_cqlType)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_ASCII, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_COUNTER, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DOUBLE, CqlParserK_FLOAT, CqlParserK_INET, CqlParserK_INT, CqlParserK_SMALLINT, CqlParserK_TEXT, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_UUID, CqlParserK_VARCHAR, CqlParserK_VARINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.CqlNativeType()
		}

	case CqlParserK_LIST, CqlParserK_MAP, CqlParserK_SET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.CqlCollectionType()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryKeyKeywordsContext is an interface to support dynamic dispatch.
type IPrimaryKeyKeywordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_PRIMARY() antlr.TerminalNode
	K_KEY() antlr.TerminalNode

	// IsPrimaryKeyKeywordsContext differentiates from other interfaces.
	IsPrimaryKeyKeywordsContext()
}

type PrimaryKeyKeywordsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryKeyKeywordsContext() *PrimaryKeyKeywordsContext {
	var p = new(PrimaryKeyKeywordsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_primaryKeyKeywords
	return p
}

func InitEmptyPrimaryKeyKeywordsContext(p *PrimaryKeyKeywordsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_primaryKeyKeywords
}

func (*PrimaryKeyKeywordsContext) IsPrimaryKeyKeywordsContext() {}

func NewPrimaryKeyKeywordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryKeyKeywordsContext {
	var p = new(PrimaryKeyKeywordsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_primaryKeyKeywords

	return p
}

func (s *PrimaryKeyKeywordsContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryKeyKeywordsContext) K_PRIMARY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_PRIMARY, 0)
}

func (s *PrimaryKeyKeywordsContext) K_KEY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_KEY, 0)
}

func (s *PrimaryKeyKeywordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryKeyKeywordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryKeyKeywordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterPrimaryKeyKeywords(s)
	}
}

func (s *PrimaryKeyKeywordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitPrimaryKeyKeywords(s)
	}
}

func (p *CqlParser) PrimaryKeyKeywords() (localctx IPrimaryKeyKeywordsContext) {
	localctx = NewPrimaryKeyKeywordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CqlParserRULE_primaryKeyKeywords)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(CqlParserK_PRIMARY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(CqlParserK_KEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfNotExistContext is an interface to support dynamic dispatch.
type IIfNotExistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_IF() antlr.TerminalNode
	K_NOT() antlr.TerminalNode
	K_EXISTS() antlr.TerminalNode

	// IsIfNotExistContext differentiates from other interfaces.
	IsIfNotExistContext()
}

type IfNotExistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfNotExistContext() *IfNotExistContext {
	var p = new(IfNotExistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_ifNotExist
	return p
}

func InitEmptyIfNotExistContext(p *IfNotExistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_ifNotExist
}

func (*IfNotExistContext) IsIfNotExistContext() {}

func NewIfNotExistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfNotExistContext {
	var p = new(IfNotExistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_ifNotExist

	return p
}

func (s *IfNotExistContext) GetParser() antlr.Parser { return s.parser }

func (s *IfNotExistContext) K_IF() antlr.TerminalNode {
	return s.GetToken(CqlParserK_IF, 0)
}

func (s *IfNotExistContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NOT, 0)
}

func (s *IfNotExistContext) K_EXISTS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_EXISTS, 0)
}

func (s *IfNotExistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfNotExistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfNotExistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterIfNotExist(s)
	}
}

func (s *IfNotExistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitIfNotExist(s)
	}
}

func (p *CqlParser) IfNotExist() (localctx IIfNotExistContext) {
	localctx = NewIfNotExistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CqlParserRULE_ifNotExist)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(CqlParserK_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(CqlParserK_NOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(CqlParserK_EXISTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICqlNativeTypeContext is an interface to support dynamic dispatch.
type ICqlNativeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_ASCII() antlr.TerminalNode
	K_BIGINT() antlr.TerminalNode
	K_BLOB() antlr.TerminalNode
	K_BOOLEAN() antlr.TerminalNode
	K_COUNTER() antlr.TerminalNode
	K_DATE() antlr.TerminalNode
	K_DECIMAL() antlr.TerminalNode
	K_DOUBLE() antlr.TerminalNode
	K_FLOAT() antlr.TerminalNode
	K_INET() antlr.TerminalNode
	K_INT() antlr.TerminalNode
	K_SMALLINT() antlr.TerminalNode
	K_TEXT() antlr.TerminalNode
	K_TIME() antlr.TerminalNode
	K_TIMESTAMP() antlr.TerminalNode
	K_TIMEUUID() antlr.TerminalNode
	K_TINYINT() antlr.TerminalNode
	K_UUID() antlr.TerminalNode
	K_VARCHAR() antlr.TerminalNode
	K_VARINT() antlr.TerminalNode

	// IsCqlNativeTypeContext differentiates from other interfaces.
	IsCqlNativeTypeContext()
}

type CqlNativeTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCqlNativeTypeContext() *CqlNativeTypeContext {
	var p = new(CqlNativeTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlNativeType
	return p
}

func InitEmptyCqlNativeTypeContext(p *CqlNativeTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlNativeType
}

func (*CqlNativeTypeContext) IsCqlNativeTypeContext() {}

func NewCqlNativeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlNativeTypeContext {
	var p = new(CqlNativeTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_cqlNativeType

	return p
}

func (s *CqlNativeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlNativeTypeContext) K_ASCII() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ASCII, 0)
}

func (s *CqlNativeTypeContext) K_BIGINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BIGINT, 0)
}

func (s *CqlNativeTypeContext) K_BLOB() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BLOB, 0)
}

func (s *CqlNativeTypeContext) K_BOOLEAN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BOOLEAN, 0)
}

func (s *CqlNativeTypeContext) K_COUNTER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_COUNTER, 0)
}

func (s *CqlNativeTypeContext) K_DATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DATE, 0)
}

func (s *CqlNativeTypeContext) K_DECIMAL() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DECIMAL, 0)
}

func (s *CqlNativeTypeContext) K_DOUBLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DOUBLE, 0)
}

func (s *CqlNativeTypeContext) K_FLOAT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FLOAT, 0)
}

func (s *CqlNativeTypeContext) K_INET() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INET, 0)
}

func (s *CqlNativeTypeContext) K_INT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INT, 0)
}

func (s *CqlNativeTypeContext) K_SMALLINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SMALLINT, 0)
}

func (s *CqlNativeTypeContext) K_TEXT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TEXT, 0)
}

func (s *CqlNativeTypeContext) K_TIME() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIME, 0)
}

func (s *CqlNativeTypeContext) K_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIMESTAMP, 0)
}

func (s *CqlNativeTypeContext) K_TIMEUUID() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIMEUUID, 0)
}

func (s *CqlNativeTypeContext) K_TINYINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TINYINT, 0)
}

func (s *CqlNativeTypeContext) K_UUID() antlr.TerminalNode {
	return s.GetToken(CqlParserK_UUID, 0)
}

func (s *CqlNativeTypeContext) K_VARCHAR() antlr.TerminalNode {
	return s.GetToken(CqlParserK_VARCHAR, 0)
}

func (s *CqlNativeTypeContext) K_VARINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_VARINT, 0)
}

func (s *CqlNativeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlNativeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlNativeTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterCqlNativeType(s)
	}
}

func (s *CqlNativeTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitCqlNativeType(s)
	}
}

func (p *CqlParser) CqlNativeType() (localctx ICqlNativeTypeContext) {
	localctx = NewCqlNativeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CqlParserRULE_cqlNativeType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&37155805446915072) != 0) || ((int64((_la-93)) & ^0x3f) == 0 && ((int64(1)<<(_la-93))&218105793) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICqlCollectionTypeContext is an interface to support dynamic dispatch.
type ICqlCollectionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_MAP() antlr.TerminalNode
	L_ANGLE_BRACKET() antlr.TerminalNode
	AllCqlNativeType() []ICqlNativeTypeContext
	CqlNativeType(i int) ICqlNativeTypeContext
	COMMA() antlr.TerminalNode
	R_ANGLE_BRACKET() antlr.TerminalNode
	K_SET() antlr.TerminalNode
	K_LIST() antlr.TerminalNode

	// IsCqlCollectionTypeContext differentiates from other interfaces.
	IsCqlCollectionTypeContext()
}

type CqlCollectionTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCqlCollectionTypeContext() *CqlCollectionTypeContext {
	var p = new(CqlCollectionTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlCollectionType
	return p
}

func InitEmptyCqlCollectionTypeContext(p *CqlCollectionTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlCollectionType
}

func (*CqlCollectionTypeContext) IsCqlCollectionTypeContext() {}

func NewCqlCollectionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlCollectionTypeContext {
	var p = new(CqlCollectionTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_cqlCollectionType

	return p
}

func (s *CqlCollectionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlCollectionTypeContext) K_MAP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_MAP, 0)
}

func (s *CqlCollectionTypeContext) L_ANGLE_BRACKET() antlr.TerminalNode {
	return s.GetToken(CqlParserL_ANGLE_BRACKET, 0)
}

func (s *CqlCollectionTypeContext) AllCqlNativeType() []ICqlNativeTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICqlNativeTypeContext); ok {
			len++
		}
	}

	tst := make([]ICqlNativeTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICqlNativeTypeContext); ok {
			tst[i] = t.(ICqlNativeTypeContext)
			i++
		}
	}

	return tst
}

func (s *CqlCollectionTypeContext) CqlNativeType(i int) ICqlNativeTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICqlNativeTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICqlNativeTypeContext)
}

func (s *CqlCollectionTypeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CqlParserCOMMA, 0)
}

func (s *CqlCollectionTypeContext) R_ANGLE_BRACKET() antlr.TerminalNode {
	return s.GetToken(CqlParserR_ANGLE_BRACKET, 0)
}

func (s *CqlCollectionTypeContext) K_SET() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SET, 0)
}

func (s *CqlCollectionTypeContext) K_LIST() antlr.TerminalNode {
	return s.GetToken(CqlParserK_LIST, 0)
}

func (s *CqlCollectionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlCollectionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlCollectionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterCqlCollectionType(s)
	}
}

func (s *CqlCollectionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitCqlCollectionType(s)
	}
}

func (p *CqlParser) CqlCollectionType() (localctx ICqlCollectionTypeContext) {
	localctx = NewCqlCollectionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CqlParserRULE_cqlCollectionType)
	var _la int

	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_MAP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Match(CqlParserK_MAP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Match(CqlParserL_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.CqlNativeType()
		}
		{
			p.SetState(152)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.CqlNativeType()
		}
		{
			p.SetState(154)
			p.Match(CqlParserR_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserK_LIST, CqlParserK_SET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CqlParserK_LIST || _la == CqlParserK_SET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(157)
			p.Match(CqlParserL_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.CqlNativeType()
		}
		{
			p.SetState(159)
			p.Match(CqlParserR_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWihtTableOptionsContext is an interface to support dynamic dispatch.
type IWihtTableOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_WITH() antlr.TerminalNode
	AllNonSemicolonToken() []INonSemicolonTokenContext
	NonSemicolonToken(i int) INonSemicolonTokenContext

	// IsWihtTableOptionsContext differentiates from other interfaces.
	IsWihtTableOptionsContext()
}

type WihtTableOptionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWihtTableOptionsContext() *WihtTableOptionsContext {
	var p = new(WihtTableOptionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_wihtTableOptions
	return p
}

func InitEmptyWihtTableOptionsContext(p *WihtTableOptionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_wihtTableOptions
}

func (*WihtTableOptionsContext) IsWihtTableOptionsContext() {}

func NewWihtTableOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WihtTableOptionsContext {
	var p = new(WihtTableOptionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_wihtTableOptions

	return p
}

func (s *WihtTableOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *WihtTableOptionsContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(CqlParserK_WITH, 0)
}

func (s *WihtTableOptionsContext) AllNonSemicolonToken() []INonSemicolonTokenContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INonSemicolonTokenContext); ok {
			len++
		}
	}

	tst := make([]INonSemicolonTokenContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INonSemicolonTokenContext); ok {
			tst[i] = t.(INonSemicolonTokenContext)
			i++
		}
	}

	return tst
}

func (s *WihtTableOptionsContext) NonSemicolonToken(i int) INonSemicolonTokenContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INonSemicolonTokenContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INonSemicolonTokenContext)
}

func (s *WihtTableOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WihtTableOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WihtTableOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterWihtTableOptions(s)
	}
}

func (s *WihtTableOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitWihtTableOptions(s)
	}
}

func (p *CqlParser) WihtTableOptions() (localctx IWihtTableOptionsContext) {
	localctx = NewWihtTableOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CqlParserRULE_wihtTableOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(CqlParserK_WITH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1152921504606846977) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&575) != 0) {
		{
			p.SetState(164)
			p.NonSemicolonToken()
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INonSemicolonTokenContext is an interface to support dynamic dispatch.
type INonSemicolonTokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword() IKeywordContext
	DQUOTE() antlr.TerminalNode
	DOT() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	L_ANGLE_BRACKET() antlr.TerminalNode
	R_ANGLE_BRACKET() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	IDENTIFIER_WITH_HYPHEN() antlr.TerminalNode
	UNKNOWN() antlr.TerminalNode

	// IsNonSemicolonTokenContext differentiates from other interfaces.
	IsNonSemicolonTokenContext()
}

type NonSemicolonTokenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonSemicolonTokenContext() *NonSemicolonTokenContext {
	var p = new(NonSemicolonTokenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_nonSemicolonToken
	return p
}

func InitEmptyNonSemicolonTokenContext(p *NonSemicolonTokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_nonSemicolonToken
}

func (*NonSemicolonTokenContext) IsNonSemicolonTokenContext() {}

func NewNonSemicolonTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonSemicolonTokenContext {
	var p = new(NonSemicolonTokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_nonSemicolonToken

	return p
}

func (s *NonSemicolonTokenContext) GetParser() antlr.Parser { return s.parser }

func (s *NonSemicolonTokenContext) Keyword() IKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *NonSemicolonTokenContext) DQUOTE() antlr.TerminalNode {
	return s.GetToken(CqlParserDQUOTE, 0)
}

func (s *NonSemicolonTokenContext) DOT() antlr.TerminalNode {
	return s.GetToken(CqlParserDOT, 0)
}

func (s *NonSemicolonTokenContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CqlParserCOMMA, 0)
}

func (s *NonSemicolonTokenContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserL_PAREN, 0)
}

func (s *NonSemicolonTokenContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserR_PAREN, 0)
}

func (s *NonSemicolonTokenContext) L_ANGLE_BRACKET() antlr.TerminalNode {
	return s.GetToken(CqlParserL_ANGLE_BRACKET, 0)
}

func (s *NonSemicolonTokenContext) R_ANGLE_BRACKET() antlr.TerminalNode {
	return s.GetToken(CqlParserR_ANGLE_BRACKET, 0)
}

func (s *NonSemicolonTokenContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CqlParserIDENTIFIER, 0)
}

func (s *NonSemicolonTokenContext) IDENTIFIER_WITH_HYPHEN() antlr.TerminalNode {
	return s.GetToken(CqlParserIDENTIFIER_WITH_HYPHEN, 0)
}

func (s *NonSemicolonTokenContext) UNKNOWN() antlr.TerminalNode {
	return s.GetToken(CqlParserUNKNOWN, 0)
}

func (s *NonSemicolonTokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonSemicolonTokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonSemicolonTokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterNonSemicolonToken(s)
	}
}

func (s *NonSemicolonTokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitNonSemicolonToken(s)
	}
}

func (p *CqlParser) NonSemicolonToken() (localctx INonSemicolonTokenContext) {
	localctx = NewNonSemicolonTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CqlParserRULE_nonSemicolonToken)
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_ADD, CqlParserK_AGGREGATE, CqlParserK_ALL, CqlParserK_ALLOW, CqlParserK_ALTER, CqlParserK_AND, CqlParserK_APPLY, CqlParserK_AS, CqlParserK_ASC, CqlParserK_ASCII, CqlParserK_AUTHORIZE, CqlParserK_BATCH, CqlParserK_BEGIN, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_BY, CqlParserK_CALLED, CqlParserK_CLUSTERING, CqlParserK_COLUMNFAMILY, CqlParserK_COMPACT, CqlParserK_CONTAINS, CqlParserK_COUNT, CqlParserK_COUNTER, CqlParserK_CREATE, CqlParserK_CUSTOM, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DELETE, CqlParserK_DESC, CqlParserK_DESCRIBE, CqlParserK_DISTINCT, CqlParserK_DOUBLE, CqlParserK_DROP, CqlParserK_ENTRIES, CqlParserK_EXECUTE, CqlParserK_EXISTS, CqlParserK_FILTERING, CqlParserK_FINALFUNC, CqlParserK_FLOAT, CqlParserK_FROM, CqlParserK_FROZEN, CqlParserK_FULL, CqlParserK_FUNCTION, CqlParserK_FUNCTIONS, CqlParserK_GRANT, CqlParserK_IF, CqlParserK_IN, CqlParserK_INDEX, CqlParserK_INET, CqlParserK_INFINITY, CqlParserK_INITCOND, CqlParserK_INPUT, CqlParserK_INSERT, CqlParserK_INT, CqlParserK_INTO, CqlParserK_JSON, CqlParserK_KEY, CqlParserK_KEYS, CqlParserK_KEYSPACE, CqlParserK_KEYSPACES, CqlParserK_LANGUAGE, CqlParserK_LIMIT, CqlParserK_LIST, CqlParserK_LOGIN, CqlParserK_MAP, CqlParserK_MODIFY, CqlParserK_NAN, CqlParserK_NOLOGIN, CqlParserK_NORECURSIVE, CqlParserK_NOSUPERUSER, CqlParserK_NOT, CqlParserK_NULL, CqlParserK_OF, CqlParserK_ON, CqlParserK_OPTIONS, CqlParserK_OR, CqlParserK_ORDER, CqlParserK_PASSWORD, CqlParserK_PERMISSION, CqlParserK_PERMISSIONS, CqlParserK_PRIMARY, CqlParserK_RENAME, CqlParserK_REPLACE, CqlParserK_RETURNS, CqlParserK_REVOKE, CqlParserK_ROLE, CqlParserK_ROLES, CqlParserK_SCHEMA, CqlParserK_SELECT, CqlParserK_SET, CqlParserK_SFUNC, CqlParserK_SMALLINT, CqlParserK_STATIC, CqlParserK_STORAGE, CqlParserK_STYPE, CqlParserK_SUPERUSER, CqlParserK_TABLE, CqlParserK_TEXT, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_TO, CqlParserK_TOKEN, CqlParserK_TRIGGER, CqlParserK_TRUNCATE, CqlParserK_TTL, CqlParserK_TUPLE, CqlParserK_TYPE, CqlParserK_UNLOGGED, CqlParserK_UPDATE, CqlParserK_USE, CqlParserK_USER, CqlParserK_USERS, CqlParserK_USING, CqlParserK_UUID, CqlParserK_VALUES, CqlParserK_VARCHAR, CqlParserK_VARINT, CqlParserK_WHERE, CqlParserK_WITH, CqlParserK_WRITETIME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.Keyword()
		}

	case CqlParserDQUOTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
			p.Match(CqlParserDQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserDOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(172)
			p.Match(CqlParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserCOMMA:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(173)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserL_PAREN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(174)
			p.Match(CqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserR_PAREN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(175)
			p.Match(CqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserL_ANGLE_BRACKET:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(176)
			p.Match(CqlParserL_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserR_ANGLE_BRACKET:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(177)
			p.Match(CqlParserR_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(178)
			p.Match(CqlParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserIDENTIFIER_WITH_HYPHEN:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(179)
			p.Match(CqlParserIDENTIFIER_WITH_HYPHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserUNKNOWN:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(180)
			p.Match(CqlParserUNKNOWN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ReservedKeyword() IReservedKeywordContext
	NonReservedKeyword() INonReservedKeywordContext

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_keyword
	return p
}

func InitEmptyKeywordContext(p *KeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_keyword
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) ReservedKeyword() IReservedKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReservedKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReservedKeywordContext)
}

func (s *KeywordContext) NonReservedKeyword() INonReservedKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INonReservedKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INonReservedKeywordContext)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *CqlParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CqlParserRULE_keyword)
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_ADD, CqlParserK_ALLOW, CqlParserK_ALTER, CqlParserK_AND, CqlParserK_APPLY, CqlParserK_ASC, CqlParserK_AUTHORIZE, CqlParserK_BATCH, CqlParserK_BEGIN, CqlParserK_BY, CqlParserK_COLUMNFAMILY, CqlParserK_CREATE, CqlParserK_DELETE, CqlParserK_DESC, CqlParserK_DESCRIBE, CqlParserK_DROP, CqlParserK_ENTRIES, CqlParserK_EXECUTE, CqlParserK_FROM, CqlParserK_FULL, CqlParserK_GRANT, CqlParserK_IF, CqlParserK_IN, CqlParserK_INDEX, CqlParserK_INFINITY, CqlParserK_INSERT, CqlParserK_INTO, CqlParserK_KEYSPACE, CqlParserK_LIMIT, CqlParserK_MODIFY, CqlParserK_NAN, CqlParserK_NORECURSIVE, CqlParserK_NOT, CqlParserK_NULL, CqlParserK_OF, CqlParserK_ON, CqlParserK_OR, CqlParserK_ORDER, CqlParserK_PRIMARY, CqlParserK_RENAME, CqlParserK_REPLACE, CqlParserK_REVOKE, CqlParserK_SCHEMA, CqlParserK_SELECT, CqlParserK_SET, CqlParserK_TABLE, CqlParserK_TO, CqlParserK_TOKEN, CqlParserK_TRUNCATE, CqlParserK_UNLOGGED, CqlParserK_UPDATE, CqlParserK_USE, CqlParserK_USING, CqlParserK_WHERE, CqlParserK_WITH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.ReservedKeyword()
		}

	case CqlParserK_AGGREGATE, CqlParserK_ALL, CqlParserK_AS, CqlParserK_ASCII, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_CALLED, CqlParserK_CLUSTERING, CqlParserK_COMPACT, CqlParserK_CONTAINS, CqlParserK_COUNT, CqlParserK_COUNTER, CqlParserK_CUSTOM, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DISTINCT, CqlParserK_DOUBLE, CqlParserK_EXISTS, CqlParserK_FILTERING, CqlParserK_FINALFUNC, CqlParserK_FLOAT, CqlParserK_FROZEN, CqlParserK_FUNCTION, CqlParserK_FUNCTIONS, CqlParserK_INET, CqlParserK_INITCOND, CqlParserK_INPUT, CqlParserK_INT, CqlParserK_JSON, CqlParserK_KEY, CqlParserK_KEYS, CqlParserK_KEYSPACES, CqlParserK_LANGUAGE, CqlParserK_LIST, CqlParserK_LOGIN, CqlParserK_MAP, CqlParserK_NOLOGIN, CqlParserK_NOSUPERUSER, CqlParserK_OPTIONS, CqlParserK_PASSWORD, CqlParserK_PERMISSION, CqlParserK_PERMISSIONS, CqlParserK_RETURNS, CqlParserK_ROLE, CqlParserK_ROLES, CqlParserK_SFUNC, CqlParserK_SMALLINT, CqlParserK_STATIC, CqlParserK_STORAGE, CqlParserK_STYPE, CqlParserK_SUPERUSER, CqlParserK_TEXT, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_TRIGGER, CqlParserK_TTL, CqlParserK_TUPLE, CqlParserK_TYPE, CqlParserK_USER, CqlParserK_USERS, CqlParserK_UUID, CqlParserK_VALUES, CqlParserK_VARCHAR, CqlParserK_VARINT, CqlParserK_WRITETIME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.NonReservedKeyword()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReservedKeywordContext is an interface to support dynamic dispatch.
type IReservedKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_ADD() antlr.TerminalNode
	K_ALLOW() antlr.TerminalNode
	K_ALTER() antlr.TerminalNode
	K_AND() antlr.TerminalNode
	K_APPLY() antlr.TerminalNode
	K_ASC() antlr.TerminalNode
	K_AUTHORIZE() antlr.TerminalNode
	K_BATCH() antlr.TerminalNode
	K_BEGIN() antlr.TerminalNode
	K_BY() antlr.TerminalNode
	K_COLUMNFAMILY() antlr.TerminalNode
	K_CREATE() antlr.TerminalNode
	K_DELETE() antlr.TerminalNode
	K_DESC() antlr.TerminalNode
	K_DESCRIBE() antlr.TerminalNode
	K_DROP() antlr.TerminalNode
	K_ENTRIES() antlr.TerminalNode
	K_EXECUTE() antlr.TerminalNode
	K_FROM() antlr.TerminalNode
	K_FULL() antlr.TerminalNode
	K_GRANT() antlr.TerminalNode
	K_IF() antlr.TerminalNode
	K_IN() antlr.TerminalNode
	K_INDEX() antlr.TerminalNode
	K_INFINITY() antlr.TerminalNode
	K_INSERT() antlr.TerminalNode
	K_INTO() antlr.TerminalNode
	K_KEYSPACE() antlr.TerminalNode
	K_LIMIT() antlr.TerminalNode
	K_MODIFY() antlr.TerminalNode
	K_NAN() antlr.TerminalNode
	K_NORECURSIVE() antlr.TerminalNode
	K_NOT() antlr.TerminalNode
	K_NULL() antlr.TerminalNode
	K_OF() antlr.TerminalNode
	K_ON() antlr.TerminalNode
	K_OR() antlr.TerminalNode
	K_ORDER() antlr.TerminalNode
	K_PRIMARY() antlr.TerminalNode
	K_RENAME() antlr.TerminalNode
	K_REPLACE() antlr.TerminalNode
	K_REVOKE() antlr.TerminalNode
	K_SCHEMA() antlr.TerminalNode
	K_SELECT() antlr.TerminalNode
	K_SET() antlr.TerminalNode
	K_TABLE() antlr.TerminalNode
	K_TO() antlr.TerminalNode
	K_TOKEN() antlr.TerminalNode
	K_TRUNCATE() antlr.TerminalNode
	K_UNLOGGED() antlr.TerminalNode
	K_UPDATE() antlr.TerminalNode
	K_USE() antlr.TerminalNode
	K_USING() antlr.TerminalNode
	K_WHERE() antlr.TerminalNode
	K_WITH() antlr.TerminalNode

	// IsReservedKeywordContext differentiates from other interfaces.
	IsReservedKeywordContext()
}

type ReservedKeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedKeywordContext() *ReservedKeywordContext {
	var p = new(ReservedKeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_reservedKeyword
	return p
}

func InitEmptyReservedKeywordContext(p *ReservedKeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_reservedKeyword
}

func (*ReservedKeywordContext) IsReservedKeywordContext() {}

func NewReservedKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedKeywordContext {
	var p = new(ReservedKeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_reservedKeyword

	return p
}

func (s *ReservedKeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedKeywordContext) K_ADD() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ADD, 0)
}

func (s *ReservedKeywordContext) K_ALLOW() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ALLOW, 0)
}

func (s *ReservedKeywordContext) K_ALTER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ALTER, 0)
}

func (s *ReservedKeywordContext) K_AND() antlr.TerminalNode {
	return s.GetToken(CqlParserK_AND, 0)
}

func (s *ReservedKeywordContext) K_APPLY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_APPLY, 0)
}

func (s *ReservedKeywordContext) K_ASC() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ASC, 0)
}

func (s *ReservedKeywordContext) K_AUTHORIZE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_AUTHORIZE, 0)
}

func (s *ReservedKeywordContext) K_BATCH() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BATCH, 0)
}

func (s *ReservedKeywordContext) K_BEGIN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BEGIN, 0)
}

func (s *ReservedKeywordContext) K_BY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BY, 0)
}

func (s *ReservedKeywordContext) K_COLUMNFAMILY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_COLUMNFAMILY, 0)
}

func (s *ReservedKeywordContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CREATE, 0)
}

func (s *ReservedKeywordContext) K_DELETE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DELETE, 0)
}

func (s *ReservedKeywordContext) K_DESC() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DESC, 0)
}

func (s *ReservedKeywordContext) K_DESCRIBE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DESCRIBE, 0)
}

func (s *ReservedKeywordContext) K_DROP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DROP, 0)
}

func (s *ReservedKeywordContext) K_ENTRIES() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ENTRIES, 0)
}

func (s *ReservedKeywordContext) K_EXECUTE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_EXECUTE, 0)
}

func (s *ReservedKeywordContext) K_FROM() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FROM, 0)
}

func (s *ReservedKeywordContext) K_FULL() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FULL, 0)
}

func (s *ReservedKeywordContext) K_GRANT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_GRANT, 0)
}

func (s *ReservedKeywordContext) K_IF() antlr.TerminalNode {
	return s.GetToken(CqlParserK_IF, 0)
}

func (s *ReservedKeywordContext) K_IN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_IN, 0)
}

func (s *ReservedKeywordContext) K_INDEX() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INDEX, 0)
}

func (s *ReservedKeywordContext) K_INFINITY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INFINITY, 0)
}

func (s *ReservedKeywordContext) K_INSERT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INSERT, 0)
}

func (s *ReservedKeywordContext) K_INTO() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INTO, 0)
}

func (s *ReservedKeywordContext) K_KEYSPACE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_KEYSPACE, 0)
}

func (s *ReservedKeywordContext) K_LIMIT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_LIMIT, 0)
}

func (s *ReservedKeywordContext) K_MODIFY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_MODIFY, 0)
}

func (s *ReservedKeywordContext) K_NAN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NAN, 0)
}

func (s *ReservedKeywordContext) K_NORECURSIVE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NORECURSIVE, 0)
}

func (s *ReservedKeywordContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NOT, 0)
}

func (s *ReservedKeywordContext) K_NULL() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NULL, 0)
}

func (s *ReservedKeywordContext) K_OF() antlr.TerminalNode {
	return s.GetToken(CqlParserK_OF, 0)
}

func (s *ReservedKeywordContext) K_ON() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ON, 0)
}

func (s *ReservedKeywordContext) K_OR() antlr.TerminalNode {
	return s.GetToken(CqlParserK_OR, 0)
}

func (s *ReservedKeywordContext) K_ORDER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ORDER, 0)
}

func (s *ReservedKeywordContext) K_PRIMARY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_PRIMARY, 0)
}

func (s *ReservedKeywordContext) K_RENAME() antlr.TerminalNode {
	return s.GetToken(CqlParserK_RENAME, 0)
}

func (s *ReservedKeywordContext) K_REPLACE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_REPLACE, 0)
}

func (s *ReservedKeywordContext) K_REVOKE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_REVOKE, 0)
}

func (s *ReservedKeywordContext) K_SCHEMA() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SCHEMA, 0)
}

func (s *ReservedKeywordContext) K_SELECT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SELECT, 0)
}

func (s *ReservedKeywordContext) K_SET() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SET, 0)
}

func (s *ReservedKeywordContext) K_TABLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TABLE, 0)
}

func (s *ReservedKeywordContext) K_TO() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TO, 0)
}

func (s *ReservedKeywordContext) K_TOKEN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TOKEN, 0)
}

func (s *ReservedKeywordContext) K_TRUNCATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TRUNCATE, 0)
}

func (s *ReservedKeywordContext) K_UNLOGGED() antlr.TerminalNode {
	return s.GetToken(CqlParserK_UNLOGGED, 0)
}

func (s *ReservedKeywordContext) K_UPDATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_UPDATE, 0)
}

func (s *ReservedKeywordContext) K_USE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_USE, 0)
}

func (s *ReservedKeywordContext) K_USING() antlr.TerminalNode {
	return s.GetToken(CqlParserK_USING, 0)
}

func (s *ReservedKeywordContext) K_WHERE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_WHERE, 0)
}

func (s *ReservedKeywordContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(CqlParserK_WITH, 0)
}

func (s *ReservedKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterReservedKeyword(s)
	}
}

func (s *ReservedKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitReservedKeyword(s)
	}
}

func (p *CqlParser) ReservedKeyword() (localctx IReservedKeywordContext) {
	localctx = NewReservedKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CqlParserRULE_reservedKeyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-7977060089555961102) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&54730804790267371) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INonReservedKeywordContext is an interface to support dynamic dispatch.
type INonReservedKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_AGGREGATE() antlr.TerminalNode
	K_ALL() antlr.TerminalNode
	K_AS() antlr.TerminalNode
	K_ASCII() antlr.TerminalNode
	K_BIGINT() antlr.TerminalNode
	K_BLOB() antlr.TerminalNode
	K_BOOLEAN() antlr.TerminalNode
	K_CALLED() antlr.TerminalNode
	K_CLUSTERING() antlr.TerminalNode
	K_COMPACT() antlr.TerminalNode
	K_CONTAINS() antlr.TerminalNode
	K_COUNT() antlr.TerminalNode
	K_COUNTER() antlr.TerminalNode
	K_CUSTOM() antlr.TerminalNode
	K_DATE() antlr.TerminalNode
	K_DECIMAL() antlr.TerminalNode
	K_DISTINCT() antlr.TerminalNode
	K_DOUBLE() antlr.TerminalNode
	K_EXISTS() antlr.TerminalNode
	K_FILTERING() antlr.TerminalNode
	K_FINALFUNC() antlr.TerminalNode
	K_FLOAT() antlr.TerminalNode
	K_FROZEN() antlr.TerminalNode
	K_FUNCTION() antlr.TerminalNode
	K_FUNCTIONS() antlr.TerminalNode
	K_INET() antlr.TerminalNode
	K_INITCOND() antlr.TerminalNode
	K_INPUT() antlr.TerminalNode
	K_INT() antlr.TerminalNode
	K_JSON() antlr.TerminalNode
	K_KEY() antlr.TerminalNode
	K_KEYS() antlr.TerminalNode
	K_KEYSPACES() antlr.TerminalNode
	K_LANGUAGE() antlr.TerminalNode
	K_LIST() antlr.TerminalNode
	K_LOGIN() antlr.TerminalNode
	K_MAP() antlr.TerminalNode
	K_NOLOGIN() antlr.TerminalNode
	K_NOSUPERUSER() antlr.TerminalNode
	K_OPTIONS() antlr.TerminalNode
	K_PASSWORD() antlr.TerminalNode
	K_PERMISSION() antlr.TerminalNode
	K_PERMISSIONS() antlr.TerminalNode
	K_RETURNS() antlr.TerminalNode
	K_ROLE() antlr.TerminalNode
	K_ROLES() antlr.TerminalNode
	K_SFUNC() antlr.TerminalNode
	K_SMALLINT() antlr.TerminalNode
	K_STATIC() antlr.TerminalNode
	K_STORAGE() antlr.TerminalNode
	K_STYPE() antlr.TerminalNode
	K_SUPERUSER() antlr.TerminalNode
	K_TEXT() antlr.TerminalNode
	K_TIME() antlr.TerminalNode
	K_TIMESTAMP() antlr.TerminalNode
	K_TIMEUUID() antlr.TerminalNode
	K_TINYINT() antlr.TerminalNode
	K_TRIGGER() antlr.TerminalNode
	K_TTL() antlr.TerminalNode
	K_TUPLE() antlr.TerminalNode
	K_TYPE() antlr.TerminalNode
	K_USER() antlr.TerminalNode
	K_USERS() antlr.TerminalNode
	K_UUID() antlr.TerminalNode
	K_VALUES() antlr.TerminalNode
	K_VARCHAR() antlr.TerminalNode
	K_VARINT() antlr.TerminalNode
	K_WRITETIME() antlr.TerminalNode

	// IsNonReservedKeywordContext differentiates from other interfaces.
	IsNonReservedKeywordContext()
}

type NonReservedKeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonReservedKeywordContext() *NonReservedKeywordContext {
	var p = new(NonReservedKeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_nonReservedKeyword
	return p
}

func InitEmptyNonReservedKeywordContext(p *NonReservedKeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_nonReservedKeyword
}

func (*NonReservedKeywordContext) IsNonReservedKeywordContext() {}

func NewNonReservedKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonReservedKeywordContext {
	var p = new(NonReservedKeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_nonReservedKeyword

	return p
}

func (s *NonReservedKeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *NonReservedKeywordContext) K_AGGREGATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_AGGREGATE, 0)
}

func (s *NonReservedKeywordContext) K_ALL() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ALL, 0)
}

func (s *NonReservedKeywordContext) K_AS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_AS, 0)
}

func (s *NonReservedKeywordContext) K_ASCII() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ASCII, 0)
}

func (s *NonReservedKeywordContext) K_BIGINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BIGINT, 0)
}

func (s *NonReservedKeywordContext) K_BLOB() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BLOB, 0)
}

func (s *NonReservedKeywordContext) K_BOOLEAN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BOOLEAN, 0)
}

func (s *NonReservedKeywordContext) K_CALLED() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CALLED, 0)
}

func (s *NonReservedKeywordContext) K_CLUSTERING() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CLUSTERING, 0)
}

func (s *NonReservedKeywordContext) K_COMPACT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_COMPACT, 0)
}

func (s *NonReservedKeywordContext) K_CONTAINS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CONTAINS, 0)
}

func (s *NonReservedKeywordContext) K_COUNT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_COUNT, 0)
}

func (s *NonReservedKeywordContext) K_COUNTER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_COUNTER, 0)
}

func (s *NonReservedKeywordContext) K_CUSTOM() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CUSTOM, 0)
}

func (s *NonReservedKeywordContext) K_DATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DATE, 0)
}

func (s *NonReservedKeywordContext) K_DECIMAL() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DECIMAL, 0)
}

func (s *NonReservedKeywordContext) K_DISTINCT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DISTINCT, 0)
}

func (s *NonReservedKeywordContext) K_DOUBLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DOUBLE, 0)
}

func (s *NonReservedKeywordContext) K_EXISTS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_EXISTS, 0)
}

func (s *NonReservedKeywordContext) K_FILTERING() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FILTERING, 0)
}

func (s *NonReservedKeywordContext) K_FINALFUNC() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FINALFUNC, 0)
}

func (s *NonReservedKeywordContext) K_FLOAT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FLOAT, 0)
}

func (s *NonReservedKeywordContext) K_FROZEN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FROZEN, 0)
}

func (s *NonReservedKeywordContext) K_FUNCTION() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FUNCTION, 0)
}

func (s *NonReservedKeywordContext) K_FUNCTIONS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FUNCTIONS, 0)
}

func (s *NonReservedKeywordContext) K_INET() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INET, 0)
}

func (s *NonReservedKeywordContext) K_INITCOND() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INITCOND, 0)
}

func (s *NonReservedKeywordContext) K_INPUT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INPUT, 0)
}

func (s *NonReservedKeywordContext) K_INT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INT, 0)
}

func (s *NonReservedKeywordContext) K_JSON() antlr.TerminalNode {
	return s.GetToken(CqlParserK_JSON, 0)
}

func (s *NonReservedKeywordContext) K_KEY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_KEY, 0)
}

func (s *NonReservedKeywordContext) K_KEYS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_KEYS, 0)
}

func (s *NonReservedKeywordContext) K_KEYSPACES() antlr.TerminalNode {
	return s.GetToken(CqlParserK_KEYSPACES, 0)
}

func (s *NonReservedKeywordContext) K_LANGUAGE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_LANGUAGE, 0)
}

func (s *NonReservedKeywordContext) K_LIST() antlr.TerminalNode {
	return s.GetToken(CqlParserK_LIST, 0)
}

func (s *NonReservedKeywordContext) K_LOGIN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_LOGIN, 0)
}

func (s *NonReservedKeywordContext) K_MAP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_MAP, 0)
}

func (s *NonReservedKeywordContext) K_NOLOGIN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NOLOGIN, 0)
}

func (s *NonReservedKeywordContext) K_NOSUPERUSER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NOSUPERUSER, 0)
}

func (s *NonReservedKeywordContext) K_OPTIONS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_OPTIONS, 0)
}

func (s *NonReservedKeywordContext) K_PASSWORD() antlr.TerminalNode {
	return s.GetToken(CqlParserK_PASSWORD, 0)
}

func (s *NonReservedKeywordContext) K_PERMISSION() antlr.TerminalNode {
	return s.GetToken(CqlParserK_PERMISSION, 0)
}

func (s *NonReservedKeywordContext) K_PERMISSIONS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_PERMISSIONS, 0)
}

func (s *NonReservedKeywordContext) K_RETURNS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_RETURNS, 0)
}

func (s *NonReservedKeywordContext) K_ROLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ROLE, 0)
}

func (s *NonReservedKeywordContext) K_ROLES() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ROLES, 0)
}

func (s *NonReservedKeywordContext) K_SFUNC() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SFUNC, 0)
}

func (s *NonReservedKeywordContext) K_SMALLINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SMALLINT, 0)
}

func (s *NonReservedKeywordContext) K_STATIC() antlr.TerminalNode {
	return s.GetToken(CqlParserK_STATIC, 0)
}

func (s *NonReservedKeywordContext) K_STORAGE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_STORAGE, 0)
}

func (s *NonReservedKeywordContext) K_STYPE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_STYPE, 0)
}

func (s *NonReservedKeywordContext) K_SUPERUSER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SUPERUSER, 0)
}

func (s *NonReservedKeywordContext) K_TEXT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TEXT, 0)
}

func (s *NonReservedKeywordContext) K_TIME() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIME, 0)
}

func (s *NonReservedKeywordContext) K_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIMESTAMP, 0)
}

func (s *NonReservedKeywordContext) K_TIMEUUID() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIMEUUID, 0)
}

func (s *NonReservedKeywordContext) K_TINYINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TINYINT, 0)
}

func (s *NonReservedKeywordContext) K_TRIGGER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TRIGGER, 0)
}

func (s *NonReservedKeywordContext) K_TTL() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TTL, 0)
}

func (s *NonReservedKeywordContext) K_TUPLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TUPLE, 0)
}

func (s *NonReservedKeywordContext) K_TYPE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TYPE, 0)
}

func (s *NonReservedKeywordContext) K_USER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_USER, 0)
}

func (s *NonReservedKeywordContext) K_USERS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_USERS, 0)
}

func (s *NonReservedKeywordContext) K_UUID() antlr.TerminalNode {
	return s.GetToken(CqlParserK_UUID, 0)
}

func (s *NonReservedKeywordContext) K_VALUES() antlr.TerminalNode {
	return s.GetToken(CqlParserK_VALUES, 0)
}

func (s *NonReservedKeywordContext) K_VARCHAR() antlr.TerminalNode {
	return s.GetToken(CqlParserK_VARCHAR, 0)
}

func (s *NonReservedKeywordContext) K_VARINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_VARINT, 0)
}

func (s *NonReservedKeywordContext) K_WRITETIME() antlr.TerminalNode {
	return s.GetToken(CqlParserK_WRITETIME, 0)
}

func (s *NonReservedKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonReservedKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonReservedKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterNonReservedKeyword(s)
	}
}

func (s *NonReservedKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitNonReservedKeyword(s)
	}
}

func (p *CqlParser) NonReservedKeyword() (localctx INonReservedKeywordContext) {
	localctx = NewNonReservedKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CqlParserRULE_nonReservedKeyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7977060089555961100) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&715075066284708007) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
