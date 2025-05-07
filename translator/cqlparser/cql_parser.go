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
		"'ANY'", "'APPLY'", "'ARRAY'", "'AS'", "'ASC'", "'ASCII'", "'ASSERT_ROWS_MODIFIED'",
		"'AT'", "'AUTHORIZE'", "'BATCH'", "'BEGIN'", "'BETWEEN'", "'BIGINT'",
		"'BLOB'", "'BOOLEAN'", "'BY'", "'CALLED'", "'CASE'", "'CAST'", "'CLUSTERING'",
		"'COLLATE'", "'COLUMNFAMILY'", "'COMPACT'", "'CONTAINS'", "'COUNT'",
		"'COUNTER'", "'CREATE'", "'CROSS'", "'CUBE'", "'CURRENT'", "'CUSTOM'",
		"'DATE'", "'DECIMAL'", "'DEFAULT'", "'DEFINE'", "'DELETE'", "'DESC'",
		"'DESCRIBE'", "'DISTINCT'", "'DOUBLE'", "'DROP'", "'ELSE'", "'END'",
		"'ENTRIES'", "'ENUM'", "'ESCAPE'", "'EXCEPT'", "'EXCLUDE'", "'EXECUTE'",
		"'EXISTS'", "'EXTRACT'", "'FALSE'", "'FETCH'", "'FILTERING'", "'FINALFUNC'",
		"'FLOAT'", "'FOLLOWING'", "'FOR'", "'FROM'", "'FROZEN'", "'FULL'", "'FUNCTION'",
		"'FUNCTIONS'", "'GRANT'", "'GROUP'", "'GROUPING'", "'GROUPS'", "'HASH'",
		"'HAVING'", "'IF'", "'IGNORE'", "'IN'", "'INDEX'", "'INET'", "'INFINITY'",
		"'INITCOND'", "'INNER'", "'INPUT'", "'INSERT'", "'INT'", "'INTERSECT'",
		"'INTERVAL'", "'INTO'", "'IS'", "'JOIN'", "'JSON'", "'KEY'", "'KEYS'",
		"'KEYSPACE'", "'KEYSPACES'", "'LANGUAGE'", "'LATERAL'", "'LEFT'", "'LIKE'",
		"'LIMIT'", "'LIST'", "'LOGIN'", "'LOOKUP'", "'MAP'", "'MERGE'", "'MODIFY'",
		"'NAN'", "'NATURAL'", "'NEW'", "'NO'", "'NOLOGIN'", "'NORECURSIVE'",
		"'NOSUPERUSER'", "'NOT'", "'NULL'", "'NULLS'", "'OF'", "'ON'", "'OPTIONS'",
		"'OR'", "'ORDER'", "'OUTER'", "'OVER'", "'PARTITION'", "'PASSWORD'",
		"'PERMISSION'", "'PERMISSIONS'", "'PRECEDING'", "'PRIMARY'", "'PROTO'",
		"'QUALIFY'", "'RANGE'", "'RECURSIVE'", "'RENAME'", "'REPLACE'", "'RESPECT'",
		"'RETURNS'", "'REVOKE'", "'RIGHT'", "'ROLE'", "'ROLES'", "'ROLLUP'",
		"'ROWS'", "'SCHEMA'", "'SELECT'", "'SET'", "'SFUNC'", "'SMALLINT'",
		"'SOME'", "'STATIC'", "'STORAGE'", "'STRUCT'", "'STYPE'", "'SUPERUSER'",
		"'TABLE'", "'TABLESAMPLE'", "'TEXT'", "'THEN'", "'TIME'", "'TIMESTAMP'",
		"'TIMEUUID'", "'TINYINT'", "'TO'", "'TOKEN'", "'TREAT'", "'TRIGGER'",
		"'TRUE'", "'TRUNCATE'", "'TTL'", "'TUPLE'", "'TYPE'", "'UNBOUNDED'",
		"'UNION'", "'UNLOGGED'", "'UNNEST'", "'UPDATE'", "'USE'", "'USER'",
		"'USERS'", "'USING'", "'UUID'", "'VALUES'", "'VARCHAR'", "'VARINT'",
		"'WHEN'", "'WHERE'", "'WINDOW'", "'WITH'", "'WITHIN'", "'WRITETIME'",
		"'MATERIALIZED'", "'VIEW'", "';'", "'''", "'\"'", "'.'", "','", "'('",
		"')'", "'<'", "'>'",
	}
	staticData.SymbolicNames = []string{
		"", "K_ADD", "K_AGGREGATE", "K_ALL", "K_ALLOW", "K_ALTER", "K_AND",
		"K_ANY", "K_APPLY", "K_ARRAY", "K_AS", "K_ASC", "K_ASCII", "K_ASSERT_ROWS_MODIFIED",
		"K_AT", "K_AUTHORIZE", "K_BATCH", "K_BEGIN", "K_BETWEEN", "K_BIGINT",
		"K_BLOB", "K_BOOLEAN", "K_BY", "K_CALLED", "K_CASE", "K_CAST", "K_CLUSTERING",
		"K_COLLATE", "K_COLUMNFAMILY", "K_COMPACT", "K_CONTAINS", "K_COUNT",
		"K_COUNTER", "K_CREATE", "K_CROSS", "K_CUBE", "K_CURRENT", "K_CUSTOM",
		"K_DATE", "K_DECIMAL", "K_DEFAULT", "K_DEFINE", "K_DELETE", "K_DESC",
		"K_DESCRIBE", "K_DISTINCT", "K_DOUBLE", "K_DROP", "K_ELSE", "K_END",
		"K_ENTRIES", "K_ENUM", "K_ESCAPE", "K_EXCEPT", "K_EXCLUDE", "K_EXECUTE",
		"K_EXISTS", "K_EXTRACT", "K_FALSE", "K_FETCH", "K_FILTERING", "K_FINALFUNC",
		"K_FLOAT", "K_FOLLOWING", "K_FOR", "K_FROM", "K_FROZEN", "K_FULL", "K_FUNCTION",
		"K_FUNCTIONS", "K_GRANT", "K_GROUP", "K_GROUPING", "K_GROUPS", "K_HASH",
		"K_HAVING", "K_IF", "K_IGNORE", "K_IN", "K_INDEX", "K_INET", "K_INFINITY",
		"K_INITCOND", "K_INNER", "K_INPUT", "K_INSERT", "K_INT", "K_INTERSECT",
		"K_INTERVAL", "K_INTO", "K_IS", "K_JOIN", "K_JSON", "K_KEY", "K_KEYS",
		"K_KEYSPACE", "K_KEYSPACES", "K_LANGUAGE", "K_LATERAL", "K_LEFT", "K_LIKE",
		"K_LIMIT", "K_LIST", "K_LOGIN", "K_LOOKUP", "K_MAP", "K_MERGE", "K_MODIFY",
		"K_NAN", "K_NATURAL", "K_NEW", "K_NO", "K_NOLOGIN", "K_NORECURSIVE",
		"K_NOSUPERUSER", "K_NOT", "K_NULL", "K_NULLS", "K_OF", "K_ON", "K_OPTIONS",
		"K_OR", "K_ORDER", "K_OUTER", "K_OVER", "K_PARTITION", "K_PASSWORD",
		"K_PERMISSION", "K_PERMISSIONS", "K_PRECEDING", "K_PRIMARY", "K_PROTO",
		"K_QUALIFY", "K_RANGE", "K_RECURSIVE", "K_RENAME", "K_REPLACE", "K_RESPECT",
		"K_RETURNS", "K_REVOKE", "K_RIGHT", "K_ROLE", "K_ROLES", "K_ROLLUP",
		"K_ROWS", "K_SCHEMA", "K_SELECT", "K_SET", "K_SFUNC", "K_SMALLINT",
		"K_SOME", "K_STATIC", "K_STORAGE", "K_STRUCT", "K_STYPE", "K_SUPERUSER",
		"K_TABLE", "K_TABLESAMPLE", "K_TEXT", "K_THEN", "K_TIME", "K_TIMESTAMP",
		"K_TIMEUUID", "K_TINYINT", "K_TO", "K_TOKEN", "K_TREAT", "K_TRIGGER",
		"K_TRUE", "K_TRUNCATE", "K_TTL", "K_TUPLE", "K_TYPE", "K_UNBOUNDED",
		"K_UNION", "K_UNLOGGED", "K_UNNEST", "K_UPDATE", "K_USE", "K_USER",
		"K_USERS", "K_USING", "K_UUID", "K_VALUES", "K_VARCHAR", "K_VARINT",
		"K_WHEN", "K_WHERE", "K_WINDOW", "K_WITH", "K_WITHIN", "K_WRITETIME",
		"K_MATERIALIZED", "K_VIEW", "SEMICOLON", "SQUOTE", "DQUOTE", "DOT",
		"COMMA", "L_PAREN", "R_PAREN", "L_ANGLE_BRACKET", "R_ANGLE_BRACKET",
		"IDENTIFIER", "IDENTIFIER_WITH_HYPHEN", "WS", "COMMENT", "MULTILINE_COMMENT",
		"UNKNOWN",
	}
	staticData.RuleNames = []string{
		"root", "cqlStatement", "createTable", "unsupportedStatement", "columnDefinitionList",
		"columnDefinition", "primaryKeyClause", "primaryKey", "partitionKey",
		"clusteringColumns", "columnName", "tableName", "keyspaceName", "generalIdentifier",
		"identifier", "cqlType", "primaryKeyKeywords", "ifExists", "ifNotExists",
		"cqlNativeType", "cqlCollectionType", "wihtTableOptions", "nonSemicolonToken",
		"keyword", "reservedKeyword", "nonReservedKeyword",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 208, 545, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		1, 0, 3, 0, 55, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 61, 8, 1, 1, 2, 1,
		2, 1, 2, 3, 2, 66, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 73, 8, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 80, 8, 3, 1, 3, 1, 3, 1, 3, 5, 3, 85,
		8, 3, 10, 3, 12, 3, 88, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 95, 8,
		3, 10, 3, 12, 3, 98, 9, 3, 1, 3, 1, 3, 1, 3, 3, 3, 103, 8, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 5, 3, 110, 8, 3, 10, 3, 12, 3, 113, 9, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 118, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3, 123, 8, 3, 1, 3, 1, 3,
		1, 3, 5, 3, 128, 8, 3, 10, 3, 12, 3, 131, 9, 3, 1, 3, 1, 3, 1, 3, 5, 3,
		136, 8, 3, 10, 3, 12, 3, 139, 9, 3, 1, 3, 1, 3, 1, 3, 5, 3, 144, 8, 3,
		10, 3, 12, 3, 147, 9, 3, 1, 3, 1, 3, 5, 3, 151, 8, 3, 10, 3, 12, 3, 154,
		9, 3, 1, 3, 1, 3, 3, 3, 158, 8, 3, 1, 3, 1, 3, 5, 3, 162, 8, 3, 10, 3,
		12, 3, 165, 9, 3, 1, 3, 4, 3, 168, 8, 3, 11, 3, 12, 3, 169, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 3, 3, 177, 8, 3, 1, 3, 1, 3, 5, 3, 181, 8, 3, 10, 3,
		12, 3, 184, 9, 3, 1, 3, 1, 3, 1, 3, 3, 3, 189, 8, 3, 1, 3, 5, 3, 192, 8,
		3, 10, 3, 12, 3, 195, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 201, 8, 3, 1,
		3, 5, 3, 204, 8, 3, 10, 3, 12, 3, 207, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5,
		3, 213, 8, 3, 10, 3, 12, 3, 216, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 222,
		8, 3, 1, 3, 5, 3, 225, 8, 3, 10, 3, 12, 3, 228, 9, 3, 1, 3, 1, 3, 1, 3,
		3, 3, 233, 8, 3, 1, 3, 5, 3, 236, 8, 3, 10, 3, 12, 3, 239, 9, 3, 1, 3,
		1, 3, 1, 3, 5, 3, 244, 8, 3, 10, 3, 12, 3, 247, 9, 3, 1, 3, 1, 3, 1, 3,
		5, 3, 252, 8, 3, 10, 3, 12, 3, 255, 9, 3, 1, 3, 1, 3, 5, 3, 259, 8, 3,
		10, 3, 12, 3, 262, 9, 3, 1, 3, 1, 3, 5, 3, 266, 8, 3, 10, 3, 12, 3, 269,
		9, 3, 1, 3, 1, 3, 5, 3, 273, 8, 3, 10, 3, 12, 3, 276, 9, 3, 1, 3, 1, 3,
		1, 3, 5, 3, 281, 8, 3, 10, 3, 12, 3, 284, 9, 3, 1, 3, 1, 3, 1, 3, 5, 3,
		289, 8, 3, 10, 3, 12, 3, 292, 9, 3, 1, 3, 1, 3, 1, 3, 5, 3, 297, 8, 3,
		10, 3, 12, 3, 300, 9, 3, 1, 3, 1, 3, 1, 3, 3, 3, 305, 8, 3, 1, 3, 1, 3,
		3, 3, 309, 8, 3, 1, 3, 5, 3, 312, 8, 3, 10, 3, 12, 3, 315, 9, 3, 1, 3,
		1, 3, 1, 3, 5, 3, 320, 8, 3, 10, 3, 12, 3, 323, 9, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 330, 8, 3, 1, 3, 5, 3, 333, 8, 3, 10, 3, 12, 3, 336,
		9, 3, 1, 3, 1, 3, 1, 3, 3, 3, 341, 8, 3, 1, 3, 1, 3, 3, 3, 345, 8, 3, 1,
		3, 5, 3, 348, 8, 3, 10, 3, 12, 3, 351, 9, 3, 1, 3, 1, 3, 1, 3, 3, 3, 356,
		8, 3, 1, 3, 5, 3, 359, 8, 3, 10, 3, 12, 3, 362, 9, 3, 1, 3, 1, 3, 1, 3,
		3, 3, 367, 8, 3, 1, 3, 5, 3, 370, 8, 3, 10, 3, 12, 3, 373, 9, 3, 1, 3,
		1, 3, 1, 3, 5, 3, 378, 8, 3, 10, 3, 12, 3, 381, 9, 3, 1, 3, 1, 3, 1, 3,
		3, 3, 386, 8, 3, 1, 3, 5, 3, 389, 8, 3, 10, 3, 12, 3, 392, 9, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 397, 8, 3, 1, 3, 5, 3, 400, 8, 3, 10, 3, 12, 3, 403,
		9, 3, 1, 3, 1, 3, 1, 3, 3, 3, 408, 8, 3, 1, 3, 5, 3, 411, 8, 3, 10, 3,
		12, 3, 414, 9, 3, 3, 3, 416, 8, 3, 1, 4, 1, 4, 1, 4, 5, 4, 421, 8, 4, 10,
		4, 12, 4, 424, 9, 4, 1, 4, 1, 4, 3, 4, 428, 8, 4, 1, 5, 1, 5, 1, 5, 3,
		5, 433, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 443,
		8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 450, 8, 8, 10, 8, 12, 8, 453,
		9, 8, 1, 8, 1, 8, 3, 8, 457, 8, 8, 1, 9, 1, 9, 1, 9, 5, 9, 462, 8, 9, 10,
		9, 12, 9, 465, 9, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 3, 11, 472, 8,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 3, 13, 480, 8, 13, 1, 14,
		1, 14, 3, 14, 484, 8, 14, 1, 15, 1, 15, 3, 15, 488, 8, 15, 1, 16, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 3, 20, 514, 8, 20, 1, 21, 1, 21, 5, 21, 518, 8, 21, 10, 21, 12,
		21, 521, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 535, 8, 22, 1, 23, 1, 23, 3, 23, 539,
		8, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 0, 0, 26, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 0, 5, 2, 0, 32, 32, 175, 175, 13, 0, 12, 12, 19, 21, 32, 32, 38,
		39, 46, 46, 62, 62, 80, 80, 86, 86, 149, 149, 158, 158, 160, 163, 182,
		182, 184, 185, 2, 0, 102, 102, 147, 147, 37, 0, 1, 1, 3, 11, 13, 18, 22,
		22, 24, 25, 27, 28, 30, 30, 33, 36, 40, 45, 47, 59, 63, 65, 67, 67, 70,
		79, 81, 81, 83, 83, 85, 85, 87, 91, 95, 95, 98, 101, 104, 104, 106, 111,
		113, 113, 115, 119, 121, 125, 129, 137, 139, 140, 143, 147, 150, 150, 153,
		153, 156, 157, 159, 159, 164, 166, 168, 169, 173, 178, 181, 181, 186, 190,
		192, 193, 36, 0, 2, 2, 12, 12, 19, 21, 23, 23, 26, 26, 29, 29, 31, 32,
		37, 39, 46, 46, 60, 62, 66, 66, 68, 69, 80, 80, 82, 82, 84, 84, 86, 86,
		92, 94, 96, 97, 102, 103, 105, 105, 112, 112, 114, 114, 120, 120, 126,
		128, 138, 138, 141, 142, 148, 149, 151, 152, 154, 155, 158, 158, 160, 163,
		167, 167, 170, 172, 179, 180, 182, 185, 191, 191, 634, 0, 52, 1, 0, 0,
		0, 2, 60, 1, 0, 0, 0, 4, 62, 1, 0, 0, 0, 6, 415, 1, 0, 0, 0, 8, 417, 1,
		0, 0, 0, 10, 429, 1, 0, 0, 0, 12, 434, 1, 0, 0, 0, 14, 439, 1, 0, 0, 0,
		16, 456, 1, 0, 0, 0, 18, 458, 1, 0, 0, 0, 20, 466, 1, 0, 0, 0, 22, 471,
		1, 0, 0, 0, 24, 475, 1, 0, 0, 0, 26, 479, 1, 0, 0, 0, 28, 483, 1, 0, 0,
		0, 30, 487, 1, 0, 0, 0, 32, 489, 1, 0, 0, 0, 34, 492, 1, 0, 0, 0, 36, 495,
		1, 0, 0, 0, 38, 499, 1, 0, 0, 0, 40, 513, 1, 0, 0, 0, 42, 515, 1, 0, 0,
		0, 44, 534, 1, 0, 0, 0, 46, 538, 1, 0, 0, 0, 48, 540, 1, 0, 0, 0, 50, 542,
		1, 0, 0, 0, 52, 54, 3, 2, 1, 0, 53, 55, 5, 194, 0, 0, 54, 53, 1, 0, 0,
		0, 54, 55, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 5, 0, 0, 1, 57, 1, 1,
		0, 0, 0, 58, 61, 3, 4, 2, 0, 59, 61, 3, 6, 3, 0, 60, 58, 1, 0, 0, 0, 60,
		59, 1, 0, 0, 0, 61, 3, 1, 0, 0, 0, 62, 63, 5, 33, 0, 0, 63, 65, 5, 156,
		0, 0, 64, 66, 3, 36, 18, 0, 65, 64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66,
		67, 1, 0, 0, 0, 67, 68, 3, 22, 11, 0, 68, 69, 5, 199, 0, 0, 69, 70, 3,
		8, 4, 0, 70, 72, 5, 200, 0, 0, 71, 73, 3, 42, 21, 0, 72, 71, 1, 0, 0, 0,
		72, 73, 1, 0, 0, 0, 73, 5, 1, 0, 0, 0, 74, 75, 5, 178, 0, 0, 75, 416, 3,
		24, 12, 0, 76, 77, 5, 33, 0, 0, 77, 79, 5, 95, 0, 0, 78, 80, 3, 36, 18,
		0, 79, 78, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 82,
		3, 24, 12, 0, 82, 86, 5, 189, 0, 0, 83, 85, 3, 44, 22, 0, 84, 83, 1, 0,
		0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 416,
		1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 90, 5, 5, 0, 0, 90, 91, 5, 95, 0, 0,
		91, 92, 3, 24, 12, 0, 92, 96, 5, 189, 0, 0, 93, 95, 3, 44, 22, 0, 94, 93,
		1, 0, 0, 0, 95, 98, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0,
		97, 416, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99, 100, 5, 47, 0, 0, 100, 102,
		5, 95, 0, 0, 101, 103, 3, 34, 17, 0, 102, 101, 1, 0, 0, 0, 102, 103, 1,
		0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 416, 3, 24, 12, 0, 105, 106, 5, 5,
		0, 0, 106, 107, 5, 156, 0, 0, 107, 111, 3, 22, 11, 0, 108, 110, 3, 44,
		22, 0, 109, 108, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0,
		111, 112, 1, 0, 0, 0, 112, 416, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114,
		115, 5, 47, 0, 0, 115, 117, 5, 156, 0, 0, 116, 118, 3, 34, 17, 0, 117,
		116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 416,
		3, 22, 11, 0, 120, 122, 5, 169, 0, 0, 121, 123, 5, 156, 0, 0, 122, 121,
		1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 416, 3, 22,
		11, 0, 125, 129, 5, 146, 0, 0, 126, 128, 3, 44, 22, 0, 127, 126, 1, 0,
		0, 0, 128, 131, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0,
		130, 416, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 132, 133, 5, 85, 0, 0, 133,
		137, 5, 89, 0, 0, 134, 136, 3, 44, 22, 0, 135, 134, 1, 0, 0, 0, 136, 139,
		1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 416, 1, 0,
		0, 0, 139, 137, 1, 0, 0, 0, 140, 141, 5, 177, 0, 0, 141, 145, 3, 22, 11,
		0, 142, 144, 3, 44, 22, 0, 143, 142, 1, 0, 0, 0, 144, 147, 1, 0, 0, 0,
		145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 416, 1, 0, 0, 0, 147,
		145, 1, 0, 0, 0, 148, 152, 5, 42, 0, 0, 149, 151, 3, 44, 22, 0, 150, 149,
		1, 0, 0, 0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0,
		0, 0, 153, 416, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 157, 5, 17, 0, 0,
		156, 158, 7, 0, 0, 0, 157, 156, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158,
		159, 1, 0, 0, 0, 159, 167, 5, 16, 0, 0, 160, 162, 3, 44, 22, 0, 161, 160,
		1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0,
		0, 0, 164, 166, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 168, 5, 194, 0,
		0, 167, 163, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169,
		170, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 172, 5, 8, 0, 0, 172, 173,
		5, 16, 0, 0, 173, 416, 5, 194, 0, 0, 174, 176, 5, 33, 0, 0, 175, 177, 5,
		37, 0, 0, 176, 175, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 1, 0, 0,
		0, 178, 182, 5, 79, 0, 0, 179, 181, 3, 44, 22, 0, 180, 179, 1, 0, 0, 0,
		181, 184, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183,
		416, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 185, 186, 5, 47, 0, 0, 186, 188,
		5, 79, 0, 0, 187, 189, 3, 34, 17, 0, 188, 187, 1, 0, 0, 0, 188, 189, 1,
		0, 0, 0, 189, 193, 1, 0, 0, 0, 190, 192, 3, 44, 22, 0, 191, 190, 1, 0,
		0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0,
		194, 416, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 196, 197, 5, 33, 0, 0, 197,
		198, 5, 192, 0, 0, 198, 200, 5, 193, 0, 0, 199, 201, 3, 36, 18, 0, 200,
		199, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 205, 1, 0, 0, 0, 202, 204,
		3, 44, 22, 0, 203, 202, 1, 0, 0, 0, 204, 207, 1, 0, 0, 0, 205, 203, 1,
		0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 416, 1, 0, 0, 0, 207, 205, 1, 0, 0,
		0, 208, 209, 5, 5, 0, 0, 209, 210, 5, 192, 0, 0, 210, 214, 5, 193, 0, 0,
		211, 213, 3, 44, 22, 0, 212, 211, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214,
		212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 416, 1, 0, 0, 0, 216, 214,
		1, 0, 0, 0, 217, 218, 5, 47, 0, 0, 218, 219, 5, 192, 0, 0, 219, 221, 5,
		193, 0, 0, 220, 222, 3, 34, 17, 0, 221, 220, 1, 0, 0, 0, 221, 222, 1, 0,
		0, 0, 222, 226, 1, 0, 0, 0, 223, 225, 3, 44, 22, 0, 224, 223, 1, 0, 0,
		0, 225, 228, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227,
		416, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 229, 230, 5, 33, 0, 0, 230, 232,
		5, 141, 0, 0, 231, 233, 3, 36, 18, 0, 232, 231, 1, 0, 0, 0, 232, 233, 1,
		0, 0, 0, 233, 237, 1, 0, 0, 0, 234, 236, 3, 44, 22, 0, 235, 234, 1, 0,
		0, 0, 236, 239, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0,
		238, 416, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 240, 241, 5, 5, 0, 0, 241,
		245, 5, 141, 0, 0, 242, 244, 3, 44, 22, 0, 243, 242, 1, 0, 0, 0, 244, 247,
		1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 416, 1, 0,
		0, 0, 247, 245, 1, 0, 0, 0, 248, 249, 5, 47, 0, 0, 249, 253, 5, 141, 0,
		0, 250, 252, 3, 44, 22, 0, 251, 250, 1, 0, 0, 0, 252, 255, 1, 0, 0, 0,
		253, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 416, 1, 0, 0, 0, 255,
		253, 1, 0, 0, 0, 256, 260, 5, 70, 0, 0, 257, 259, 3, 44, 22, 0, 258, 257,
		1, 0, 0, 0, 259, 262, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 260, 261, 1, 0,
		0, 0, 261, 416, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 263, 267, 5, 139, 0,
		0, 264, 266, 3, 44, 22, 0, 265, 264, 1, 0, 0, 0, 266, 269, 1, 0, 0, 0,
		267, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 416, 1, 0, 0, 0, 269,
		267, 1, 0, 0, 0, 270, 274, 5, 102, 0, 0, 271, 273, 3, 44, 22, 0, 272, 271,
		1, 0, 0, 0, 273, 276, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 274, 275, 1, 0,
		0, 0, 275, 416, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0, 277, 278, 5, 33, 0, 0,
		278, 282, 5, 179, 0, 0, 279, 281, 3, 44, 22, 0, 280, 279, 1, 0, 0, 0, 281,
		284, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 416,
		1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 285, 286, 5, 5, 0, 0, 286, 290, 5, 179,
		0, 0, 287, 289, 3, 44, 22, 0, 288, 287, 1, 0, 0, 0, 289, 292, 1, 0, 0,
		0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 416, 1, 0, 0, 0, 292,
		290, 1, 0, 0, 0, 293, 294, 5, 47, 0, 0, 294, 298, 5, 179, 0, 0, 295, 297,
		3, 44, 22, 0, 296, 295, 1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 296, 1,
		0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 416, 1, 0, 0, 0, 300, 298, 1, 0, 0,
		0, 301, 304, 5, 33, 0, 0, 302, 303, 5, 121, 0, 0, 303, 305, 5, 136, 0,
		0, 304, 302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306,
		308, 5, 68, 0, 0, 307, 309, 3, 36, 18, 0, 308, 307, 1, 0, 0, 0, 308, 309,
		1, 0, 0, 0, 309, 313, 1, 0, 0, 0, 310, 312, 3, 44, 22, 0, 311, 310, 1,
		0, 0, 0, 312, 315, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 313, 314, 1, 0, 0,
		0, 314, 316, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 316, 317, 5, 10, 0, 0, 317,
		321, 5, 195, 0, 0, 318, 320, 3, 44, 22, 0, 319, 318, 1, 0, 0, 0, 320, 323,
		1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 324, 1, 0,
		0, 0, 323, 321, 1, 0, 0, 0, 324, 325, 5, 194, 0, 0, 325, 416, 5, 195, 0,
		0, 326, 327, 5, 47, 0, 0, 327, 329, 5, 68, 0, 0, 328, 330, 3, 34, 17, 0,
		329, 328, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 334, 1, 0, 0, 0, 331,
		333, 3, 44, 22, 0, 332, 331, 1, 0, 0, 0, 333, 336, 1, 0, 0, 0, 334, 332,
		1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 416, 1, 0, 0, 0, 336, 334, 1, 0,
		0, 0, 337, 340, 5, 33, 0, 0, 338, 339, 5, 121, 0, 0, 339, 341, 5, 136,
		0, 0, 340, 338, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0,
		342, 344, 5, 2, 0, 0, 343, 345, 3, 36, 18, 0, 344, 343, 1, 0, 0, 0, 344,
		345, 1, 0, 0, 0, 345, 349, 1, 0, 0, 0, 346, 348, 3, 44, 22, 0, 347, 346,
		1, 0, 0, 0, 348, 351, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0,
		0, 0, 350, 416, 1, 0, 0, 0, 351, 349, 1, 0, 0, 0, 352, 353, 5, 47, 0, 0,
		353, 355, 5, 2, 0, 0, 354, 356, 3, 34, 17, 0, 355, 354, 1, 0, 0, 0, 355,
		356, 1, 0, 0, 0, 356, 360, 1, 0, 0, 0, 357, 359, 3, 44, 22, 0, 358, 357,
		1, 0, 0, 0, 359, 362, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 360, 361, 1, 0,
		0, 0, 361, 416, 1, 0, 0, 0, 362, 360, 1, 0, 0, 0, 363, 364, 5, 33, 0, 0,
		364, 366, 5, 172, 0, 0, 365, 367, 3, 36, 18, 0, 366, 365, 1, 0, 0, 0, 366,
		367, 1, 0, 0, 0, 367, 371, 1, 0, 0, 0, 368, 370, 3, 44, 22, 0, 369, 368,
		1, 0, 0, 0, 370, 373, 1, 0, 0, 0, 371, 369, 1, 0, 0, 0, 371, 372, 1, 0,
		0, 0, 372, 416, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 374, 375, 5, 5, 0, 0,
		375, 379, 5, 172, 0, 0, 376, 378, 3, 44, 22, 0, 377, 376, 1, 0, 0, 0, 378,
		381, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380, 416,
		1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 382, 383, 5, 47, 0, 0, 383, 385, 5, 172,
		0, 0, 384, 386, 3, 34, 17, 0, 385, 384, 1, 0, 0, 0, 385, 386, 1, 0, 0,
		0, 386, 390, 1, 0, 0, 0, 387, 389, 3, 44, 22, 0, 388, 387, 1, 0, 0, 0,
		389, 392, 1, 0, 0, 0, 390, 388, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391,
		416, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 393, 394, 5, 33, 0, 0, 394, 396,
		5, 167, 0, 0, 395, 397, 3, 36, 18, 0, 396, 395, 1, 0, 0, 0, 396, 397, 1,
		0, 0, 0, 397, 401, 1, 0, 0, 0, 398, 400, 3, 44, 22, 0, 399, 398, 1, 0,
		0, 0, 400, 403, 1, 0, 0, 0, 401, 399, 1, 0, 0, 0, 401, 402, 1, 0, 0, 0,
		402, 416, 1, 0, 0, 0, 403, 401, 1, 0, 0, 0, 404, 405, 5, 47, 0, 0, 405,
		407, 5, 167, 0, 0, 406, 408, 3, 34, 17, 0, 407, 406, 1, 0, 0, 0, 407, 408,
		1, 0, 0, 0, 408, 412, 1, 0, 0, 0, 409, 411, 3, 44, 22, 0, 410, 409, 1,
		0, 0, 0, 411, 414, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 412, 413, 1, 0, 0,
		0, 413, 416, 1, 0, 0, 0, 414, 412, 1, 0, 0, 0, 415, 74, 1, 0, 0, 0, 415,
		76, 1, 0, 0, 0, 415, 89, 1, 0, 0, 0, 415, 99, 1, 0, 0, 0, 415, 105, 1,
		0, 0, 0, 415, 114, 1, 0, 0, 0, 415, 120, 1, 0, 0, 0, 415, 125, 1, 0, 0,
		0, 415, 132, 1, 0, 0, 0, 415, 140, 1, 0, 0, 0, 415, 148, 1, 0, 0, 0, 415,
		155, 1, 0, 0, 0, 415, 174, 1, 0, 0, 0, 415, 185, 1, 0, 0, 0, 415, 196,
		1, 0, 0, 0, 415, 208, 1, 0, 0, 0, 415, 217, 1, 0, 0, 0, 415, 229, 1, 0,
		0, 0, 415, 240, 1, 0, 0, 0, 415, 248, 1, 0, 0, 0, 415, 256, 1, 0, 0, 0,
		415, 263, 1, 0, 0, 0, 415, 270, 1, 0, 0, 0, 415, 277, 1, 0, 0, 0, 415,
		285, 1, 0, 0, 0, 415, 293, 1, 0, 0, 0, 415, 301, 1, 0, 0, 0, 415, 326,
		1, 0, 0, 0, 415, 337, 1, 0, 0, 0, 415, 352, 1, 0, 0, 0, 415, 363, 1, 0,
		0, 0, 415, 374, 1, 0, 0, 0, 415, 382, 1, 0, 0, 0, 415, 393, 1, 0, 0, 0,
		415, 404, 1, 0, 0, 0, 416, 7, 1, 0, 0, 0, 417, 422, 3, 10, 5, 0, 418, 419,
		5, 198, 0, 0, 419, 421, 3, 10, 5, 0, 420, 418, 1, 0, 0, 0, 421, 424, 1,
		0, 0, 0, 422, 420, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 427, 1, 0, 0,
		0, 424, 422, 1, 0, 0, 0, 425, 426, 5, 198, 0, 0, 426, 428, 3, 12, 6, 0,
		427, 425, 1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 9, 1, 0, 0, 0, 429, 430,
		3, 20, 10, 0, 430, 432, 3, 30, 15, 0, 431, 433, 3, 32, 16, 0, 432, 431,
		1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 11, 1, 0, 0, 0, 434, 435, 3, 32,
		16, 0, 435, 436, 5, 199, 0, 0, 436, 437, 3, 14, 7, 0, 437, 438, 5, 200,
		0, 0, 438, 13, 1, 0, 0, 0, 439, 442, 3, 16, 8, 0, 440, 441, 5, 198, 0,
		0, 441, 443, 3, 18, 9, 0, 442, 440, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443,
		15, 1, 0, 0, 0, 444, 457, 3, 20, 10, 0, 445, 446, 5, 199, 0, 0, 446, 451,
		3, 20, 10, 0, 447, 448, 5, 198, 0, 0, 448, 450, 3, 20, 10, 0, 449, 447,
		1, 0, 0, 0, 450, 453, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 451, 452, 1, 0,
		0, 0, 452, 454, 1, 0, 0, 0, 453, 451, 1, 0, 0, 0, 454, 455, 5, 200, 0,
		0, 455, 457, 1, 0, 0, 0, 456, 444, 1, 0, 0, 0, 456, 445, 1, 0, 0, 0, 457,
		17, 1, 0, 0, 0, 458, 463, 3, 20, 10, 0, 459, 460, 5, 198, 0, 0, 460, 462,
		3, 20, 10, 0, 461, 459, 1, 0, 0, 0, 462, 465, 1, 0, 0, 0, 463, 461, 1,
		0, 0, 0, 463, 464, 1, 0, 0, 0, 464, 19, 1, 0, 0, 0, 465, 463, 1, 0, 0,
		0, 466, 467, 3, 28, 14, 0, 467, 21, 1, 0, 0, 0, 468, 469, 3, 24, 12, 0,
		469, 470, 5, 197, 0, 0, 470, 472, 1, 0, 0, 0, 471, 468, 1, 0, 0, 0, 471,
		472, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0, 473, 474, 3, 28, 14, 0, 474, 23,
		1, 0, 0, 0, 475, 476, 3, 26, 13, 0, 476, 25, 1, 0, 0, 0, 477, 480, 5, 204,
		0, 0, 478, 480, 3, 28, 14, 0, 479, 477, 1, 0, 0, 0, 479, 478, 1, 0, 0,
		0, 480, 27, 1, 0, 0, 0, 481, 484, 3, 50, 25, 0, 482, 484, 5, 203, 0, 0,
		483, 481, 1, 0, 0, 0, 483, 482, 1, 0, 0, 0, 484, 29, 1, 0, 0, 0, 485, 488,
		3, 38, 19, 0, 486, 488, 3, 40, 20, 0, 487, 485, 1, 0, 0, 0, 487, 486, 1,
		0, 0, 0, 488, 31, 1, 0, 0, 0, 489, 490, 5, 130, 0, 0, 490, 491, 5, 93,
		0, 0, 491, 33, 1, 0, 0, 0, 492, 493, 5, 76, 0, 0, 493, 494, 5, 56, 0, 0,
		494, 35, 1, 0, 0, 0, 495, 496, 5, 76, 0, 0, 496, 497, 5, 115, 0, 0, 497,
		498, 5, 56, 0, 0, 498, 37, 1, 0, 0, 0, 499, 500, 7, 1, 0, 0, 500, 39, 1,
		0, 0, 0, 501, 502, 5, 105, 0, 0, 502, 503, 5, 201, 0, 0, 503, 504, 3, 38,
		19, 0, 504, 505, 5, 198, 0, 0, 505, 506, 3, 38, 19, 0, 506, 507, 5, 202,
		0, 0, 507, 514, 1, 0, 0, 0, 508, 509, 7, 2, 0, 0, 509, 510, 5, 201, 0,
		0, 510, 511, 3, 38, 19, 0, 511, 512, 5, 202, 0, 0, 512, 514, 1, 0, 0, 0,
		513, 501, 1, 0, 0, 0, 513, 508, 1, 0, 0, 0, 514, 41, 1, 0, 0, 0, 515, 519,
		5, 189, 0, 0, 516, 518, 3, 44, 22, 0, 517, 516, 1, 0, 0, 0, 518, 521, 1,
		0, 0, 0, 519, 517, 1, 0, 0, 0, 519, 520, 1, 0, 0, 0, 520, 43, 1, 0, 0,
		0, 521, 519, 1, 0, 0, 0, 522, 535, 3, 46, 23, 0, 523, 535, 5, 195, 0, 0,
		524, 535, 5, 196, 0, 0, 525, 535, 5, 197, 0, 0, 526, 535, 5, 198, 0, 0,
		527, 535, 5, 199, 0, 0, 528, 535, 5, 200, 0, 0, 529, 535, 5, 201, 0, 0,
		530, 535, 5, 202, 0, 0, 531, 535, 5, 203, 0, 0, 532, 535, 5, 204, 0, 0,
		533, 535, 5, 208, 0, 0, 534, 522, 1, 0, 0, 0, 534, 523, 1, 0, 0, 0, 534,
		524, 1, 0, 0, 0, 534, 525, 1, 0, 0, 0, 534, 526, 1, 0, 0, 0, 534, 527,
		1, 0, 0, 0, 534, 528, 1, 0, 0, 0, 534, 529, 1, 0, 0, 0, 534, 530, 1, 0,
		0, 0, 534, 531, 1, 0, 0, 0, 534, 532, 1, 0, 0, 0, 534, 533, 1, 0, 0, 0,
		535, 45, 1, 0, 0, 0, 536, 539, 3, 48, 24, 0, 537, 539, 3, 50, 25, 0, 538,
		536, 1, 0, 0, 0, 538, 537, 1, 0, 0, 0, 539, 47, 1, 0, 0, 0, 540, 541, 7,
		3, 0, 0, 541, 49, 1, 0, 0, 0, 542, 543, 7, 4, 0, 0, 543, 51, 1, 0, 0, 0,
		73, 54, 60, 65, 72, 79, 86, 96, 102, 111, 117, 122, 129, 137, 145, 152,
		157, 163, 169, 176, 182, 188, 193, 200, 205, 214, 221, 226, 232, 237, 245,
		253, 260, 267, 274, 282, 290, 298, 304, 308, 313, 321, 329, 334, 340, 344,
		349, 355, 360, 366, 371, 379, 385, 390, 396, 401, 407, 412, 415, 422, 427,
		432, 442, 451, 456, 463, 471, 479, 483, 487, 513, 519, 534, 538,
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
	CqlParserK_ANY                  = 7
	CqlParserK_APPLY                = 8
	CqlParserK_ARRAY                = 9
	CqlParserK_AS                   = 10
	CqlParserK_ASC                  = 11
	CqlParserK_ASCII                = 12
	CqlParserK_ASSERT_ROWS_MODIFIED = 13
	CqlParserK_AT                   = 14
	CqlParserK_AUTHORIZE            = 15
	CqlParserK_BATCH                = 16
	CqlParserK_BEGIN                = 17
	CqlParserK_BETWEEN              = 18
	CqlParserK_BIGINT               = 19
	CqlParserK_BLOB                 = 20
	CqlParserK_BOOLEAN              = 21
	CqlParserK_BY                   = 22
	CqlParserK_CALLED               = 23
	CqlParserK_CASE                 = 24
	CqlParserK_CAST                 = 25
	CqlParserK_CLUSTERING           = 26
	CqlParserK_COLLATE              = 27
	CqlParserK_COLUMNFAMILY         = 28
	CqlParserK_COMPACT              = 29
	CqlParserK_CONTAINS             = 30
	CqlParserK_COUNT                = 31
	CqlParserK_COUNTER              = 32
	CqlParserK_CREATE               = 33
	CqlParserK_CROSS                = 34
	CqlParserK_CUBE                 = 35
	CqlParserK_CURRENT              = 36
	CqlParserK_CUSTOM               = 37
	CqlParserK_DATE                 = 38
	CqlParserK_DECIMAL              = 39
	CqlParserK_DEFAULT              = 40
	CqlParserK_DEFINE               = 41
	CqlParserK_DELETE               = 42
	CqlParserK_DESC                 = 43
	CqlParserK_DESCRIBE             = 44
	CqlParserK_DISTINCT             = 45
	CqlParserK_DOUBLE               = 46
	CqlParserK_DROP                 = 47
	CqlParserK_ELSE                 = 48
	CqlParserK_END                  = 49
	CqlParserK_ENTRIES              = 50
	CqlParserK_ENUM                 = 51
	CqlParserK_ESCAPE               = 52
	CqlParserK_EXCEPT               = 53
	CqlParserK_EXCLUDE              = 54
	CqlParserK_EXECUTE              = 55
	CqlParserK_EXISTS               = 56
	CqlParserK_EXTRACT              = 57
	CqlParserK_FALSE                = 58
	CqlParserK_FETCH                = 59
	CqlParserK_FILTERING            = 60
	CqlParserK_FINALFUNC            = 61
	CqlParserK_FLOAT                = 62
	CqlParserK_FOLLOWING            = 63
	CqlParserK_FOR                  = 64
	CqlParserK_FROM                 = 65
	CqlParserK_FROZEN               = 66
	CqlParserK_FULL                 = 67
	CqlParserK_FUNCTION             = 68
	CqlParserK_FUNCTIONS            = 69
	CqlParserK_GRANT                = 70
	CqlParserK_GROUP                = 71
	CqlParserK_GROUPING             = 72
	CqlParserK_GROUPS               = 73
	CqlParserK_HASH                 = 74
	CqlParserK_HAVING               = 75
	CqlParserK_IF                   = 76
	CqlParserK_IGNORE               = 77
	CqlParserK_IN                   = 78
	CqlParserK_INDEX                = 79
	CqlParserK_INET                 = 80
	CqlParserK_INFINITY             = 81
	CqlParserK_INITCOND             = 82
	CqlParserK_INNER                = 83
	CqlParserK_INPUT                = 84
	CqlParserK_INSERT               = 85
	CqlParserK_INT                  = 86
	CqlParserK_INTERSECT            = 87
	CqlParserK_INTERVAL             = 88
	CqlParserK_INTO                 = 89
	CqlParserK_IS                   = 90
	CqlParserK_JOIN                 = 91
	CqlParserK_JSON                 = 92
	CqlParserK_KEY                  = 93
	CqlParserK_KEYS                 = 94
	CqlParserK_KEYSPACE             = 95
	CqlParserK_KEYSPACES            = 96
	CqlParserK_LANGUAGE             = 97
	CqlParserK_LATERAL              = 98
	CqlParserK_LEFT                 = 99
	CqlParserK_LIKE                 = 100
	CqlParserK_LIMIT                = 101
	CqlParserK_LIST                 = 102
	CqlParserK_LOGIN                = 103
	CqlParserK_LOOKUP               = 104
	CqlParserK_MAP                  = 105
	CqlParserK_MERGE                = 106
	CqlParserK_MODIFY               = 107
	CqlParserK_NAN                  = 108
	CqlParserK_NATURAL              = 109
	CqlParserK_NEW                  = 110
	CqlParserK_NO                   = 111
	CqlParserK_NOLOGIN              = 112
	CqlParserK_NORECURSIVE          = 113
	CqlParserK_NOSUPERUSER          = 114
	CqlParserK_NOT                  = 115
	CqlParserK_NULL                 = 116
	CqlParserK_NULLS                = 117
	CqlParserK_OF                   = 118
	CqlParserK_ON                   = 119
	CqlParserK_OPTIONS              = 120
	CqlParserK_OR                   = 121
	CqlParserK_ORDER                = 122
	CqlParserK_OUTER                = 123
	CqlParserK_OVER                 = 124
	CqlParserK_PARTITION            = 125
	CqlParserK_PASSWORD             = 126
	CqlParserK_PERMISSION           = 127
	CqlParserK_PERMISSIONS          = 128
	CqlParserK_PRECEDING            = 129
	CqlParserK_PRIMARY              = 130
	CqlParserK_PROTO                = 131
	CqlParserK_QUALIFY              = 132
	CqlParserK_RANGE                = 133
	CqlParserK_RECURSIVE            = 134
	CqlParserK_RENAME               = 135
	CqlParserK_REPLACE              = 136
	CqlParserK_RESPECT              = 137
	CqlParserK_RETURNS              = 138
	CqlParserK_REVOKE               = 139
	CqlParserK_RIGHT                = 140
	CqlParserK_ROLE                 = 141
	CqlParserK_ROLES                = 142
	CqlParserK_ROLLUP               = 143
	CqlParserK_ROWS                 = 144
	CqlParserK_SCHEMA               = 145
	CqlParserK_SELECT               = 146
	CqlParserK_SET                  = 147
	CqlParserK_SFUNC                = 148
	CqlParserK_SMALLINT             = 149
	CqlParserK_SOME                 = 150
	CqlParserK_STATIC               = 151
	CqlParserK_STORAGE              = 152
	CqlParserK_STRUCT               = 153
	CqlParserK_STYPE                = 154
	CqlParserK_SUPERUSER            = 155
	CqlParserK_TABLE                = 156
	CqlParserK_TABLESAMPLE          = 157
	CqlParserK_TEXT                 = 158
	CqlParserK_THEN                 = 159
	CqlParserK_TIME                 = 160
	CqlParserK_TIMESTAMP            = 161
	CqlParserK_TIMEUUID             = 162
	CqlParserK_TINYINT              = 163
	CqlParserK_TO                   = 164
	CqlParserK_TOKEN                = 165
	CqlParserK_TREAT                = 166
	CqlParserK_TRIGGER              = 167
	CqlParserK_TRUE                 = 168
	CqlParserK_TRUNCATE             = 169
	CqlParserK_TTL                  = 170
	CqlParserK_TUPLE                = 171
	CqlParserK_TYPE                 = 172
	CqlParserK_UNBOUNDED            = 173
	CqlParserK_UNION                = 174
	CqlParserK_UNLOGGED             = 175
	CqlParserK_UNNEST               = 176
	CqlParserK_UPDATE               = 177
	CqlParserK_USE                  = 178
	CqlParserK_USER                 = 179
	CqlParserK_USERS                = 180
	CqlParserK_USING                = 181
	CqlParserK_UUID                 = 182
	CqlParserK_VALUES               = 183
	CqlParserK_VARCHAR              = 184
	CqlParserK_VARINT               = 185
	CqlParserK_WHEN                 = 186
	CqlParserK_WHERE                = 187
	CqlParserK_WINDOW               = 188
	CqlParserK_WITH                 = 189
	CqlParserK_WITHIN               = 190
	CqlParserK_WRITETIME            = 191
	CqlParserK_MATERIALIZED         = 192
	CqlParserK_VIEW                 = 193
	CqlParserSEMICOLON              = 194
	CqlParserSQUOTE                 = 195
	CqlParserDQUOTE                 = 196
	CqlParserDOT                    = 197
	CqlParserCOMMA                  = 198
	CqlParserL_PAREN                = 199
	CqlParserR_PAREN                = 200
	CqlParserL_ANGLE_BRACKET        = 201
	CqlParserR_ANGLE_BRACKET        = 202
	CqlParserIDENTIFIER             = 203
	CqlParserIDENTIFIER_WITH_HYPHEN = 204
	CqlParserWS                     = 205
	CqlParserCOMMENT                = 206
	CqlParserMULTILINE_COMMENT      = 207
	CqlParserUNKNOWN                = 208
)

// CqlParser rules.
const (
	CqlParserRULE_root                 = 0
	CqlParserRULE_cqlStatement         = 1
	CqlParserRULE_createTable          = 2
	CqlParserRULE_unsupportedStatement = 3
	CqlParserRULE_columnDefinitionList = 4
	CqlParserRULE_columnDefinition     = 5
	CqlParserRULE_primaryKeyClause     = 6
	CqlParserRULE_primaryKey           = 7
	CqlParserRULE_partitionKey         = 8
	CqlParserRULE_clusteringColumns    = 9
	CqlParserRULE_columnName           = 10
	CqlParserRULE_tableName            = 11
	CqlParserRULE_keyspaceName         = 12
	CqlParserRULE_generalIdentifier    = 13
	CqlParserRULE_identifier           = 14
	CqlParserRULE_cqlType              = 15
	CqlParserRULE_primaryKeyKeywords   = 16
	CqlParserRULE_ifExists             = 17
	CqlParserRULE_ifNotExists          = 18
	CqlParserRULE_cqlNativeType        = 19
	CqlParserRULE_cqlCollectionType    = 20
	CqlParserRULE_wihtTableOptions     = 21
	CqlParserRULE_nonSemicolonToken    = 22
	CqlParserRULE_keyword              = 23
	CqlParserRULE_reservedKeyword      = 24
	CqlParserRULE_nonReservedKeyword   = 25
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
		p.SetState(52)
		p.CqlStatement()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserSEMICOLON {
		{
			p.SetState(53)
			p.Match(CqlParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(56)
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
	UnsupportedStatement() IUnsupportedStatementContext

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

func (s *CqlStatementContext) UnsupportedStatement() IUnsupportedStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnsupportedStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnsupportedStatementContext)
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
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.CreateTable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.UnsupportedStatement()
		}

	case antlr.ATNInvalidAltNumber:
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
	IfNotExists() IIfNotExistsContext
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

func (s *CreateTableContext) IfNotExists() IIfNotExistsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfNotExistsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
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
		p.SetState(62)
		p.Match(CqlParserK_CREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Match(CqlParserK_TABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_IF {
		{
			p.SetState(64)
			p.IfNotExists()
		}

	}
	{
		p.SetState(67)
		p.TableName()
	}
	{
		p.SetState(68)
		p.Match(CqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.ColumnDefinitionList()
	}
	{
		p.SetState(70)
		p.Match(CqlParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_WITH {
		{
			p.SetState(71)
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

// IUnsupportedStatementContext is an interface to support dynamic dispatch.
type IUnsupportedStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_USE() antlr.TerminalNode
	KeyspaceName() IKeyspaceNameContext
	K_CREATE() antlr.TerminalNode
	K_KEYSPACE() antlr.TerminalNode
	K_WITH() antlr.TerminalNode
	IfNotExists() IIfNotExistsContext
	AllNonSemicolonToken() []INonSemicolonTokenContext
	NonSemicolonToken(i int) INonSemicolonTokenContext
	K_ALTER() antlr.TerminalNode
	K_DROP() antlr.TerminalNode
	IfExists() IIfExistsContext
	K_TABLE() antlr.TerminalNode
	TableName() ITableNameContext
	K_TRUNCATE() antlr.TerminalNode
	K_SELECT() antlr.TerminalNode
	K_INSERT() antlr.TerminalNode
	K_INTO() antlr.TerminalNode
	K_UPDATE() antlr.TerminalNode
	K_DELETE() antlr.TerminalNode
	K_BEGIN() antlr.TerminalNode
	AllK_BATCH() []antlr.TerminalNode
	K_BATCH(i int) antlr.TerminalNode
	K_APPLY() antlr.TerminalNode
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	K_UNLOGGED() antlr.TerminalNode
	K_COUNTER() antlr.TerminalNode
	K_INDEX() antlr.TerminalNode
	K_CUSTOM() antlr.TerminalNode
	K_MATERIALIZED() antlr.TerminalNode
	K_VIEW() antlr.TerminalNode
	K_ROLE() antlr.TerminalNode
	K_GRANT() antlr.TerminalNode
	K_REVOKE() antlr.TerminalNode
	K_LIST() antlr.TerminalNode
	K_USER() antlr.TerminalNode
	K_FUNCTION() antlr.TerminalNode
	K_AS() antlr.TerminalNode
	AllSQUOTE() []antlr.TerminalNode
	SQUOTE(i int) antlr.TerminalNode
	K_OR() antlr.TerminalNode
	K_REPLACE() antlr.TerminalNode
	K_AGGREGATE() antlr.TerminalNode
	K_TYPE() antlr.TerminalNode
	K_TRIGGER() antlr.TerminalNode

	// IsUnsupportedStatementContext differentiates from other interfaces.
	IsUnsupportedStatementContext()
}

type UnsupportedStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnsupportedStatementContext() *UnsupportedStatementContext {
	var p = new(UnsupportedStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_unsupportedStatement
	return p
}

func InitEmptyUnsupportedStatementContext(p *UnsupportedStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_unsupportedStatement
}

func (*UnsupportedStatementContext) IsUnsupportedStatementContext() {}

func NewUnsupportedStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnsupportedStatementContext {
	var p = new(UnsupportedStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_unsupportedStatement

	return p
}

func (s *UnsupportedStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *UnsupportedStatementContext) K_USE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_USE, 0)
}

func (s *UnsupportedStatementContext) KeyspaceName() IKeyspaceNameContext {
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

func (s *UnsupportedStatementContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CREATE, 0)
}

func (s *UnsupportedStatementContext) K_KEYSPACE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_KEYSPACE, 0)
}

func (s *UnsupportedStatementContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(CqlParserK_WITH, 0)
}

func (s *UnsupportedStatementContext) IfNotExists() IIfNotExistsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfNotExistsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *UnsupportedStatementContext) AllNonSemicolonToken() []INonSemicolonTokenContext {
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

func (s *UnsupportedStatementContext) NonSemicolonToken(i int) INonSemicolonTokenContext {
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

func (s *UnsupportedStatementContext) K_ALTER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ALTER, 0)
}

func (s *UnsupportedStatementContext) K_DROP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DROP, 0)
}

func (s *UnsupportedStatementContext) IfExists() IIfExistsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfExistsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *UnsupportedStatementContext) K_TABLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TABLE, 0)
}

func (s *UnsupportedStatementContext) TableName() ITableNameContext {
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

func (s *UnsupportedStatementContext) K_TRUNCATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TRUNCATE, 0)
}

func (s *UnsupportedStatementContext) K_SELECT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SELECT, 0)
}

func (s *UnsupportedStatementContext) K_INSERT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INSERT, 0)
}

func (s *UnsupportedStatementContext) K_INTO() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INTO, 0)
}

func (s *UnsupportedStatementContext) K_UPDATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_UPDATE, 0)
}

func (s *UnsupportedStatementContext) K_DELETE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DELETE, 0)
}

func (s *UnsupportedStatementContext) K_BEGIN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BEGIN, 0)
}

func (s *UnsupportedStatementContext) AllK_BATCH() []antlr.TerminalNode {
	return s.GetTokens(CqlParserK_BATCH)
}

func (s *UnsupportedStatementContext) K_BATCH(i int) antlr.TerminalNode {
	return s.GetToken(CqlParserK_BATCH, i)
}

func (s *UnsupportedStatementContext) K_APPLY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_APPLY, 0)
}

func (s *UnsupportedStatementContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(CqlParserSEMICOLON)
}

func (s *UnsupportedStatementContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(CqlParserSEMICOLON, i)
}

func (s *UnsupportedStatementContext) K_UNLOGGED() antlr.TerminalNode {
	return s.GetToken(CqlParserK_UNLOGGED, 0)
}

func (s *UnsupportedStatementContext) K_COUNTER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_COUNTER, 0)
}

func (s *UnsupportedStatementContext) K_INDEX() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INDEX, 0)
}

func (s *UnsupportedStatementContext) K_CUSTOM() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CUSTOM, 0)
}

func (s *UnsupportedStatementContext) K_MATERIALIZED() antlr.TerminalNode {
	return s.GetToken(CqlParserK_MATERIALIZED, 0)
}

func (s *UnsupportedStatementContext) K_VIEW() antlr.TerminalNode {
	return s.GetToken(CqlParserK_VIEW, 0)
}

func (s *UnsupportedStatementContext) K_ROLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ROLE, 0)
}

func (s *UnsupportedStatementContext) K_GRANT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_GRANT, 0)
}

func (s *UnsupportedStatementContext) K_REVOKE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_REVOKE, 0)
}

func (s *UnsupportedStatementContext) K_LIST() antlr.TerminalNode {
	return s.GetToken(CqlParserK_LIST, 0)
}

func (s *UnsupportedStatementContext) K_USER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_USER, 0)
}

func (s *UnsupportedStatementContext) K_FUNCTION() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FUNCTION, 0)
}

func (s *UnsupportedStatementContext) K_AS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_AS, 0)
}

func (s *UnsupportedStatementContext) AllSQUOTE() []antlr.TerminalNode {
	return s.GetTokens(CqlParserSQUOTE)
}

func (s *UnsupportedStatementContext) SQUOTE(i int) antlr.TerminalNode {
	return s.GetToken(CqlParserSQUOTE, i)
}

func (s *UnsupportedStatementContext) K_OR() antlr.TerminalNode {
	return s.GetToken(CqlParserK_OR, 0)
}

func (s *UnsupportedStatementContext) K_REPLACE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_REPLACE, 0)
}

func (s *UnsupportedStatementContext) K_AGGREGATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_AGGREGATE, 0)
}

func (s *UnsupportedStatementContext) K_TYPE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TYPE, 0)
}

func (s *UnsupportedStatementContext) K_TRIGGER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TRIGGER, 0)
}

func (s *UnsupportedStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnsupportedStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnsupportedStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterUnsupportedStatement(s)
	}
}

func (s *UnsupportedStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitUnsupportedStatement(s)
	}
}

func (p *CqlParser) UnsupportedStatement() (localctx IUnsupportedStatementContext) {
	localctx = NewUnsupportedStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CqlParserRULE_unsupportedStatement)
	var _la int

	var _alt int

	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Match(CqlParserK_USE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.KeyspaceName()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.Match(CqlParserK_CREATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(CqlParserK_KEYSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CqlParserK_IF {
			{
				p.SetState(78)
				p.IfNotExists()
			}

		}
		{
			p.SetState(81)
			p.KeyspaceName()
		}
		{
			p.SetState(82)
			p.Match(CqlParserK_WITH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(83)
				p.NonSemicolonToken()
			}

			p.SetState(88)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(89)
			p.Match(CqlParserK_ALTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Match(CqlParserK_KEYSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.KeyspaceName()
		}
		{
			p.SetState(92)
			p.Match(CqlParserK_WITH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(93)
				p.NonSemicolonToken()
			}

			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(99)
			p.Match(CqlParserK_DROP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Match(CqlParserK_KEYSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CqlParserK_IF {
			{
				p.SetState(101)
				p.IfExists()
			}

		}
		{
			p.SetState(104)
			p.KeyspaceName()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(105)
			p.Match(CqlParserK_ALTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(CqlParserK_TABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.TableName()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(108)
				p.NonSemicolonToken()
			}

			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(114)
			p.Match(CqlParserK_DROP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Match(CqlParserK_TABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CqlParserK_IF {
			{
				p.SetState(116)
				p.IfExists()
			}

		}
		{
			p.SetState(119)
			p.TableName()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(120)
			p.Match(CqlParserK_TRUNCATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CqlParserK_TABLE {
			{
				p.SetState(121)
				p.Match(CqlParserK_TABLE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(124)
			p.TableName()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(125)
			p.Match(CqlParserK_SELECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(126)
				p.NonSemicolonToken()
			}

			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(132)
			p.Match(CqlParserK_INSERT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Match(CqlParserK_INTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(134)
				p.NonSemicolonToken()
			}

			p.SetState(139)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(140)
			p.Match(CqlParserK_UPDATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.TableName()
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(142)
				p.NonSemicolonToken()
			}

			p.SetState(147)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(148)
			p.Match(CqlParserK_DELETE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(149)
				p.NonSemicolonToken()
			}

			p.SetState(154)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(155)
			p.Match(CqlParserK_BEGIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CqlParserK_COUNTER || _la == CqlParserK_UNLOGGED {
			{
				p.SetState(156)
				_la = p.GetTokenStream().LA(1)

				if !(_la == CqlParserK_COUNTER || _la == CqlParserK_UNLOGGED) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(159)
			p.Match(CqlParserK_BATCH)
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
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(163)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
					{
						p.SetState(160)
						p.NonSemicolonToken()
					}

					p.SetState(165)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(166)
					p.Match(CqlParserSEMICOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(169)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(CqlParserK_APPLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Match(CqlParserK_BATCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Match(CqlParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(174)
			p.Match(CqlParserK_CREATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CqlParserK_CUSTOM {
			{
				p.SetState(175)
				p.Match(CqlParserK_CUSTOM)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(178)
			p.Match(CqlParserK_INDEX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(179)
				p.NonSemicolonToken()
			}

			p.SetState(184)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(185)
			p.Match(CqlParserK_DROP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(CqlParserK_INDEX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(187)
				p.IfExists()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(190)
				p.NonSemicolonToken()
			}

			p.SetState(195)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(196)
			p.Match(CqlParserK_CREATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Match(CqlParserK_MATERIALIZED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Match(CqlParserK_VIEW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(199)
				p.IfNotExists()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(202)
				p.NonSemicolonToken()
			}

			p.SetState(207)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(208)
			p.Match(CqlParserK_ALTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.Match(CqlParserK_MATERIALIZED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.Match(CqlParserK_VIEW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(211)
				p.NonSemicolonToken()
			}

			p.SetState(216)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(217)
			p.Match(CqlParserK_DROP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.Match(CqlParserK_MATERIALIZED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Match(CqlParserK_VIEW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(220)
				p.IfExists()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(223)
				p.NonSemicolonToken()
			}

			p.SetState(228)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(229)
			p.Match(CqlParserK_CREATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.Match(CqlParserK_ROLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(231)
				p.IfNotExists()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(234)
				p.NonSemicolonToken()
			}

			p.SetState(239)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(240)
			p.Match(CqlParserK_ALTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Match(CqlParserK_ROLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(242)
				p.NonSemicolonToken()
			}

			p.SetState(247)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(248)
			p.Match(CqlParserK_DROP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.Match(CqlParserK_ROLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(250)
				p.NonSemicolonToken()
			}

			p.SetState(255)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(256)
			p.Match(CqlParserK_GRANT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(257)
				p.NonSemicolonToken()
			}

			p.SetState(262)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(263)
			p.Match(CqlParserK_REVOKE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(264)
				p.NonSemicolonToken()
			}

			p.SetState(269)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(270)
			p.Match(CqlParserK_LIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(271)
				p.NonSemicolonToken()
			}

			p.SetState(276)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(277)
			p.Match(CqlParserK_CREATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)
			p.Match(CqlParserK_USER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(279)
				p.NonSemicolonToken()
			}

			p.SetState(284)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(285)
			p.Match(CqlParserK_ALTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)
			p.Match(CqlParserK_USER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(287)
				p.NonSemicolonToken()
			}

			p.SetState(292)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(293)
			p.Match(CqlParserK_DROP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)
			p.Match(CqlParserK_USER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(295)
				p.NonSemicolonToken()
			}

			p.SetState(300)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 27:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(301)
			p.Match(CqlParserK_CREATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CqlParserK_OR {
			{
				p.SetState(302)
				p.Match(CqlParserK_OR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(303)
				p.Match(CqlParserK_REPLACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(306)
			p.Match(CqlParserK_FUNCTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(307)
				p.IfNotExists()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(310)
					p.NonSemicolonToken()
				}

			}
			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.Match(CqlParserK_AS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)
			p.Match(CqlParserSQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(318)
				p.NonSemicolonToken()
			}

			p.SetState(323)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(324)
			p.Match(CqlParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)
			p.Match(CqlParserSQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 28:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(326)
			p.Match(CqlParserK_DROP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Match(CqlParserK_FUNCTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(328)
				p.IfExists()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(331)
				p.NonSemicolonToken()
			}

			p.SetState(336)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 29:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(337)
			p.Match(CqlParserK_CREATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CqlParserK_OR {
			{
				p.SetState(338)
				p.Match(CqlParserK_OR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(339)
				p.Match(CqlParserK_REPLACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(342)
			p.Match(CqlParserK_AGGREGATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(343)
				p.IfNotExists()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(346)
				p.NonSemicolonToken()
			}

			p.SetState(351)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 30:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(352)
			p.Match(CqlParserK_DROP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.Match(CqlParserK_AGGREGATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(355)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(354)
				p.IfExists()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(357)
				p.NonSemicolonToken()
			}

			p.SetState(362)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 31:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(363)
			p.Match(CqlParserK_CREATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)
			p.Match(CqlParserK_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(365)
				p.IfNotExists()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(368)
				p.NonSemicolonToken()
			}

			p.SetState(373)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 32:
		p.EnterOuterAlt(localctx, 32)
		{
			p.SetState(374)
			p.Match(CqlParserK_ALTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.Match(CqlParserK_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(379)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(376)
				p.NonSemicolonToken()
			}

			p.SetState(381)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 33:
		p.EnterOuterAlt(localctx, 33)
		{
			p.SetState(382)
			p.Match(CqlParserK_DROP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.Match(CqlParserK_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(385)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(384)
				p.IfExists()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(387)
				p.NonSemicolonToken()
			}

			p.SetState(392)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 34:
		p.EnterOuterAlt(localctx, 34)
		{
			p.SetState(393)
			p.Match(CqlParserK_CREATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(394)
			p.Match(CqlParserK_TRIGGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(396)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(395)
				p.IfNotExists()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(398)
				p.NonSemicolonToken()
			}

			p.SetState(403)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 35:
		p.EnterOuterAlt(localctx, 35)
		{
			p.SetState(404)
			p.Match(CqlParserK_DROP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.Match(CqlParserK_TRIGGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(407)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(406)
				p.IfExists()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
			{
				p.SetState(409)
				p.NonSemicolonToken()
			}

			p.SetState(414)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case antlr.ATNInvalidAltNumber:
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
	p.EnterRule(localctx, 8, CqlParserRULE_columnDefinitionList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(417)
		p.ColumnDefinition()
	}
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(418)
				p.Match(CqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(419)
				p.ColumnDefinition()
			}

		}
		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserCOMMA {
		{
			p.SetState(425)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(426)
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
	p.EnterRule(localctx, 10, CqlParserRULE_columnDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(429)
		p.ColumnName()
	}
	{
		p.SetState(430)
		p.CqlType()
	}
	p.SetState(432)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_PRIMARY {
		{
			p.SetState(431)
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
	p.EnterRule(localctx, 12, CqlParserRULE_primaryKeyClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		p.PrimaryKeyKeywords()
	}
	{
		p.SetState(435)
		p.Match(CqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(436)
		p.PrimaryKey()
	}
	{
		p.SetState(437)
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
	p.EnterRule(localctx, 14, CqlParserRULE_primaryKey)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(439)
		p.PartitionKey()
	}
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserCOMMA {
		{
			p.SetState(440)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(441)
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
	p.EnterRule(localctx, 16, CqlParserRULE_partitionKey)
	var _la int

	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_AGGREGATE, CqlParserK_ASCII, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_CALLED, CqlParserK_CLUSTERING, CqlParserK_COMPACT, CqlParserK_COUNT, CqlParserK_COUNTER, CqlParserK_CUSTOM, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DOUBLE, CqlParserK_FILTERING, CqlParserK_FINALFUNC, CqlParserK_FLOAT, CqlParserK_FROZEN, CqlParserK_FUNCTION, CqlParserK_FUNCTIONS, CqlParserK_INET, CqlParserK_INITCOND, CqlParserK_INPUT, CqlParserK_INT, CqlParserK_JSON, CqlParserK_KEY, CqlParserK_KEYS, CqlParserK_KEYSPACES, CqlParserK_LANGUAGE, CqlParserK_LIST, CqlParserK_LOGIN, CqlParserK_MAP, CqlParserK_NOLOGIN, CqlParserK_NOSUPERUSER, CqlParserK_OPTIONS, CqlParserK_PASSWORD, CqlParserK_PERMISSION, CqlParserK_PERMISSIONS, CqlParserK_RETURNS, CqlParserK_ROLE, CqlParserK_ROLES, CqlParserK_SFUNC, CqlParserK_SMALLINT, CqlParserK_STATIC, CqlParserK_STORAGE, CqlParserK_STYPE, CqlParserK_SUPERUSER, CqlParserK_TEXT, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_TRIGGER, CqlParserK_TTL, CqlParserK_TUPLE, CqlParserK_TYPE, CqlParserK_USER, CqlParserK_USERS, CqlParserK_UUID, CqlParserK_VALUES, CqlParserK_VARCHAR, CqlParserK_VARINT, CqlParserK_WRITETIME, CqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(444)
			p.ColumnName()
		}

	case CqlParserL_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(445)
			p.Match(CqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)
			p.ColumnName()
		}
		p.SetState(451)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CqlParserCOMMA {
			{
				p.SetState(447)
				p.Match(CqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(448)
				p.ColumnName()
			}

			p.SetState(453)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(454)
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
	p.EnterRule(localctx, 18, CqlParserRULE_clusteringColumns)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(458)
		p.ColumnName()
	}
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CqlParserCOMMA {
		{
			p.SetState(459)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(460)
			p.ColumnName()
		}

		p.SetState(465)
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
	p.EnterRule(localctx, 20, CqlParserRULE_columnName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(466)
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
	p.EnterRule(localctx, 22, CqlParserRULE_tableName)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(471)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 65, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(468)
			p.KeyspaceName()
		}
		{
			p.SetState(469)
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
		p.SetState(473)
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
	p.EnterRule(localctx, 24, CqlParserRULE_keyspaceName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(475)
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
	p.EnterRule(localctx, 26, CqlParserRULE_generalIdentifier)
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserIDENTIFIER_WITH_HYPHEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(477)
			p.Match(CqlParserIDENTIFIER_WITH_HYPHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserK_AGGREGATE, CqlParserK_ASCII, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_CALLED, CqlParserK_CLUSTERING, CqlParserK_COMPACT, CqlParserK_COUNT, CqlParserK_COUNTER, CqlParserK_CUSTOM, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DOUBLE, CqlParserK_FILTERING, CqlParserK_FINALFUNC, CqlParserK_FLOAT, CqlParserK_FROZEN, CqlParserK_FUNCTION, CqlParserK_FUNCTIONS, CqlParserK_INET, CqlParserK_INITCOND, CqlParserK_INPUT, CqlParserK_INT, CqlParserK_JSON, CqlParserK_KEY, CqlParserK_KEYS, CqlParserK_KEYSPACES, CqlParserK_LANGUAGE, CqlParserK_LIST, CqlParserK_LOGIN, CqlParserK_MAP, CqlParserK_NOLOGIN, CqlParserK_NOSUPERUSER, CqlParserK_OPTIONS, CqlParserK_PASSWORD, CqlParserK_PERMISSION, CqlParserK_PERMISSIONS, CqlParserK_RETURNS, CqlParserK_ROLE, CqlParserK_ROLES, CqlParserK_SFUNC, CqlParserK_SMALLINT, CqlParserK_STATIC, CqlParserK_STORAGE, CqlParserK_STYPE, CqlParserK_SUPERUSER, CqlParserK_TEXT, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_TRIGGER, CqlParserK_TTL, CqlParserK_TUPLE, CqlParserK_TYPE, CqlParserK_USER, CqlParserK_USERS, CqlParserK_UUID, CqlParserK_VALUES, CqlParserK_VARCHAR, CqlParserK_VARINT, CqlParserK_WRITETIME, CqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(478)
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
	p.EnterRule(localctx, 28, CqlParserRULE_identifier)
	p.SetState(483)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_AGGREGATE, CqlParserK_ASCII, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_CALLED, CqlParserK_CLUSTERING, CqlParserK_COMPACT, CqlParserK_COUNT, CqlParserK_COUNTER, CqlParserK_CUSTOM, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DOUBLE, CqlParserK_FILTERING, CqlParserK_FINALFUNC, CqlParserK_FLOAT, CqlParserK_FROZEN, CqlParserK_FUNCTION, CqlParserK_FUNCTIONS, CqlParserK_INET, CqlParserK_INITCOND, CqlParserK_INPUT, CqlParserK_INT, CqlParserK_JSON, CqlParserK_KEY, CqlParserK_KEYS, CqlParserK_KEYSPACES, CqlParserK_LANGUAGE, CqlParserK_LIST, CqlParserK_LOGIN, CqlParserK_MAP, CqlParserK_NOLOGIN, CqlParserK_NOSUPERUSER, CqlParserK_OPTIONS, CqlParserK_PASSWORD, CqlParserK_PERMISSION, CqlParserK_PERMISSIONS, CqlParserK_RETURNS, CqlParserK_ROLE, CqlParserK_ROLES, CqlParserK_SFUNC, CqlParserK_SMALLINT, CqlParserK_STATIC, CqlParserK_STORAGE, CqlParserK_STYPE, CqlParserK_SUPERUSER, CqlParserK_TEXT, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_TRIGGER, CqlParserK_TTL, CqlParserK_TUPLE, CqlParserK_TYPE, CqlParserK_USER, CqlParserK_USERS, CqlParserK_UUID, CqlParserK_VALUES, CqlParserK_VARCHAR, CqlParserK_VARINT, CqlParserK_WRITETIME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(481)
			p.NonReservedKeyword()
		}

	case CqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(482)
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
	p.EnterRule(localctx, 30, CqlParserRULE_cqlType)
	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_ASCII, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_COUNTER, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DOUBLE, CqlParserK_FLOAT, CqlParserK_INET, CqlParserK_INT, CqlParserK_SMALLINT, CqlParserK_TEXT, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_UUID, CqlParserK_VARCHAR, CqlParserK_VARINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(485)
			p.CqlNativeType()
		}

	case CqlParserK_LIST, CqlParserK_MAP, CqlParserK_SET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(486)
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
	p.EnterRule(localctx, 32, CqlParserRULE_primaryKeyKeywords)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(489)
		p.Match(CqlParserK_PRIMARY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(490)
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

// IIfExistsContext is an interface to support dynamic dispatch.
type IIfExistsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_IF() antlr.TerminalNode
	K_EXISTS() antlr.TerminalNode

	// IsIfExistsContext differentiates from other interfaces.
	IsIfExistsContext()
}

type IfExistsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfExistsContext() *IfExistsContext {
	var p = new(IfExistsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_ifExists
	return p
}

func InitEmptyIfExistsContext(p *IfExistsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_ifExists
}

func (*IfExistsContext) IsIfExistsContext() {}

func NewIfExistsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfExistsContext {
	var p = new(IfExistsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_ifExists

	return p
}

func (s *IfExistsContext) GetParser() antlr.Parser { return s.parser }

func (s *IfExistsContext) K_IF() antlr.TerminalNode {
	return s.GetToken(CqlParserK_IF, 0)
}

func (s *IfExistsContext) K_EXISTS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_EXISTS, 0)
}

func (s *IfExistsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExistsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfExistsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterIfExists(s)
	}
}

func (s *IfExistsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitIfExists(s)
	}
}

func (p *CqlParser) IfExists() (localctx IIfExistsContext) {
	localctx = NewIfExistsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CqlParserRULE_ifExists)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(492)
		p.Match(CqlParserK_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(493)
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

// IIfNotExistsContext is an interface to support dynamic dispatch.
type IIfNotExistsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_IF() antlr.TerminalNode
	K_NOT() antlr.TerminalNode
	K_EXISTS() antlr.TerminalNode

	// IsIfNotExistsContext differentiates from other interfaces.
	IsIfNotExistsContext()
}

type IfNotExistsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfNotExistsContext() *IfNotExistsContext {
	var p = new(IfNotExistsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_ifNotExists
	return p
}

func InitEmptyIfNotExistsContext(p *IfNotExistsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_ifNotExists
}

func (*IfNotExistsContext) IsIfNotExistsContext() {}

func NewIfNotExistsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfNotExistsContext {
	var p = new(IfNotExistsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_ifNotExists

	return p
}

func (s *IfNotExistsContext) GetParser() antlr.Parser { return s.parser }

func (s *IfNotExistsContext) K_IF() antlr.TerminalNode {
	return s.GetToken(CqlParserK_IF, 0)
}

func (s *IfNotExistsContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NOT, 0)
}

func (s *IfNotExistsContext) K_EXISTS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_EXISTS, 0)
}

func (s *IfNotExistsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfNotExistsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfNotExistsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterIfNotExists(s)
	}
}

func (s *IfNotExistsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitIfNotExists(s)
	}
}

func (p *CqlParser) IfNotExists() (localctx IIfNotExistsContext) {
	localctx = NewIfNotExistsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CqlParserRULE_ifNotExists)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(495)
		p.Match(CqlParserK_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(496)
		p.Match(CqlParserK_NOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(497)
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
	p.EnterRule(localctx, 38, CqlParserRULE_cqlNativeType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(499)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4611757216103927808) != 0) || _la == CqlParserK_INET || _la == CqlParserK_INT || ((int64((_la-149)) & ^0x3f) == 0 && ((int64(1)<<(_la-149))&111669180929) != 0)) {
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
	p.EnterRule(localctx, 40, CqlParserRULE_cqlCollectionType)
	var _la int

	p.SetState(513)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_MAP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(501)
			p.Match(CqlParserK_MAP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(502)
			p.Match(CqlParserL_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(503)
			p.CqlNativeType()
		}
		{
			p.SetState(504)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(505)
			p.CqlNativeType()
		}
		{
			p.SetState(506)
			p.Match(CqlParserR_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserK_LIST, CqlParserK_SET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(508)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CqlParserK_LIST || _la == CqlParserK_SET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(509)
			p.Match(CqlParserL_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(510)
			p.CqlNativeType()
		}
		{
			p.SetState(511)
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
	p.EnterRule(localctx, 42, CqlParserRULE_wihtTableOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(515)
		p.Match(CqlParserK_WITH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(519)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-1) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-1) != 0) || ((int64((_la-192)) & ^0x3f) == 0 && ((int64(1)<<(_la-192))&73723) != 0) {
		{
			p.SetState(516)
			p.NonSemicolonToken()
		}

		p.SetState(521)
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
	SQUOTE() antlr.TerminalNode
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

func (s *NonSemicolonTokenContext) SQUOTE() antlr.TerminalNode {
	return s.GetToken(CqlParserSQUOTE, 0)
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
	p.EnterRule(localctx, 44, CqlParserRULE_nonSemicolonToken)
	p.SetState(534)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_ADD, CqlParserK_AGGREGATE, CqlParserK_ALL, CqlParserK_ALLOW, CqlParserK_ALTER, CqlParserK_AND, CqlParserK_ANY, CqlParserK_APPLY, CqlParserK_ARRAY, CqlParserK_AS, CqlParserK_ASC, CqlParserK_ASCII, CqlParserK_ASSERT_ROWS_MODIFIED, CqlParserK_AT, CqlParserK_AUTHORIZE, CqlParserK_BATCH, CqlParserK_BEGIN, CqlParserK_BETWEEN, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_BY, CqlParserK_CALLED, CqlParserK_CASE, CqlParserK_CAST, CqlParserK_CLUSTERING, CqlParserK_COLLATE, CqlParserK_COLUMNFAMILY, CqlParserK_COMPACT, CqlParserK_CONTAINS, CqlParserK_COUNT, CqlParserK_COUNTER, CqlParserK_CREATE, CqlParserK_CROSS, CqlParserK_CUBE, CqlParserK_CURRENT, CqlParserK_CUSTOM, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DEFAULT, CqlParserK_DEFINE, CqlParserK_DELETE, CqlParserK_DESC, CqlParserK_DESCRIBE, CqlParserK_DISTINCT, CqlParserK_DOUBLE, CqlParserK_DROP, CqlParserK_ELSE, CqlParserK_END, CqlParserK_ENTRIES, CqlParserK_ENUM, CqlParserK_ESCAPE, CqlParserK_EXCEPT, CqlParserK_EXCLUDE, CqlParserK_EXECUTE, CqlParserK_EXISTS, CqlParserK_EXTRACT, CqlParserK_FALSE, CqlParserK_FETCH, CqlParserK_FILTERING, CqlParserK_FINALFUNC, CqlParserK_FLOAT, CqlParserK_FOLLOWING, CqlParserK_FOR, CqlParserK_FROM, CqlParserK_FROZEN, CqlParserK_FULL, CqlParserK_FUNCTION, CqlParserK_FUNCTIONS, CqlParserK_GRANT, CqlParserK_GROUP, CqlParserK_GROUPING, CqlParserK_GROUPS, CqlParserK_HASH, CqlParserK_HAVING, CqlParserK_IF, CqlParserK_IGNORE, CqlParserK_IN, CqlParserK_INDEX, CqlParserK_INET, CqlParserK_INFINITY, CqlParserK_INITCOND, CqlParserK_INNER, CqlParserK_INPUT, CqlParserK_INSERT, CqlParserK_INT, CqlParserK_INTERSECT, CqlParserK_INTERVAL, CqlParserK_INTO, CqlParserK_IS, CqlParserK_JOIN, CqlParserK_JSON, CqlParserK_KEY, CqlParserK_KEYS, CqlParserK_KEYSPACE, CqlParserK_KEYSPACES, CqlParserK_LANGUAGE, CqlParserK_LATERAL, CqlParserK_LEFT, CqlParserK_LIKE, CqlParserK_LIMIT, CqlParserK_LIST, CqlParserK_LOGIN, CqlParserK_LOOKUP, CqlParserK_MAP, CqlParserK_MERGE, CqlParserK_MODIFY, CqlParserK_NAN, CqlParserK_NATURAL, CqlParserK_NEW, CqlParserK_NO, CqlParserK_NOLOGIN, CqlParserK_NORECURSIVE, CqlParserK_NOSUPERUSER, CqlParserK_NOT, CqlParserK_NULL, CqlParserK_NULLS, CqlParserK_OF, CqlParserK_ON, CqlParserK_OPTIONS, CqlParserK_OR, CqlParserK_ORDER, CqlParserK_OUTER, CqlParserK_OVER, CqlParserK_PARTITION, CqlParserK_PASSWORD, CqlParserK_PERMISSION, CqlParserK_PERMISSIONS, CqlParserK_PRECEDING, CqlParserK_PRIMARY, CqlParserK_PROTO, CqlParserK_QUALIFY, CqlParserK_RANGE, CqlParserK_RECURSIVE, CqlParserK_RENAME, CqlParserK_REPLACE, CqlParserK_RESPECT, CqlParserK_RETURNS, CqlParserK_REVOKE, CqlParserK_RIGHT, CqlParserK_ROLE, CqlParserK_ROLES, CqlParserK_ROLLUP, CqlParserK_ROWS, CqlParserK_SCHEMA, CqlParserK_SELECT, CqlParserK_SET, CqlParserK_SFUNC, CqlParserK_SMALLINT, CqlParserK_SOME, CqlParserK_STATIC, CqlParserK_STORAGE, CqlParserK_STRUCT, CqlParserK_STYPE, CqlParserK_SUPERUSER, CqlParserK_TABLE, CqlParserK_TABLESAMPLE, CqlParserK_TEXT, CqlParserK_THEN, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_TO, CqlParserK_TOKEN, CqlParserK_TREAT, CqlParserK_TRIGGER, CqlParserK_TRUE, CqlParserK_TRUNCATE, CqlParserK_TTL, CqlParserK_TUPLE, CqlParserK_TYPE, CqlParserK_UNBOUNDED, CqlParserK_UNION, CqlParserK_UNLOGGED, CqlParserK_UNNEST, CqlParserK_UPDATE, CqlParserK_USE, CqlParserK_USER, CqlParserK_USERS, CqlParserK_USING, CqlParserK_UUID, CqlParserK_VALUES, CqlParserK_VARCHAR, CqlParserK_VARINT, CqlParserK_WHEN, CqlParserK_WHERE, CqlParserK_WINDOW, CqlParserK_WITH, CqlParserK_WITHIN, CqlParserK_WRITETIME, CqlParserK_MATERIALIZED, CqlParserK_VIEW:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(522)
			p.Keyword()
		}

	case CqlParserSQUOTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(523)
			p.Match(CqlParserSQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserDQUOTE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(524)
			p.Match(CqlParserDQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserDOT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(525)
			p.Match(CqlParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserCOMMA:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(526)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserL_PAREN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(527)
			p.Match(CqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserR_PAREN:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(528)
			p.Match(CqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserL_ANGLE_BRACKET:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(529)
			p.Match(CqlParserL_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserR_ANGLE_BRACKET:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(530)
			p.Match(CqlParserR_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(531)
			p.Match(CqlParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserIDENTIFIER_WITH_HYPHEN:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(532)
			p.Match(CqlParserIDENTIFIER_WITH_HYPHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserUNKNOWN:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(533)
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
	p.EnterRule(localctx, 46, CqlParserRULE_keyword)
	p.SetState(538)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_ADD, CqlParserK_ALL, CqlParserK_ALLOW, CqlParserK_ALTER, CqlParserK_AND, CqlParserK_ANY, CqlParserK_APPLY, CqlParserK_ARRAY, CqlParserK_AS, CqlParserK_ASC, CqlParserK_ASSERT_ROWS_MODIFIED, CqlParserK_AT, CqlParserK_AUTHORIZE, CqlParserK_BATCH, CqlParserK_BEGIN, CqlParserK_BETWEEN, CqlParserK_BY, CqlParserK_CASE, CqlParserK_CAST, CqlParserK_COLLATE, CqlParserK_COLUMNFAMILY, CqlParserK_CONTAINS, CqlParserK_CREATE, CqlParserK_CROSS, CqlParserK_CUBE, CqlParserK_CURRENT, CqlParserK_DEFAULT, CqlParserK_DEFINE, CqlParserK_DELETE, CqlParserK_DESC, CqlParserK_DESCRIBE, CqlParserK_DISTINCT, CqlParserK_DROP, CqlParserK_ELSE, CqlParserK_END, CqlParserK_ENTRIES, CqlParserK_ENUM, CqlParserK_ESCAPE, CqlParserK_EXCEPT, CqlParserK_EXCLUDE, CqlParserK_EXECUTE, CqlParserK_EXISTS, CqlParserK_EXTRACT, CqlParserK_FALSE, CqlParserK_FETCH, CqlParserK_FOLLOWING, CqlParserK_FOR, CqlParserK_FROM, CqlParserK_FULL, CqlParserK_GRANT, CqlParserK_GROUP, CqlParserK_GROUPING, CqlParserK_GROUPS, CqlParserK_HASH, CqlParserK_HAVING, CqlParserK_IF, CqlParserK_IGNORE, CqlParserK_IN, CqlParserK_INDEX, CqlParserK_INFINITY, CqlParserK_INNER, CqlParserK_INSERT, CqlParserK_INTERSECT, CqlParserK_INTERVAL, CqlParserK_INTO, CqlParserK_IS, CqlParserK_JOIN, CqlParserK_KEYSPACE, CqlParserK_LATERAL, CqlParserK_LEFT, CqlParserK_LIKE, CqlParserK_LIMIT, CqlParserK_LOOKUP, CqlParserK_MERGE, CqlParserK_MODIFY, CqlParserK_NAN, CqlParserK_NATURAL, CqlParserK_NEW, CqlParserK_NO, CqlParserK_NORECURSIVE, CqlParserK_NOT, CqlParserK_NULL, CqlParserK_NULLS, CqlParserK_OF, CqlParserK_ON, CqlParserK_OR, CqlParserK_ORDER, CqlParserK_OUTER, CqlParserK_OVER, CqlParserK_PARTITION, CqlParserK_PRECEDING, CqlParserK_PRIMARY, CqlParserK_PROTO, CqlParserK_QUALIFY, CqlParserK_RANGE, CqlParserK_RECURSIVE, CqlParserK_RENAME, CqlParserK_REPLACE, CqlParserK_RESPECT, CqlParserK_REVOKE, CqlParserK_RIGHT, CqlParserK_ROLLUP, CqlParserK_ROWS, CqlParserK_SCHEMA, CqlParserK_SELECT, CqlParserK_SET, CqlParserK_SOME, CqlParserK_STRUCT, CqlParserK_TABLE, CqlParserK_TABLESAMPLE, CqlParserK_THEN, CqlParserK_TO, CqlParserK_TOKEN, CqlParserK_TREAT, CqlParserK_TRUE, CqlParserK_TRUNCATE, CqlParserK_UNBOUNDED, CqlParserK_UNION, CqlParserK_UNLOGGED, CqlParserK_UNNEST, CqlParserK_UPDATE, CqlParserK_USE, CqlParserK_USING, CqlParserK_WHEN, CqlParserK_WHERE, CqlParserK_WINDOW, CqlParserK_WITH, CqlParserK_WITHIN, CqlParserK_MATERIALIZED, CqlParserK_VIEW:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(536)
			p.ReservedKeyword()
		}

	case CqlParserK_AGGREGATE, CqlParserK_ASCII, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_CALLED, CqlParserK_CLUSTERING, CqlParserK_COMPACT, CqlParserK_COUNT, CqlParserK_COUNTER, CqlParserK_CUSTOM, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DOUBLE, CqlParserK_FILTERING, CqlParserK_FINALFUNC, CqlParserK_FLOAT, CqlParserK_FROZEN, CqlParserK_FUNCTION, CqlParserK_FUNCTIONS, CqlParserK_INET, CqlParserK_INITCOND, CqlParserK_INPUT, CqlParserK_INT, CqlParserK_JSON, CqlParserK_KEY, CqlParserK_KEYS, CqlParserK_KEYSPACES, CqlParserK_LANGUAGE, CqlParserK_LIST, CqlParserK_LOGIN, CqlParserK_MAP, CqlParserK_NOLOGIN, CqlParserK_NOSUPERUSER, CqlParserK_OPTIONS, CqlParserK_PASSWORD, CqlParserK_PERMISSION, CqlParserK_PERMISSIONS, CqlParserK_RETURNS, CqlParserK_ROLE, CqlParserK_ROLES, CqlParserK_SFUNC, CqlParserK_SMALLINT, CqlParserK_STATIC, CqlParserK_STORAGE, CqlParserK_STYPE, CqlParserK_SUPERUSER, CqlParserK_TEXT, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_TRIGGER, CqlParserK_TTL, CqlParserK_TUPLE, CqlParserK_TYPE, CqlParserK_USER, CqlParserK_USERS, CqlParserK_UUID, CqlParserK_VALUES, CqlParserK_VARCHAR, CqlParserK_VARINT, CqlParserK_WRITETIME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(537)
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
	K_ALL() antlr.TerminalNode
	K_ALLOW() antlr.TerminalNode
	K_ALTER() antlr.TerminalNode
	K_AND() antlr.TerminalNode
	K_ANY() antlr.TerminalNode
	K_APPLY() antlr.TerminalNode
	K_ARRAY() antlr.TerminalNode
	K_AS() antlr.TerminalNode
	K_ASC() antlr.TerminalNode
	K_ASSERT_ROWS_MODIFIED() antlr.TerminalNode
	K_AT() antlr.TerminalNode
	K_AUTHORIZE() antlr.TerminalNode
	K_BATCH() antlr.TerminalNode
	K_BEGIN() antlr.TerminalNode
	K_BETWEEN() antlr.TerminalNode
	K_BY() antlr.TerminalNode
	K_CASE() antlr.TerminalNode
	K_CAST() antlr.TerminalNode
	K_COLLATE() antlr.TerminalNode
	K_COLUMNFAMILY() antlr.TerminalNode
	K_CONTAINS() antlr.TerminalNode
	K_CREATE() antlr.TerminalNode
	K_CROSS() antlr.TerminalNode
	K_CUBE() antlr.TerminalNode
	K_CURRENT() antlr.TerminalNode
	K_DEFAULT() antlr.TerminalNode
	K_DEFINE() antlr.TerminalNode
	K_DELETE() antlr.TerminalNode
	K_DESC() antlr.TerminalNode
	K_DESCRIBE() antlr.TerminalNode
	K_DISTINCT() antlr.TerminalNode
	K_DROP() antlr.TerminalNode
	K_ELSE() antlr.TerminalNode
	K_END() antlr.TerminalNode
	K_ENTRIES() antlr.TerminalNode
	K_ENUM() antlr.TerminalNode
	K_ESCAPE() antlr.TerminalNode
	K_EXCEPT() antlr.TerminalNode
	K_EXCLUDE() antlr.TerminalNode
	K_EXECUTE() antlr.TerminalNode
	K_EXISTS() antlr.TerminalNode
	K_EXTRACT() antlr.TerminalNode
	K_FALSE() antlr.TerminalNode
	K_FETCH() antlr.TerminalNode
	K_FOLLOWING() antlr.TerminalNode
	K_FOR() antlr.TerminalNode
	K_FROM() antlr.TerminalNode
	K_FULL() antlr.TerminalNode
	K_GRANT() antlr.TerminalNode
	K_GROUP() antlr.TerminalNode
	K_GROUPING() antlr.TerminalNode
	K_GROUPS() antlr.TerminalNode
	K_HASH() antlr.TerminalNode
	K_HAVING() antlr.TerminalNode
	K_IF() antlr.TerminalNode
	K_IGNORE() antlr.TerminalNode
	K_IN() antlr.TerminalNode
	K_INDEX() antlr.TerminalNode
	K_INFINITY() antlr.TerminalNode
	K_INNER() antlr.TerminalNode
	K_INSERT() antlr.TerminalNode
	K_INTERSECT() antlr.TerminalNode
	K_INTERVAL() antlr.TerminalNode
	K_INTO() antlr.TerminalNode
	K_IS() antlr.TerminalNode
	K_JOIN() antlr.TerminalNode
	K_KEYSPACE() antlr.TerminalNode
	K_LATERAL() antlr.TerminalNode
	K_LEFT() antlr.TerminalNode
	K_LIKE() antlr.TerminalNode
	K_LIMIT() antlr.TerminalNode
	K_LOOKUP() antlr.TerminalNode
	K_MATERIALIZED() antlr.TerminalNode
	K_MERGE() antlr.TerminalNode
	K_MODIFY() antlr.TerminalNode
	K_NAN() antlr.TerminalNode
	K_NATURAL() antlr.TerminalNode
	K_NEW() antlr.TerminalNode
	K_NO() antlr.TerminalNode
	K_NORECURSIVE() antlr.TerminalNode
	K_NOT() antlr.TerminalNode
	K_NULL() antlr.TerminalNode
	K_NULLS() antlr.TerminalNode
	K_OF() antlr.TerminalNode
	K_ON() antlr.TerminalNode
	K_OR() antlr.TerminalNode
	K_ORDER() antlr.TerminalNode
	K_OUTER() antlr.TerminalNode
	K_OVER() antlr.TerminalNode
	K_PARTITION() antlr.TerminalNode
	K_PRECEDING() antlr.TerminalNode
	K_PRIMARY() antlr.TerminalNode
	K_PROTO() antlr.TerminalNode
	K_QUALIFY() antlr.TerminalNode
	K_RANGE() antlr.TerminalNode
	K_RECURSIVE() antlr.TerminalNode
	K_RENAME() antlr.TerminalNode
	K_REPLACE() antlr.TerminalNode
	K_RESPECT() antlr.TerminalNode
	K_REVOKE() antlr.TerminalNode
	K_RIGHT() antlr.TerminalNode
	K_ROLLUP() antlr.TerminalNode
	K_ROWS() antlr.TerminalNode
	K_SCHEMA() antlr.TerminalNode
	K_SELECT() antlr.TerminalNode
	K_SET() antlr.TerminalNode
	K_SOME() antlr.TerminalNode
	K_STRUCT() antlr.TerminalNode
	K_TABLE() antlr.TerminalNode
	K_TABLESAMPLE() antlr.TerminalNode
	K_THEN() antlr.TerminalNode
	K_TO() antlr.TerminalNode
	K_TOKEN() antlr.TerminalNode
	K_TREAT() antlr.TerminalNode
	K_TRUE() antlr.TerminalNode
	K_TRUNCATE() antlr.TerminalNode
	K_UNBOUNDED() antlr.TerminalNode
	K_UNION() antlr.TerminalNode
	K_UNLOGGED() antlr.TerminalNode
	K_UNNEST() antlr.TerminalNode
	K_UPDATE() antlr.TerminalNode
	K_USE() antlr.TerminalNode
	K_USING() antlr.TerminalNode
	K_VIEW() antlr.TerminalNode
	K_WHEN() antlr.TerminalNode
	K_WHERE() antlr.TerminalNode
	K_WINDOW() antlr.TerminalNode
	K_WITH() antlr.TerminalNode
	K_WITHIN() antlr.TerminalNode

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

func (s *ReservedKeywordContext) K_ALL() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ALL, 0)
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

func (s *ReservedKeywordContext) K_ANY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ANY, 0)
}

func (s *ReservedKeywordContext) K_APPLY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_APPLY, 0)
}

func (s *ReservedKeywordContext) K_ARRAY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ARRAY, 0)
}

func (s *ReservedKeywordContext) K_AS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_AS, 0)
}

func (s *ReservedKeywordContext) K_ASC() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ASC, 0)
}

func (s *ReservedKeywordContext) K_ASSERT_ROWS_MODIFIED() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ASSERT_ROWS_MODIFIED, 0)
}

func (s *ReservedKeywordContext) K_AT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_AT, 0)
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

func (s *ReservedKeywordContext) K_BETWEEN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BETWEEN, 0)
}

func (s *ReservedKeywordContext) K_BY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BY, 0)
}

func (s *ReservedKeywordContext) K_CASE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CASE, 0)
}

func (s *ReservedKeywordContext) K_CAST() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CAST, 0)
}

func (s *ReservedKeywordContext) K_COLLATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_COLLATE, 0)
}

func (s *ReservedKeywordContext) K_COLUMNFAMILY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_COLUMNFAMILY, 0)
}

func (s *ReservedKeywordContext) K_CONTAINS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CONTAINS, 0)
}

func (s *ReservedKeywordContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CREATE, 0)
}

func (s *ReservedKeywordContext) K_CROSS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CROSS, 0)
}

func (s *ReservedKeywordContext) K_CUBE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CUBE, 0)
}

func (s *ReservedKeywordContext) K_CURRENT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CURRENT, 0)
}

func (s *ReservedKeywordContext) K_DEFAULT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DEFAULT, 0)
}

func (s *ReservedKeywordContext) K_DEFINE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DEFINE, 0)
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

func (s *ReservedKeywordContext) K_DISTINCT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DISTINCT, 0)
}

func (s *ReservedKeywordContext) K_DROP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DROP, 0)
}

func (s *ReservedKeywordContext) K_ELSE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ELSE, 0)
}

func (s *ReservedKeywordContext) K_END() antlr.TerminalNode {
	return s.GetToken(CqlParserK_END, 0)
}

func (s *ReservedKeywordContext) K_ENTRIES() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ENTRIES, 0)
}

func (s *ReservedKeywordContext) K_ENUM() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ENUM, 0)
}

func (s *ReservedKeywordContext) K_ESCAPE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ESCAPE, 0)
}

func (s *ReservedKeywordContext) K_EXCEPT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_EXCEPT, 0)
}

func (s *ReservedKeywordContext) K_EXCLUDE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_EXCLUDE, 0)
}

func (s *ReservedKeywordContext) K_EXECUTE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_EXECUTE, 0)
}

func (s *ReservedKeywordContext) K_EXISTS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_EXISTS, 0)
}

func (s *ReservedKeywordContext) K_EXTRACT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_EXTRACT, 0)
}

func (s *ReservedKeywordContext) K_FALSE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FALSE, 0)
}

func (s *ReservedKeywordContext) K_FETCH() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FETCH, 0)
}

func (s *ReservedKeywordContext) K_FOLLOWING() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FOLLOWING, 0)
}

func (s *ReservedKeywordContext) K_FOR() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FOR, 0)
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

func (s *ReservedKeywordContext) K_GROUP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_GROUP, 0)
}

func (s *ReservedKeywordContext) K_GROUPING() antlr.TerminalNode {
	return s.GetToken(CqlParserK_GROUPING, 0)
}

func (s *ReservedKeywordContext) K_GROUPS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_GROUPS, 0)
}

func (s *ReservedKeywordContext) K_HASH() antlr.TerminalNode {
	return s.GetToken(CqlParserK_HASH, 0)
}

func (s *ReservedKeywordContext) K_HAVING() antlr.TerminalNode {
	return s.GetToken(CqlParserK_HAVING, 0)
}

func (s *ReservedKeywordContext) K_IF() antlr.TerminalNode {
	return s.GetToken(CqlParserK_IF, 0)
}

func (s *ReservedKeywordContext) K_IGNORE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_IGNORE, 0)
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

func (s *ReservedKeywordContext) K_INNER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INNER, 0)
}

func (s *ReservedKeywordContext) K_INSERT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INSERT, 0)
}

func (s *ReservedKeywordContext) K_INTERSECT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INTERSECT, 0)
}

func (s *ReservedKeywordContext) K_INTERVAL() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INTERVAL, 0)
}

func (s *ReservedKeywordContext) K_INTO() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INTO, 0)
}

func (s *ReservedKeywordContext) K_IS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_IS, 0)
}

func (s *ReservedKeywordContext) K_JOIN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_JOIN, 0)
}

func (s *ReservedKeywordContext) K_KEYSPACE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_KEYSPACE, 0)
}

func (s *ReservedKeywordContext) K_LATERAL() antlr.TerminalNode {
	return s.GetToken(CqlParserK_LATERAL, 0)
}

func (s *ReservedKeywordContext) K_LEFT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_LEFT, 0)
}

func (s *ReservedKeywordContext) K_LIKE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_LIKE, 0)
}

func (s *ReservedKeywordContext) K_LIMIT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_LIMIT, 0)
}

func (s *ReservedKeywordContext) K_LOOKUP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_LOOKUP, 0)
}

func (s *ReservedKeywordContext) K_MATERIALIZED() antlr.TerminalNode {
	return s.GetToken(CqlParserK_MATERIALIZED, 0)
}

func (s *ReservedKeywordContext) K_MERGE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_MERGE, 0)
}

func (s *ReservedKeywordContext) K_MODIFY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_MODIFY, 0)
}

func (s *ReservedKeywordContext) K_NAN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NAN, 0)
}

func (s *ReservedKeywordContext) K_NATURAL() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NATURAL, 0)
}

func (s *ReservedKeywordContext) K_NEW() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NEW, 0)
}

func (s *ReservedKeywordContext) K_NO() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NO, 0)
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

func (s *ReservedKeywordContext) K_NULLS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NULLS, 0)
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

func (s *ReservedKeywordContext) K_OUTER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_OUTER, 0)
}

func (s *ReservedKeywordContext) K_OVER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_OVER, 0)
}

func (s *ReservedKeywordContext) K_PARTITION() antlr.TerminalNode {
	return s.GetToken(CqlParserK_PARTITION, 0)
}

func (s *ReservedKeywordContext) K_PRECEDING() antlr.TerminalNode {
	return s.GetToken(CqlParserK_PRECEDING, 0)
}

func (s *ReservedKeywordContext) K_PRIMARY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_PRIMARY, 0)
}

func (s *ReservedKeywordContext) K_PROTO() antlr.TerminalNode {
	return s.GetToken(CqlParserK_PROTO, 0)
}

func (s *ReservedKeywordContext) K_QUALIFY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_QUALIFY, 0)
}

func (s *ReservedKeywordContext) K_RANGE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_RANGE, 0)
}

func (s *ReservedKeywordContext) K_RECURSIVE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_RECURSIVE, 0)
}

func (s *ReservedKeywordContext) K_RENAME() antlr.TerminalNode {
	return s.GetToken(CqlParserK_RENAME, 0)
}

func (s *ReservedKeywordContext) K_REPLACE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_REPLACE, 0)
}

func (s *ReservedKeywordContext) K_RESPECT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_RESPECT, 0)
}

func (s *ReservedKeywordContext) K_REVOKE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_REVOKE, 0)
}

func (s *ReservedKeywordContext) K_RIGHT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_RIGHT, 0)
}

func (s *ReservedKeywordContext) K_ROLLUP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ROLLUP, 0)
}

func (s *ReservedKeywordContext) K_ROWS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ROWS, 0)
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

func (s *ReservedKeywordContext) K_SOME() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SOME, 0)
}

func (s *ReservedKeywordContext) K_STRUCT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_STRUCT, 0)
}

func (s *ReservedKeywordContext) K_TABLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TABLE, 0)
}

func (s *ReservedKeywordContext) K_TABLESAMPLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TABLESAMPLE, 0)
}

func (s *ReservedKeywordContext) K_THEN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_THEN, 0)
}

func (s *ReservedKeywordContext) K_TO() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TO, 0)
}

func (s *ReservedKeywordContext) K_TOKEN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TOKEN, 0)
}

func (s *ReservedKeywordContext) K_TREAT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TREAT, 0)
}

func (s *ReservedKeywordContext) K_TRUE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TRUE, 0)
}

func (s *ReservedKeywordContext) K_TRUNCATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TRUNCATE, 0)
}

func (s *ReservedKeywordContext) K_UNBOUNDED() antlr.TerminalNode {
	return s.GetToken(CqlParserK_UNBOUNDED, 0)
}

func (s *ReservedKeywordContext) K_UNION() antlr.TerminalNode {
	return s.GetToken(CqlParserK_UNION, 0)
}

func (s *ReservedKeywordContext) K_UNLOGGED() antlr.TerminalNode {
	return s.GetToken(CqlParserK_UNLOGGED, 0)
}

func (s *ReservedKeywordContext) K_UNNEST() antlr.TerminalNode {
	return s.GetToken(CqlParserK_UNNEST, 0)
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

func (s *ReservedKeywordContext) K_VIEW() antlr.TerminalNode {
	return s.GetToken(CqlParserK_VIEW, 0)
}

func (s *ReservedKeywordContext) K_WHEN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_WHEN, 0)
}

func (s *ReservedKeywordContext) K_WHERE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_WHERE, 0)
}

func (s *ReservedKeywordContext) K_WINDOW() antlr.TerminalNode {
	return s.GetToken(CqlParserK_WINDOW, 0)
}

func (s *ReservedKeywordContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(CqlParserK_WITH, 0)
}

func (s *ReservedKeywordContext) K_WITHIN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_WITHIN, 0)
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
	p.EnterRule(localctx, 48, CqlParserRULE_reservedKeyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(540)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8070521870123274246) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&4538218011079409611) != 0) || ((int64((_la-129)) & ^0x3f) == 0 && ((int64(1)<<(_la-129))&-4750187407873683969) != 0) || _la == CqlParserK_VIEW) {
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
	K_ASCII() antlr.TerminalNode
	K_BIGINT() antlr.TerminalNode
	K_BLOB() antlr.TerminalNode
	K_BOOLEAN() antlr.TerminalNode
	K_CALLED() antlr.TerminalNode
	K_CLUSTERING() antlr.TerminalNode
	K_COMPACT() antlr.TerminalNode
	K_COUNT() antlr.TerminalNode
	K_COUNTER() antlr.TerminalNode
	K_CUSTOM() antlr.TerminalNode
	K_DATE() antlr.TerminalNode
	K_DECIMAL() antlr.TerminalNode
	K_DOUBLE() antlr.TerminalNode
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

func (s *NonReservedKeywordContext) K_DOUBLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DOUBLE, 0)
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
	p.EnterRule(localctx, 50, CqlParserRULE_nonReservedKeyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(542)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8070521870123274244) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&8088817534084923405) != 0) || ((int64((_la-138)) & ^0x3f) == 0 && ((int64(1)<<(_la-138))&9277709781003289) != 0)) {
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
