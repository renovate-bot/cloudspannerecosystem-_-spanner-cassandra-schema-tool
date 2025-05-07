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
 
/**
 * Grammar for CQL.
 */
parser grammar CqlParser;

options {
  tokenVocab = CqlLexer;
}

root
    : cqlStatement SEMICOLON? EOF
    ;

cqlStatement
    : createTable
	| unsupportedStatement
    ;

createTable
   : K_CREATE K_TABLE ifNotExists? tableName L_PAREN columnDefinitionList R_PAREN wihtTableOptions?
   ;

// Reference: https://cassandra.apache.org/doc/4.0/cassandra/cql/definitions.html#statements.
// The syntax is not precisely defined; it is a simplification.
unsupportedStatement
	: K_USE keyspaceName
	| K_CREATE K_KEYSPACE ifNotExists? keyspaceName K_WITH nonSemicolonToken*
	| K_ALTER K_KEYSPACE keyspaceName K_WITH nonSemicolonToken*
	| K_DROP K_KEYSPACE ifExists? keyspaceName
	| K_ALTER K_TABLE tableName nonSemicolonToken* 
	| K_DROP K_TABLE ifExists? tableName
	| K_TRUNCATE K_TABLE? tableName

	| K_SELECT nonSemicolonToken* 
	| K_INSERT K_INTO nonSemicolonToken* 
	| K_UPDATE tableName nonSemicolonToken* 
	| K_DELETE nonSemicolonToken* 
	| K_BEGIN (K_UNLOGGED | K_COUNTER)? K_BATCH (nonSemicolonToken* SEMICOLON)+ K_APPLY K_BATCH SEMICOLON

	| K_CREATE K_CUSTOM? K_INDEX nonSemicolonToken*
	| K_DROP K_INDEX ifExists? nonSemicolonToken*

	| K_CREATE K_MATERIALIZED K_VIEW ifNotExists? nonSemicolonToken*
	| K_ALTER K_MATERIALIZED K_VIEW nonSemicolonToken*
	| K_DROP K_MATERIALIZED K_VIEW ifExists? nonSemicolonToken*

	| K_CREATE K_ROLE ifNotExists? nonSemicolonToken*
	| K_ALTER K_ROLE nonSemicolonToken*
	| K_DROP K_ROLE nonSemicolonToken*
	| K_GRANT nonSemicolonToken*
	| K_REVOKE nonSemicolonToken*
	| K_LIST nonSemicolonToken*
	| K_CREATE K_USER nonSemicolonToken*
	| K_ALTER K_USER nonSemicolonToken*
	| K_DROP K_USER nonSemicolonToken*

	| K_CREATE (K_OR K_REPLACE)? K_FUNCTION ifNotExists? nonSemicolonToken* K_AS SQUOTE nonSemicolonToken* SEMICOLON SQUOTE
	| K_DROP K_FUNCTION ifExists? nonSemicolonToken*
	| K_CREATE (K_OR K_REPLACE)? K_AGGREGATE ifNotExists? nonSemicolonToken*
	| K_DROP K_AGGREGATE ifExists? nonSemicolonToken*

	| K_CREATE K_TYPE ifNotExists? nonSemicolonToken*
	| K_ALTER K_TYPE nonSemicolonToken*
	| K_DROP K_TYPE ifExists? nonSemicolonToken*

	| K_CREATE K_TRIGGER ifNotExists? nonSemicolonToken*
	| K_DROP K_TRIGGER ifExists? nonSemicolonToken*
	;

columnDefinitionList
   : columnDefinition (COMMA columnDefinition)* (COMMA primaryKeyClause)?
   ;

columnDefinition
   : columnName cqlType primaryKeyKeywords?
   ;

// TODO: Merge the rule into columnDefinition.
primaryKeyClause
   : primaryKeyKeywords L_PAREN primaryKey R_PAREN
   ;

primaryKey
    : partitionKey (COMMA clusteringColumns)?
    ;

partitionKey
    : columnName
    | L_PAREN columnName (COMMA columnName)* R_PAREN
    ;

clusteringColumns
    : columnName (COMMA columnName)*
    ;

columnName
    : identifier
    ;

tableName
    : (keyspaceName DOT)? identifier
    ;

// A Cassandra keyspace name is conventionally defined as an IDENTIFIER. Spanner
// database names, however, permit hyphens. To ensure compatibility when mapping
// Spanner databases to Cassandra keyspaces, the parsing rule is relaxed to
// allow hyphens in keyspace names.
keyspaceName
    : generalIdentifier
    ;

generalIdentifier
    : IDENTIFIER_WITH_HYPHEN
    | identifier
    ;

identifier
    : nonReservedKeyword
    | IDENTIFIER
    ;

// TODO: Add support on frozen, duration, and tuple.
cqlType
    : cqlNativeType
    | cqlCollectionType
    ;

primaryKeyKeywords
    : K_PRIMARY K_KEY
    ;

ifExists
	: K_IF K_EXISTS
	;

ifNotExists
   : K_IF K_NOT K_EXISTS
   ;

cqlNativeType
    : K_ASCII
    | K_BIGINT
    | K_BLOB
    | K_BOOLEAN
    | K_COUNTER
    | K_DATE
    | K_DECIMAL
    | K_DOUBLE
    | K_FLOAT
    | K_INET
    | K_INT
    | K_SMALLINT
    | K_TEXT
    | K_TIME
    | K_TIMESTAMP
    | K_TIMEUUID
    | K_TINYINT
    | K_UUID
    | K_VARCHAR
    | K_VARINT
    ;

cqlCollectionType
    : K_MAP L_ANGLE_BRACKET cqlNativeType COMMA cqlNativeType R_ANGLE_BRACKET
    | (K_SET | K_LIST) L_ANGLE_BRACKET cqlNativeType R_ANGLE_BRACKET
    ;

// TODO: Consider using a more precise syntax.
wihtTableOptions
    : K_WITH nonSemicolonToken*
    ;

nonSemicolonToken
    : keyword
	| SQUOTE
    | DQUOTE
    | DOT
    | COMMA
    | L_PAREN
    | R_PAREN
    | L_ANGLE_BRACKET
    | R_ANGLE_BRACKET
    | IDENTIFIER
    | IDENTIFIER_WITH_HYPHEN
    | UNKNOWN
    ;

keyword
    : reservedKeyword
    | nonReservedKeyword
    ;

reservedKeyword
	: K_ADD
	| K_ALL
	| K_ALLOW
	| K_ALTER
	| K_AND
	| K_ANY
	| K_APPLY
	| K_ARRAY
	| K_AS
	| K_ASC
	| K_ASSERT_ROWS_MODIFIED
	| K_AT
	| K_AUTHORIZE
	| K_BATCH
	| K_BEGIN
	| K_BETWEEN
	| K_BY
	| K_CASE
	| K_CAST
	| K_COLLATE
	| K_COLUMNFAMILY
	| K_CONTAINS
	| K_CREATE
	| K_CROSS
	| K_CUBE
	| K_CURRENT
	| K_DEFAULT
	| K_DEFINE
	| K_DELETE
	| K_DESC
	| K_DESCRIBE
	| K_DISTINCT
	| K_DROP
	| K_ELSE
	| K_END
	| K_ENTRIES
	| K_ENUM
	| K_ESCAPE
	| K_EXCEPT
	| K_EXCLUDE
	| K_EXECUTE
	| K_EXISTS
	| K_EXTRACT
	| K_FALSE
	| K_FETCH
	| K_FOLLOWING
	| K_FOR
	| K_FROM
	| K_FULL
	| K_GRANT
	| K_GROUP
	| K_GROUPING
	| K_GROUPS
	| K_HASH
	| K_HAVING
	| K_IF
	| K_IGNORE
	| K_IN
	| K_INDEX
	| K_INFINITY
	| K_INNER
	| K_INSERT
	| K_INTERSECT
	| K_INTERVAL
	| K_INTO
	| K_IS
	| K_JOIN
	| K_KEYSPACE
	| K_LATERAL
	| K_LEFT
	| K_LIKE
	| K_LIMIT
	| K_LOOKUP
	| K_MATERIALIZED
	| K_MERGE
	| K_MODIFY
	| K_NAN
	| K_NATURAL
	| K_NEW
	| K_NO
	| K_NORECURSIVE
	| K_NOT
	| K_NULL
	| K_NULLS
	| K_OF
	| K_ON
	| K_OR
	| K_ORDER
	| K_OUTER
	| K_OVER
	| K_PARTITION
	| K_PRECEDING
	| K_PRIMARY
	| K_PROTO
	| K_QUALIFY
	| K_RANGE
	| K_RECURSIVE
	| K_RENAME
	| K_REPLACE
	| K_RESPECT
	| K_REVOKE
	| K_RIGHT
	| K_ROLLUP
	| K_ROWS
	| K_SCHEMA
	| K_SELECT
	| K_SET
	| K_SOME
	| K_STRUCT
	| K_TABLE
	| K_TABLESAMPLE
	| K_THEN
	| K_TO
	| K_TOKEN
	| K_TREAT
	| K_TRUE
	| K_TRUNCATE
	| K_UNBOUNDED
	| K_UNION
	| K_UNLOGGED
	| K_UNNEST
	| K_UPDATE
	| K_USE
	| K_USING
	| K_VIEW
	| K_WHEN
	| K_WHERE
	| K_WINDOW
	| K_WITH
	| K_WITHIN
	;

nonReservedKeyword
	: K_AGGREGATE
	| K_ASCII
	| K_BIGINT
	| K_BLOB
	| K_BOOLEAN
	| K_CALLED
	| K_CLUSTERING
	| K_COMPACT
	| K_COUNT
	| K_COUNTER
	| K_CUSTOM
	| K_DATE
	| K_DECIMAL
	| K_DOUBLE
	| K_FILTERING
	| K_FINALFUNC
	| K_FLOAT
	| K_FROZEN
	| K_FUNCTION
	| K_FUNCTIONS
	| K_INET
	| K_INITCOND
	| K_INPUT
	| K_INT
	| K_JSON
	| K_KEY
	| K_KEYS
	| K_KEYSPACES
	| K_LANGUAGE
	| K_LIST
	| K_LOGIN
	| K_MAP
	| K_NOLOGIN
	| K_NOSUPERUSER
	| K_OPTIONS
	| K_PASSWORD
	| K_PERMISSION
	| K_PERMISSIONS
	| K_RETURNS
	| K_ROLE
	| K_ROLES
	| K_SFUNC
	| K_SMALLINT
	| K_STATIC
	| K_STORAGE
	| K_STYPE
	| K_SUPERUSER
	| K_TEXT
	| K_TIME
	| K_TIMESTAMP
	| K_TIMEUUID
	| K_TINYINT
	| K_TRIGGER
	| K_TTL
	| K_TUPLE
	| K_TYPE
	| K_USER
	| K_USERS
	| K_UUID
	| K_VALUES
	| K_VARCHAR
	| K_VARINT
	| K_WRITETIME
	;