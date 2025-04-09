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
 * Lexical definitions for CQL.
 */
lexer grammar CqlLexer;

options {
  caseInsensitive = true;
}

fragment ALPHA: [A-Z];

// Keywords (Reference: https://cassandra.apache.org/doc/stable/cassandra/cql/appendices.html#appendix-A)
// When adding a new keyword, add entry to reservedKeyword or nonReservedKeyword in the CqlParser.g4.
K_ADD          : 'ADD'; // reserved
K_AGGREGATE    : 'AGGREGATE'; // non-reserved
K_ALL          : 'ALL'; // non-reserved
K_ALLOW        : 'ALLOW'; // reserved
K_ALTER        : 'ALTER'; // reserved
K_AND          : 'AND'; // reserved
K_APPLY        : 'APPLY'; // reserved
K_AS           : 'AS'; // non-reserved
K_ASC          : 'ASC'; // reserved
K_ASCII        : 'ASCII'; // non-reserved
K_AUTHORIZE    : 'AUTHORIZE'; // reserved
K_BATCH        : 'BATCH'; // reserved
K_BEGIN        : 'BEGIN'; // reserved
K_BIGINT       : 'BIGINT'; // non-reserved
K_BLOB         : 'BLOB'; // non-reserved
K_BOOLEAN      : 'BOOLEAN'; // non-reserved
K_BY           : 'BY'; // reserved
K_CALLED       : 'CALLED'; // non-reserved
K_CLUSTERING   : 'CLUSTERING'; // non-reserved
K_COLUMNFAMILY : 'COLUMNFAMILY'; // reserved
K_COMPACT      : 'COMPACT'; // non-reserved
K_CONTAINS     : 'CONTAINS'; // non-reserved
K_COUNT        : 'COUNT'; // non-reserved
K_COUNTER      : 'COUNTER'; // non-reserved
K_CREATE       : 'CREATE'; // reserved
K_CUSTOM       : 'CUSTOM'; // non-reserved
K_DATE         : 'DATE'; // non-reserved
K_DECIMAL      : 'DECIMAL'; // non-reserved
K_DELETE       : 'DELETE'; // reserved
K_DESC         : 'DESC'; // reserved
K_DESCRIBE     : 'DESCRIBE'; // reserved
K_DISTINCT     : 'DISTINCT'; // non-reserved
K_DOUBLE       : 'DOUBLE'; // non-reserved
K_DROP         : 'DROP'; // reserved
K_ENTRIES      : 'ENTRIES'; // reserved
K_EXECUTE      : 'EXECUTE'; // reserved
K_EXISTS       : 'EXISTS'; // non-reserved
K_FILTERING    : 'FILTERING'; // non-reserved
K_FINALFUNC    : 'FINALFUNC'; // non-reserved
K_FLOAT        : 'FLOAT'; // non-reserved
K_FROM         : 'FROM'; // reserved
K_FROZEN       : 'FROZEN'; // non-reserved
K_FULL         : 'FULL'; // reserved
K_FUNCTION     : 'FUNCTION'; // non-reserved
K_FUNCTIONS    : 'FUNCTIONS'; // non-reserved
K_GRANT        : 'GRANT'; // reserved
K_IF           : 'IF'; // reserved
K_IN           : 'IN'; // reserved
K_INDEX        : 'INDEX'; // reserved
K_INET         : 'INET'; // non-reserved
K_INFINITY     : 'INFINITY'; // reserved
K_INITCOND     : 'INITCOND'; // non-reserved
K_INPUT        : 'INPUT'; // non-reserved
K_INSERT       : 'INSERT'; // reserved
K_INT          : 'INT'; // non-reserved
K_INTO         : 'INTO'; // reserved
K_JSON         : 'JSON'; // non-reserved
K_KEY          : 'KEY'; // non-reserved
K_KEYS         : 'KEYS'; // non-reserved
K_KEYSPACE     : 'KEYSPACE'; // reserved
K_KEYSPACES    : 'KEYSPACES'; // non-reserved
K_LANGUAGE     : 'LANGUAGE'; // non-reserved
K_LIMIT        : 'LIMIT'; // reserved
K_LIST         : 'LIST'; // non-reserved
K_LOGIN        : 'LOGIN'; // non-reserved
K_MAP          : 'MAP'; // non-reserved
K_MODIFY       : 'MODIFY'; // reserved
K_NAN          : 'NAN'; // reserved
K_NOLOGIN      : 'NOLOGIN'; // non-reserved
K_NORECURSIVE  : 'NORECURSIVE'; // reserved
K_NOSUPERUSER  : 'NOSUPERUSER'; // non-reserved
K_NOT          : 'NOT'; // reserved
K_NULL         : 'NULL'; // reserved
K_OF           : 'OF'; // reserved
K_ON           : 'ON'; // reserved
K_OPTIONS      : 'OPTIONS'; // non-reserved
K_OR           : 'OR'; // reserved
K_ORDER        : 'ORDER'; // reserved
K_PASSWORD     : 'PASSWORD'; // non-reserved
K_PERMISSION   : 'PERMISSION'; // non-reserved
K_PERMISSIONS  : 'PERMISSIONS'; // non-reserved
K_PRIMARY      : 'PRIMARY'; // reserved
K_RENAME       : 'RENAME'; // reserved
K_REPLACE      : 'REPLACE'; // reserved
K_RETURNS      : 'RETURNS'; // non-reserved
K_REVOKE       : 'REVOKE'; // reserved
K_ROLE         : 'ROLE'; // non-reserved
K_ROLES        : 'ROLES'; // non-reserved
K_SCHEMA       : 'SCHEMA'; // reserved
K_SELECT       : 'SELECT'; // reserved
K_SET          : 'SET'; // reserved
K_SFUNC        : 'SFUNC'; // non-reserved
K_SMALLINT     : 'SMALLINT'; // non-reserved
K_STATIC       : 'STATIC'; // non-reserved
K_STORAGE      : 'STORAGE'; // non-reserved
K_STYPE        : 'STYPE'; // non-reserved
K_SUPERUSER    : 'SUPERUSER'; // non-reserved
K_TABLE        : 'TABLE'; // reserved
K_TEXT         : 'TEXT'; // non-reserved
K_TIME         : 'TIME'; // non-reserved
K_TIMESTAMP    : 'TIMESTAMP'; // non-reserved
K_TIMEUUID     : 'TIMEUUID'; // non-reserved
K_TINYINT      : 'TINYINT'; // non-reserved
K_TO           : 'TO'; // reserved
K_TOKEN        : 'TOKEN'; // reserved
K_TRIGGER      : 'TRIGGER'; // non-reserved
K_TRUNCATE     : 'TRUNCATE'; // reserved
K_TTL          : 'TTL'; // non-reserved
K_TUPLE        : 'TUPLE'; // non-reserved
K_TYPE         : 'TYPE'; // non-reserved
K_UNLOGGED     : 'UNLOGGED'; // reserved
K_UPDATE       : 'UPDATE'; // reserved
K_USE          : 'USE'; // reserved
K_USER         : 'USER'; // non-reserved
K_USERS        : 'USERS'; // non-reserved
K_USING        : 'USING'; // reserved
K_UUID         : 'UUID'; // non-reserved
K_VALUES       : 'VALUES'; // non-reserved
K_VARCHAR      : 'VARCHAR'; // non-reserved
K_VARINT       : 'VARINT'; // non-reserved
K_WHERE        : 'WHERE'; // reserved
K_WITH         : 'WITH'; // reserved
K_WRITETIME    : 'WRITETIME'; // non-reserved

// Punctuation
// When adding a new punctuation, add entry to nonSemicolonToken in the CqlParser.g4.
SEMICOLON       : ';';
DQUOTE          : '"';
DOT             : '.';
COMMA           : ',';
L_PAREN         : '(';
R_PAREN         : ')';
L_ANGLE_BRACKET : '<';
R_ANGLE_BRACKET : '>';

// Identifiers
IDENTIFIER: ALPHA[A-Z0-9_]*;
IDENTIFIER_WITH_HYPHEN: ALPHA [A-Z0-9_-]* '-' [A-Z0-9_-]*;

// Hidden
WS
    : [ \t\r\n]+ -> channel (HIDDEN)
    ;

COMMENT
    : ('-- ' | '//') .*? ('\n' | '\r') -> channel (HIDDEN)
    ;

MULTILINE_COMMENT
    : '/*' .*? '*/' -> channel (HIDDEN)
    ;

UNKNOWN
    : .
    ;