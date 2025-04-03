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
		"'KEY'", "'ASCII'", "'BIGINT'", "'BLOB'", "'BOOLEAN'", "'COUNTER'",
		"'DATE'", "'DECIMAL'", "'DOUBLE'", "'FLOAT'", "'INET'", "'INT'", "'SMALLINT'",
		"'TEXT'", "'TIME'", "'TIMESTAMP'", "'TIMEUUID'", "'TINYINT'", "'UUID'",
		"'VARCHAR'", "'VARINT'", "'MAP'", "'SET'", "'LIST'", "';'", "'\"'",
		"'.'", "','", "'('", "')'", "'<'", "'>'",
	}
	staticData.SymbolicNames = []string{
		"", "K_CREATE", "K_TABLE", "K_IF", "K_NOT", "K_EXISTS", "K_PRIMARY",
		"K_KEY", "K_ASCII", "K_BIGINT", "K_BLOB", "K_BOOLEAN", "K_COUNTER",
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
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 144, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 3, 0, 39, 8, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 3, 2, 48, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 5, 3, 58, 8, 3, 10, 3, 12, 3, 61, 9, 3, 1, 3, 1, 3, 3,
		3, 65, 8, 3, 1, 4, 1, 4, 1, 4, 3, 4, 70, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 3, 6, 80, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5,
		7, 87, 8, 7, 10, 7, 12, 7, 90, 9, 7, 1, 7, 1, 7, 3, 7, 94, 8, 7, 1, 8,
		1, 8, 1, 8, 5, 8, 99, 8, 8, 10, 8, 12, 8, 102, 9, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 3, 10, 109, 8, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 3, 13, 119, 8, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 142, 8, 17, 1, 17, 0, 0,
		18, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		0, 3, 1, 0, 39, 40, 1, 0, 8, 27, 1, 0, 29, 30, 137, 0, 36, 1, 0, 0, 0,
		2, 42, 1, 0, 0, 0, 4, 44, 1, 0, 0, 0, 6, 54, 1, 0, 0, 0, 8, 66, 1, 0, 0,
		0, 10, 71, 1, 0, 0, 0, 12, 76, 1, 0, 0, 0, 14, 93, 1, 0, 0, 0, 16, 95,
		1, 0, 0, 0, 18, 103, 1, 0, 0, 0, 20, 108, 1, 0, 0, 0, 22, 112, 1, 0, 0,
		0, 24, 114, 1, 0, 0, 0, 26, 118, 1, 0, 0, 0, 28, 120, 1, 0, 0, 0, 30, 123,
		1, 0, 0, 0, 32, 127, 1, 0, 0, 0, 34, 141, 1, 0, 0, 0, 36, 38, 3, 2, 1,
		0, 37, 39, 5, 31, 0, 0, 38, 37, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 40,
		1, 0, 0, 0, 40, 41, 5, 0, 0, 1, 41, 1, 1, 0, 0, 0, 42, 43, 3, 4, 2, 0,
		43, 3, 1, 0, 0, 0, 44, 45, 5, 1, 0, 0, 45, 47, 5, 2, 0, 0, 46, 48, 3, 30,
		15, 0, 47, 46, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49,
		50, 3, 20, 10, 0, 50, 51, 5, 35, 0, 0, 51, 52, 3, 6, 3, 0, 52, 53, 5, 36,
		0, 0, 53, 5, 1, 0, 0, 0, 54, 59, 3, 8, 4, 0, 55, 56, 5, 34, 0, 0, 56, 58,
		3, 8, 4, 0, 57, 55, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0,
		59, 60, 1, 0, 0, 0, 60, 64, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 63, 5,
		34, 0, 0, 63, 65, 3, 10, 5, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0,
		65, 7, 1, 0, 0, 0, 66, 67, 3, 18, 9, 0, 67, 69, 3, 26, 13, 0, 68, 70, 3,
		28, 14, 0, 69, 68, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 9, 1, 0, 0, 0, 71,
		72, 3, 28, 14, 0, 72, 73, 5, 35, 0, 0, 73, 74, 3, 12, 6, 0, 74, 75, 5,
		36, 0, 0, 75, 11, 1, 0, 0, 0, 76, 79, 3, 14, 7, 0, 77, 78, 5, 34, 0, 0,
		78, 80, 3, 16, 8, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 13, 1,
		0, 0, 0, 81, 94, 3, 18, 9, 0, 82, 83, 5, 35, 0, 0, 83, 88, 3, 18, 9, 0,
		84, 85, 5, 34, 0, 0, 85, 87, 3, 18, 9, 0, 86, 84, 1, 0, 0, 0, 87, 90, 1,
		0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 91, 1, 0, 0, 0, 90,
		88, 1, 0, 0, 0, 91, 92, 5, 36, 0, 0, 92, 94, 1, 0, 0, 0, 93, 81, 1, 0,
		0, 0, 93, 82, 1, 0, 0, 0, 94, 15, 1, 0, 0, 0, 95, 100, 3, 18, 9, 0, 96,
		97, 5, 34, 0, 0, 97, 99, 3, 18, 9, 0, 98, 96, 1, 0, 0, 0, 99, 102, 1, 0,
		0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 17, 1, 0, 0, 0, 102,
		100, 1, 0, 0, 0, 103, 104, 5, 39, 0, 0, 104, 19, 1, 0, 0, 0, 105, 106,
		3, 22, 11, 0, 106, 107, 5, 33, 0, 0, 107, 109, 1, 0, 0, 0, 108, 105, 1,
		0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 5, 39, 0,
		0, 111, 21, 1, 0, 0, 0, 112, 113, 3, 24, 12, 0, 113, 23, 1, 0, 0, 0, 114,
		115, 7, 0, 0, 0, 115, 25, 1, 0, 0, 0, 116, 119, 3, 32, 16, 0, 117, 119,
		3, 34, 17, 0, 118, 116, 1, 0, 0, 0, 118, 117, 1, 0, 0, 0, 119, 27, 1, 0,
		0, 0, 120, 121, 5, 6, 0, 0, 121, 122, 5, 7, 0, 0, 122, 29, 1, 0, 0, 0,
		123, 124, 5, 3, 0, 0, 124, 125, 5, 4, 0, 0, 125, 126, 5, 5, 0, 0, 126,
		31, 1, 0, 0, 0, 127, 128, 7, 1, 0, 0, 128, 33, 1, 0, 0, 0, 129, 130, 5,
		28, 0, 0, 130, 131, 5, 37, 0, 0, 131, 132, 3, 32, 16, 0, 132, 133, 5, 34,
		0, 0, 133, 134, 3, 32, 16, 0, 134, 135, 5, 38, 0, 0, 135, 142, 1, 0, 0,
		0, 136, 137, 7, 2, 0, 0, 137, 138, 5, 37, 0, 0, 138, 139, 3, 32, 16, 0,
		139, 140, 5, 38, 0, 0, 140, 142, 1, 0, 0, 0, 141, 129, 1, 0, 0, 0, 141,
		136, 1, 0, 0, 0, 142, 35, 1, 0, 0, 0, 12, 38, 47, 59, 64, 69, 79, 88, 93,
		100, 108, 118, 141,
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
	CqlParserK_ASCII                = 8
	CqlParserK_BIGINT               = 9
	CqlParserK_BLOB                 = 10
	CqlParserK_BOOLEAN              = 11
	CqlParserK_COUNTER              = 12
	CqlParserK_DATE                 = 13
	CqlParserK_DECIMAL              = 14
	CqlParserK_DOUBLE               = 15
	CqlParserK_FLOAT                = 16
	CqlParserK_INET                 = 17
	CqlParserK_INT                  = 18
	CqlParserK_SMALLINT             = 19
	CqlParserK_TEXT                 = 20
	CqlParserK_TIME                 = 21
	CqlParserK_TIMESTAMP            = 22
	CqlParserK_TIMEUUID             = 23
	CqlParserK_TINYINT              = 24
	CqlParserK_UUID                 = 25
	CqlParserK_VARCHAR              = 26
	CqlParserK_VARINT               = 27
	CqlParserK_MAP                  = 28
	CqlParserK_SET                  = 29
	CqlParserK_LIST                 = 30
	CqlParserSEMICOLON              = 31
	CqlParserDQUOTE                 = 32
	CqlParserDOT                    = 33
	CqlParserCOMMA                  = 34
	CqlParserL_PAREN                = 35
	CqlParserR_PAREN                = 36
	CqlParserL_ANGLE_BRACKET        = 37
	CqlParserR_ANGLE_BRACKET        = 38
	CqlParserIDENTIFIER             = 39
	CqlParserIDENTIFIER_WITH_HYPHEN = 40
	CqlParserWS                     = 41
	CqlParserCOMMENT                = 42
	CqlParserMULTILINE_COMMENT      = 43
	CqlParserUNKNOWN                = 44
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
		p.SetState(36)
		p.CqlStatement()
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserSEMICOLON {
		{
			p.SetState(37)
			p.Match(CqlParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(40)
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
		p.SetState(42)
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
		p.SetState(44)
		p.Match(CqlParserK_CREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(CqlParserK_TABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_IF {
		{
			p.SetState(46)
			p.IfNotExist()
		}

	}
	{
		p.SetState(49)
		p.TableName()
	}
	{
		p.SetState(50)
		p.Match(CqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.ColumnDefinitionList()
	}
	{
		p.SetState(52)
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
		p.SetState(54)
		p.ColumnDefinition()
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(55)
				p.Match(CqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(56)
				p.ColumnDefinition()
			}

		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserCOMMA {
		{
			p.SetState(62)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
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
		p.SetState(66)
		p.ColumnName()
	}
	{
		p.SetState(67)
		p.CqlType()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_PRIMARY {
		{
			p.SetState(68)
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
		p.SetState(71)
		p.PrimaryKeyKeywords()
	}
	{
		p.SetState(72)
		p.Match(CqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.PrimaryKey()
	}
	{
		p.SetState(74)
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
		p.SetState(76)
		p.PartitionKey()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserCOMMA {
		{
			p.SetState(77)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
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

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(81)
			p.ColumnName()
		}

	case CqlParserL_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.Match(CqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.ColumnName()
		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CqlParserCOMMA {
			{
				p.SetState(84)
				p.Match(CqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(85)
				p.ColumnName()
			}

			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(91)
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
		p.SetState(95)
		p.ColumnName()
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CqlParserCOMMA {
		{
			p.SetState(96)
			p.Match(CqlParserCOMMA)
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
		p.SetState(103)
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
	p.SetState(108)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(105)
			p.KeyspaceName()
		}
		{
			p.SetState(106)
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
		p.SetState(110)
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
		p.SetState(112)
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
		p.SetState(114)
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
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_ASCII, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_COUNTER, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DOUBLE, CqlParserK_FLOAT, CqlParserK_INET, CqlParserK_INT, CqlParserK_SMALLINT, CqlParserK_TEXT, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_UUID, CqlParserK_VARCHAR, CqlParserK_VARINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.CqlNativeType()
		}

	case CqlParserK_MAP, CqlParserK_SET, CqlParserK_LIST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
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
		p.SetState(120)
		p.Match(CqlParserK_PRIMARY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
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
		p.SetState(123)
		p.Match(CqlParserK_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(CqlParserK_NOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
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
		p.SetState(127)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268435200) != 0) {
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

	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_MAP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.Match(CqlParserK_MAP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Match(CqlParserL_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.CqlNativeType()
		}
		{
			p.SetState(132)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.CqlNativeType()
		}
		{
			p.SetState(134)
			p.Match(CqlParserR_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserK_SET, CqlParserK_LIST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CqlParserK_SET || _la == CqlParserK_LIST) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(137)
			p.Match(CqlParserL_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.CqlNativeType()
		}
		{
			p.SetState(139)
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
