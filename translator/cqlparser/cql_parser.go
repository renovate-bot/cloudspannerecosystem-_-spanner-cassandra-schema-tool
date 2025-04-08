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
		"", "'CREATE'", "'TABLE'", "'IF'", "'NOT'", "'EXISTS'", "'PRIMARY'",
		"'KEY'", "'WITH'", "'ASCII'", "'BIGINT'", "'BLOB'", "'BOOLEAN'", "'COUNTER'",
		"'DATE'", "'DECIMAL'", "'DOUBLE'", "'FLOAT'", "'INET'", "'INT'", "'SMALLINT'",
		"'TEXT'", "'TIME'", "'TIMESTAMP'", "'TIMEUUID'", "'TINYINT'", "'UUID'",
		"'VARCHAR'", "'VARINT'", "'MAP'", "'SET'", "'LIST'", "';'", "'\"'",
		"'.'", "','", "'('", "')'", "'<'", "'>'",
	}
	staticData.SymbolicNames = []string{
		"", "K_CREATE", "K_TABLE", "K_IF", "K_NOT", "K_EXISTS", "K_PRIMARY",
		"K_KEY", "K_WITH", "K_ASCII", "K_BIGINT", "K_BLOB", "K_BOOLEAN", "K_COUNTER",
		"K_DATE", "K_DECIMAL", "K_DOUBLE", "K_FLOAT", "K_INET", "K_INT", "K_SMALLINT",
		"K_TEXT", "K_TIME", "K_TIMESTAMP", "K_TIMEUUID", "K_TINYINT", "K_UUID",
		"K_VARCHAR", "K_VARINT", "K_MAP", "K_SET", "K_LIST", "SEMICOLON", "DQUOTE",
		"DOT", "COMMA", "L_PAREN", "R_PAREN", "L_ANGLE_BRACKET", "R_ANGLE_BRACKET",
		"IDENTIFIER", "IDENTIFIER_WITH_HYPHEN", "WS", "COMMENT", "MULTILINE_COMMENT",
		"UNKNOWN",
	}
	staticData.RuleNames = []string{
		"root", "cqlStatement", "createTable", "columnDefinitionList", "columnDefinition",
		"primaryKeyClause", "primaryKey", "partitionKey", "clusteringColumns",
		"columnName", "tableName", "keyspaceName", "generalIdentifier", "cqlType",
		"primaryKeyKeywords", "ifNotExist", "cqlNativeType", "cqlCollectionType",
		"wihtTableOptions", "nonSemicolonToken",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 45, 159, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 0, 3,
		0, 43, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 3, 2, 52, 8, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 59, 8, 2, 1, 3, 1, 3, 1, 3, 5, 3, 64,
		8, 3, 10, 3, 12, 3, 67, 9, 3, 1, 3, 1, 3, 3, 3, 71, 8, 3, 1, 4, 1, 4, 1,
		4, 3, 4, 76, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 3, 6,
		86, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 93, 8, 7, 10, 7, 12, 7, 96,
		9, 7, 1, 7, 1, 7, 3, 7, 100, 8, 7, 1, 8, 1, 8, 1, 8, 5, 8, 105, 8, 8, 10,
		8, 12, 8, 108, 9, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10, 115, 8, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 3, 13, 125, 8,
		13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 3, 17, 148, 8, 17, 1, 18, 1, 18, 5, 18, 152, 8, 18, 10, 18, 12, 18,
		155, 9, 18, 1, 19, 1, 19, 1, 19, 0, 0, 20, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 0, 4, 1, 0, 40, 41, 1, 0, 9,
		28, 1, 0, 30, 31, 3, 0, 1, 31, 33, 41, 45, 45, 152, 0, 40, 1, 0, 0, 0,
		2, 46, 1, 0, 0, 0, 4, 48, 1, 0, 0, 0, 6, 60, 1, 0, 0, 0, 8, 72, 1, 0, 0,
		0, 10, 77, 1, 0, 0, 0, 12, 82, 1, 0, 0, 0, 14, 99, 1, 0, 0, 0, 16, 101,
		1, 0, 0, 0, 18, 109, 1, 0, 0, 0, 20, 114, 1, 0, 0, 0, 22, 118, 1, 0, 0,
		0, 24, 120, 1, 0, 0, 0, 26, 124, 1, 0, 0, 0, 28, 126, 1, 0, 0, 0, 30, 129,
		1, 0, 0, 0, 32, 133, 1, 0, 0, 0, 34, 147, 1, 0, 0, 0, 36, 149, 1, 0, 0,
		0, 38, 156, 1, 0, 0, 0, 40, 42, 3, 2, 1, 0, 41, 43, 5, 32, 0, 0, 42, 41,
		1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 45, 5, 0, 0, 1,
		45, 1, 1, 0, 0, 0, 46, 47, 3, 4, 2, 0, 47, 3, 1, 0, 0, 0, 48, 49, 5, 1,
		0, 0, 49, 51, 5, 2, 0, 0, 50, 52, 3, 30, 15, 0, 51, 50, 1, 0, 0, 0, 51,
		52, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 54, 3, 20, 10, 0, 54, 55, 5, 36,
		0, 0, 55, 56, 3, 6, 3, 0, 56, 58, 5, 37, 0, 0, 57, 59, 3, 36, 18, 0, 58,
		57, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 5, 1, 0, 0, 0, 60, 65, 3, 8, 4,
		0, 61, 62, 5, 35, 0, 0, 62, 64, 3, 8, 4, 0, 63, 61, 1, 0, 0, 0, 64, 67,
		1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 70, 1, 0, 0, 0,
		67, 65, 1, 0, 0, 0, 68, 69, 5, 35, 0, 0, 69, 71, 3, 10, 5, 0, 70, 68, 1,
		0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 7, 1, 0, 0, 0, 72, 73, 3, 18, 9, 0, 73,
		75, 3, 26, 13, 0, 74, 76, 3, 28, 14, 0, 75, 74, 1, 0, 0, 0, 75, 76, 1,
		0, 0, 0, 76, 9, 1, 0, 0, 0, 77, 78, 3, 28, 14, 0, 78, 79, 5, 36, 0, 0,
		79, 80, 3, 12, 6, 0, 80, 81, 5, 37, 0, 0, 81, 11, 1, 0, 0, 0, 82, 85, 3,
		14, 7, 0, 83, 84, 5, 35, 0, 0, 84, 86, 3, 16, 8, 0, 85, 83, 1, 0, 0, 0,
		85, 86, 1, 0, 0, 0, 86, 13, 1, 0, 0, 0, 87, 100, 3, 18, 9, 0, 88, 89, 5,
		36, 0, 0, 89, 94, 3, 18, 9, 0, 90, 91, 5, 35, 0, 0, 91, 93, 3, 18, 9, 0,
		92, 90, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1,
		0, 0, 0, 95, 97, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 98, 5, 37, 0, 0, 98,
		100, 1, 0, 0, 0, 99, 87, 1, 0, 0, 0, 99, 88, 1, 0, 0, 0, 100, 15, 1, 0,
		0, 0, 101, 106, 3, 18, 9, 0, 102, 103, 5, 35, 0, 0, 103, 105, 3, 18, 9,
		0, 104, 102, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106,
		107, 1, 0, 0, 0, 107, 17, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109, 110, 5,
		40, 0, 0, 110, 19, 1, 0, 0, 0, 111, 112, 3, 22, 11, 0, 112, 113, 5, 34,
		0, 0, 113, 115, 1, 0, 0, 0, 114, 111, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0,
		115, 116, 1, 0, 0, 0, 116, 117, 5, 40, 0, 0, 117, 21, 1, 0, 0, 0, 118,
		119, 3, 24, 12, 0, 119, 23, 1, 0, 0, 0, 120, 121, 7, 0, 0, 0, 121, 25,
		1, 0, 0, 0, 122, 125, 3, 32, 16, 0, 123, 125, 3, 34, 17, 0, 124, 122, 1,
		0, 0, 0, 124, 123, 1, 0, 0, 0, 125, 27, 1, 0, 0, 0, 126, 127, 5, 6, 0,
		0, 127, 128, 5, 7, 0, 0, 128, 29, 1, 0, 0, 0, 129, 130, 5, 3, 0, 0, 130,
		131, 5, 4, 0, 0, 131, 132, 5, 5, 0, 0, 132, 31, 1, 0, 0, 0, 133, 134, 7,
		1, 0, 0, 134, 33, 1, 0, 0, 0, 135, 136, 5, 29, 0, 0, 136, 137, 5, 38, 0,
		0, 137, 138, 3, 32, 16, 0, 138, 139, 5, 35, 0, 0, 139, 140, 3, 32, 16,
		0, 140, 141, 5, 39, 0, 0, 141, 148, 1, 0, 0, 0, 142, 143, 7, 2, 0, 0, 143,
		144, 5, 38, 0, 0, 144, 145, 3, 32, 16, 0, 145, 146, 5, 39, 0, 0, 146, 148,
		1, 0, 0, 0, 147, 135, 1, 0, 0, 0, 147, 142, 1, 0, 0, 0, 148, 35, 1, 0,
		0, 0, 149, 153, 5, 8, 0, 0, 150, 152, 3, 38, 19, 0, 151, 150, 1, 0, 0,
		0, 152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154,
		37, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 157, 7, 3, 0, 0, 157, 39, 1,
		0, 0, 0, 14, 42, 51, 58, 65, 70, 75, 85, 94, 99, 106, 114, 124, 147, 153,
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
	CqlParserK_CREATE               = 1
	CqlParserK_TABLE                = 2
	CqlParserK_IF                   = 3
	CqlParserK_NOT                  = 4
	CqlParserK_EXISTS               = 5
	CqlParserK_PRIMARY              = 6
	CqlParserK_KEY                  = 7
	CqlParserK_WITH                 = 8
	CqlParserK_ASCII                = 9
	CqlParserK_BIGINT               = 10
	CqlParserK_BLOB                 = 11
	CqlParserK_BOOLEAN              = 12
	CqlParserK_COUNTER              = 13
	CqlParserK_DATE                 = 14
	CqlParserK_DECIMAL              = 15
	CqlParserK_DOUBLE               = 16
	CqlParserK_FLOAT                = 17
	CqlParserK_INET                 = 18
	CqlParserK_INT                  = 19
	CqlParserK_SMALLINT             = 20
	CqlParserK_TEXT                 = 21
	CqlParserK_TIME                 = 22
	CqlParserK_TIMESTAMP            = 23
	CqlParserK_TIMEUUID             = 24
	CqlParserK_TINYINT              = 25
	CqlParserK_UUID                 = 26
	CqlParserK_VARCHAR              = 27
	CqlParserK_VARINT               = 28
	CqlParserK_MAP                  = 29
	CqlParserK_SET                  = 30
	CqlParserK_LIST                 = 31
	CqlParserSEMICOLON              = 32
	CqlParserDQUOTE                 = 33
	CqlParserDOT                    = 34
	CqlParserCOMMA                  = 35
	CqlParserL_PAREN                = 36
	CqlParserR_PAREN                = 37
	CqlParserL_ANGLE_BRACKET        = 38
	CqlParserR_ANGLE_BRACKET        = 39
	CqlParserIDENTIFIER             = 40
	CqlParserIDENTIFIER_WITH_HYPHEN = 41
	CqlParserWS                     = 42
	CqlParserCOMMENT                = 43
	CqlParserMULTILINE_COMMENT      = 44
	CqlParserUNKNOWN                = 45
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
	CqlParserRULE_cqlType              = 13
	CqlParserRULE_primaryKeyKeywords   = 14
	CqlParserRULE_ifNotExist           = 15
	CqlParserRULE_cqlNativeType        = 16
	CqlParserRULE_cqlCollectionType    = 17
	CqlParserRULE_wihtTableOptions     = 18
	CqlParserRULE_nonSemicolonToken    = 19
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
		p.SetState(40)
		p.CqlStatement()
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserSEMICOLON {
		{
			p.SetState(41)
			p.Match(CqlParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(44)
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
		p.SetState(46)
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
		p.SetState(48)
		p.Match(CqlParserK_CREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(CqlParserK_TABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_IF {
		{
			p.SetState(50)
			p.IfNotExist()
		}

	}
	{
		p.SetState(53)
		p.TableName()
	}
	{
		p.SetState(54)
		p.Match(CqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.ColumnDefinitionList()
	}
	{
		p.SetState(56)
		p.Match(CqlParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_WITH {
		{
			p.SetState(57)
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
		p.SetState(60)
		p.ColumnDefinition()
	}
	p.SetState(65)
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
				p.SetState(61)
				p.Match(CqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(62)
				p.ColumnDefinition()
			}

		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserCOMMA {
		{
			p.SetState(68)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
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
		p.SetState(72)
		p.ColumnName()
	}
	{
		p.SetState(73)
		p.CqlType()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_PRIMARY {
		{
			p.SetState(74)
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
		p.SetState(77)
		p.PrimaryKeyKeywords()
	}
	{
		p.SetState(78)
		p.Match(CqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.PrimaryKey()
	}
	{
		p.SetState(80)
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
		p.SetState(82)
		p.PartitionKey()
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserCOMMA {
		{
			p.SetState(83)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
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

	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.ColumnName()
		}

	case CqlParserL_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)
			p.Match(CqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.ColumnName()
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CqlParserCOMMA {
			{
				p.SetState(90)
				p.Match(CqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(91)
				p.ColumnName()
			}

			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(97)
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
		p.SetState(101)
		p.ColumnName()
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CqlParserCOMMA {
		{
			p.SetState(102)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.ColumnName()
		}

		p.SetState(108)
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
	IDENTIFIER() antlr.TerminalNode

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

func (s *ColumnNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CqlParserIDENTIFIER, 0)
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
		p.SetState(109)
		p.Match(CqlParserIDENTIFIER)
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

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
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

func (s *TableNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CqlParserIDENTIFIER, 0)
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
	p.SetState(114)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(111)
			p.KeyspaceName()
		}
		{
			p.SetState(112)
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
		p.SetState(116)
		p.Match(CqlParserIDENTIFIER)
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
		p.SetState(118)
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
	IDENTIFIER() antlr.TerminalNode

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

func (s *GeneralIdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CqlParserIDENTIFIER, 0)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CqlParserIDENTIFIER || _la == CqlParserIDENTIFIER_WITH_HYPHEN) {
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
	p.EnterRule(localctx, 26, CqlParserRULE_cqlType)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_ASCII, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_COUNTER, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DOUBLE, CqlParserK_FLOAT, CqlParserK_INET, CqlParserK_INT, CqlParserK_SMALLINT, CqlParserK_TEXT, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_UUID, CqlParserK_VARCHAR, CqlParserK_VARINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.CqlNativeType()
		}

	case CqlParserK_MAP, CqlParserK_SET, CqlParserK_LIST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
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
	p.EnterRule(localctx, 28, CqlParserRULE_primaryKeyKeywords)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(CqlParserK_PRIMARY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
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
	p.EnterRule(localctx, 30, CqlParserRULE_ifNotExist)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(CqlParserK_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Match(CqlParserK_NOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
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
	p.EnterRule(localctx, 32, CqlParserRULE_cqlNativeType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&536870400) != 0) {
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
	p.EnterRule(localctx, 34, CqlParserRULE_cqlCollectionType)
	var _la int

	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_MAP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.Match(CqlParserK_MAP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(CqlParserL_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.CqlNativeType()
		}
		{
			p.SetState(138)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.CqlNativeType()
		}
		{
			p.SetState(140)
			p.Match(CqlParserR_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserK_SET, CqlParserK_LIST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CqlParserK_SET || _la == CqlParserK_LIST) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(143)
			p.Match(CqlParserL_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.CqlNativeType()
		}
		{
			p.SetState(145)
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
	p.EnterRule(localctx, 36, CqlParserRULE_wihtTableOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(CqlParserK_WITH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&39578123632638) != 0 {
		{
			p.SetState(150)
			p.NonSemicolonToken()
		}

		p.SetState(155)
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
	K_CREATE() antlr.TerminalNode
	K_TABLE() antlr.TerminalNode
	K_IF() antlr.TerminalNode
	K_NOT() antlr.TerminalNode
	K_EXISTS() antlr.TerminalNode
	K_PRIMARY() antlr.TerminalNode
	K_KEY() antlr.TerminalNode
	K_WITH() antlr.TerminalNode
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
	K_MAP() antlr.TerminalNode
	K_SET() antlr.TerminalNode
	K_LIST() antlr.TerminalNode
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

func (s *NonSemicolonTokenContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CREATE, 0)
}

func (s *NonSemicolonTokenContext) K_TABLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TABLE, 0)
}

func (s *NonSemicolonTokenContext) K_IF() antlr.TerminalNode {
	return s.GetToken(CqlParserK_IF, 0)
}

func (s *NonSemicolonTokenContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NOT, 0)
}

func (s *NonSemicolonTokenContext) K_EXISTS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_EXISTS, 0)
}

func (s *NonSemicolonTokenContext) K_PRIMARY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_PRIMARY, 0)
}

func (s *NonSemicolonTokenContext) K_KEY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_KEY, 0)
}

func (s *NonSemicolonTokenContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(CqlParserK_WITH, 0)
}

func (s *NonSemicolonTokenContext) K_ASCII() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ASCII, 0)
}

func (s *NonSemicolonTokenContext) K_BIGINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BIGINT, 0)
}

func (s *NonSemicolonTokenContext) K_BLOB() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BLOB, 0)
}

func (s *NonSemicolonTokenContext) K_BOOLEAN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BOOLEAN, 0)
}

func (s *NonSemicolonTokenContext) K_COUNTER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_COUNTER, 0)
}

func (s *NonSemicolonTokenContext) K_DATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DATE, 0)
}

func (s *NonSemicolonTokenContext) K_DECIMAL() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DECIMAL, 0)
}

func (s *NonSemicolonTokenContext) K_DOUBLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DOUBLE, 0)
}

func (s *NonSemicolonTokenContext) K_FLOAT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FLOAT, 0)
}

func (s *NonSemicolonTokenContext) K_INET() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INET, 0)
}

func (s *NonSemicolonTokenContext) K_INT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INT, 0)
}

func (s *NonSemicolonTokenContext) K_SMALLINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SMALLINT, 0)
}

func (s *NonSemicolonTokenContext) K_TEXT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TEXT, 0)
}

func (s *NonSemicolonTokenContext) K_TIME() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIME, 0)
}

func (s *NonSemicolonTokenContext) K_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIMESTAMP, 0)
}

func (s *NonSemicolonTokenContext) K_TIMEUUID() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIMEUUID, 0)
}

func (s *NonSemicolonTokenContext) K_TINYINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TINYINT, 0)
}

func (s *NonSemicolonTokenContext) K_UUID() antlr.TerminalNode {
	return s.GetToken(CqlParserK_UUID, 0)
}

func (s *NonSemicolonTokenContext) K_VARCHAR() antlr.TerminalNode {
	return s.GetToken(CqlParserK_VARCHAR, 0)
}

func (s *NonSemicolonTokenContext) K_VARINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_VARINT, 0)
}

func (s *NonSemicolonTokenContext) K_MAP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_MAP, 0)
}

func (s *NonSemicolonTokenContext) K_SET() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SET, 0)
}

func (s *NonSemicolonTokenContext) K_LIST() antlr.TerminalNode {
	return s.GetToken(CqlParserK_LIST, 0)
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
	p.EnterRule(localctx, 38, CqlParserRULE_nonSemicolonToken)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&39578123632638) != 0) {
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
