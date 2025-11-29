// Code generated from parser/ExprParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ExprParser
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


type ExprParser struct {
	*antlr.BaseParser
}

var ExprParserParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func exprparserParserInit() {
  staticData := &ExprParserParserStaticData
  staticData.LiteralNames = []string{
    "", "", "", "", "", "", "", "", "", "", "'nothing'", "'fraud'", "'cancelled'", 
    "'invalid'", "'alert'", "'conflict'", "'ungraded'", "'obscured'", "'absent'", 
    "'present'", "'excused'", "'toolow'", "", "'=>'", "'?'", "':'", "';'", 
    "','", "'='", "'['", "']'", "'('", "')'", "", "", "'!~'", "'~'", "'+'", 
    "'-'", "'*'", "'/'", "'>'", "'>='", "'!='", "'=='", "'<='", "'<'", "'&&'", 
    "'||'", "'!'",
  }
  staticData.SymbolicNames = []string{
    "", "APPLY", "TO", "WHEN", "EXISTS", "EXPLAIN", "MAX", "MIN", "ROUND", 
    "IN", "CONST_NOTHING", "CONST_FRAUD", "CONST_CANCELLED", "CONST_INVALID", 
    "CONST_ALERT", "COSNT_CONFLICT", "CONST_UNGRADED", "COSNT_OBSCURED", 
    "CONST_ABSENT", "CONST_PRESENT", "CONST_EXCUSED", "CONST_TOOLOW", "IDENTIFIER", 
    "LOP", "ASK", "COLON", "SEMICOLON", "COMMA", "ASSIGN", "LBRACK", "RBRACK", 
    "LPAREN", "RPAREN", "STRING", "REGEX", "NM", "MA", "ADD", "SUB", "MUL", 
    "DIV", "GT", "GE", "NE", "EQ", "LE", "LT", "AND", "OR", "NOT", "FLOAT", 
    "INT", "SPACES",
  }
  staticData.RuleNames = []string{
    "algorithm", "apply", "normal_operation", "when_operation", "function", 
    "constant", "constantList", "lambda", "list", "expression", "assignmentExpression", 
    "iterateExpression", "conditionalExpression", "logicalOrExpression", 
    "logicalAndExpression", "equalityExpression", "relationalExpression", 
    "regexExpression", "additiveExpression", "multiplicativeExpression", 
    "unaryExpression", "functionCall",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 52, 212, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2, 
	21, 7, 21, 1, 0, 1, 0, 1, 0, 5, 0, 48, 8, 0, 10, 0, 12, 0, 51, 9, 0, 1, 
	0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 2, 62, 8, 2, 1, 2, 
	1, 2, 1, 2, 3, 2, 67, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 
	4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 83, 8, 6, 10, 6, 12, 6, 86, 
	9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 95, 8, 8, 10, 8, 
	12, 8, 98, 9, 8, 1, 9, 1, 9, 1, 9, 3, 9, 103, 8, 9, 1, 10, 1, 10, 1, 10, 
	1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 
	11, 3, 11, 119, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 
	127, 8, 12, 1, 13, 1, 13, 1, 13, 5, 13, 132, 8, 13, 10, 13, 12, 13, 135, 
	9, 13, 1, 14, 1, 14, 1, 14, 5, 14, 140, 8, 14, 10, 14, 12, 14, 143, 9, 
	14, 1, 15, 1, 15, 1, 15, 5, 15, 148, 8, 15, 10, 15, 12, 15, 151, 9, 15, 
	1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 157, 8, 16, 10, 16, 12, 16, 160, 9, 
	16, 3, 16, 162, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 
	5, 18, 171, 8, 18, 10, 18, 12, 18, 174, 9, 18, 1, 19, 1, 19, 1, 19, 5, 
	19, 179, 8, 19, 10, 19, 12, 19, 182, 9, 19, 1, 20, 3, 20, 185, 8, 20, 1, 
	20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 
	3, 20, 198, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 205, 8, 21, 
	10, 21, 12, 21, 208, 9, 21, 1, 21, 1, 21, 1, 21, 0, 0, 22, 0, 2, 4, 6, 
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 
	0, 7, 1, 0, 6, 8, 1, 0, 10, 21, 1, 0, 43, 44, 3, 0, 35, 36, 41, 42, 45, 
	46, 1, 0, 37, 38, 1, 0, 39, 40, 2, 0, 38, 38, 49, 49, 216, 0, 49, 1, 0, 
	0, 0, 2, 54, 1, 0, 0, 0, 4, 61, 1, 0, 0, 0, 6, 70, 1, 0, 0, 0, 8, 75, 1, 
	0, 0, 0, 10, 77, 1, 0, 0, 0, 12, 79, 1, 0, 0, 0, 14, 87, 1, 0, 0, 0, 16, 
	91, 1, 0, 0, 0, 18, 102, 1, 0, 0, 0, 20, 104, 1, 0, 0, 0, 22, 108, 1, 0, 
	0, 0, 24, 120, 1, 0, 0, 0, 26, 128, 1, 0, 0, 0, 28, 136, 1, 0, 0, 0, 30, 
	144, 1, 0, 0, 0, 32, 161, 1, 0, 0, 0, 34, 163, 1, 0, 0, 0, 36, 167, 1, 
	0, 0, 0, 38, 175, 1, 0, 0, 0, 40, 184, 1, 0, 0, 0, 42, 199, 1, 0, 0, 0, 
	44, 48, 3, 2, 1, 0, 45, 48, 3, 4, 2, 0, 46, 48, 3, 6, 3, 0, 47, 44, 1, 
	0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 46, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 
	47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0, 51, 49, 1, 0, 0, 
	0, 52, 53, 5, 0, 0, 1, 53, 1, 1, 0, 0, 0, 54, 55, 5, 1, 0, 0, 55, 56, 3, 
	14, 7, 0, 56, 57, 5, 2, 0, 0, 57, 58, 3, 16, 8, 0, 58, 59, 5, 26, 0, 0, 
	59, 3, 1, 0, 0, 0, 60, 62, 5, 5, 0, 0, 61, 60, 1, 0, 0, 0, 61, 62, 1, 0, 
	0, 0, 62, 63, 1, 0, 0, 0, 63, 66, 3, 18, 9, 0, 64, 65, 5, 27, 0, 0, 65, 
	67, 5, 33, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 1, 0, 
	0, 0, 68, 69, 5, 26, 0, 0, 69, 5, 1, 0, 0, 0, 70, 71, 5, 3, 0, 0, 71, 72, 
	3, 16, 8, 0, 72, 73, 5, 4, 0, 0, 73, 74, 3, 4, 2, 0, 74, 7, 1, 0, 0, 0, 
	75, 76, 7, 0, 0, 0, 76, 9, 1, 0, 0, 0, 77, 78, 7, 1, 0, 0, 78, 11, 1, 0, 
	0, 0, 79, 84, 3, 10, 5, 0, 80, 81, 5, 27, 0, 0, 81, 83, 3, 10, 5, 0, 82, 
	80, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 
	0, 85, 13, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 88, 5, 22, 0, 0, 88, 89, 
	5, 23, 0, 0, 89, 90, 3, 24, 12, 0, 90, 15, 1, 0, 0, 0, 91, 96, 5, 22, 0, 
	0, 92, 93, 5, 27, 0, 0, 93, 95, 5, 22, 0, 0, 94, 92, 1, 0, 0, 0, 95, 98, 
	1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 17, 1, 0, 0, 0, 
	98, 96, 1, 0, 0, 0, 99, 103, 3, 22, 11, 0, 100, 103, 3, 24, 12, 0, 101, 
	103, 3, 20, 10, 0, 102, 99, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 101, 
	1, 0, 0, 0, 103, 19, 1, 0, 0, 0, 104, 105, 5, 22, 0, 0, 105, 106, 5, 28, 
	0, 0, 106, 107, 3, 18, 9, 0, 107, 21, 1, 0, 0, 0, 108, 109, 5, 22, 0, 0, 
	109, 110, 5, 9, 0, 0, 110, 111, 5, 29, 0, 0, 111, 112, 3, 12, 6, 0, 112, 
	118, 5, 30, 0, 0, 113, 114, 5, 24, 0, 0, 114, 115, 3, 18, 9, 0, 115, 116, 
	5, 25, 0, 0, 116, 117, 3, 18, 9, 0, 117, 119, 1, 0, 0, 0, 118, 113, 1, 
	0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 23, 1, 0, 0, 0, 120, 126, 3, 26, 13, 
	0, 121, 122, 5, 24, 0, 0, 122, 123, 3, 18, 9, 0, 123, 124, 5, 25, 0, 0, 
	124, 125, 3, 18, 9, 0, 125, 127, 1, 0, 0, 0, 126, 121, 1, 0, 0, 0, 126, 
	127, 1, 0, 0, 0, 127, 25, 1, 0, 0, 0, 128, 133, 3, 28, 14, 0, 129, 130, 
	5, 48, 0, 0, 130, 132, 3, 28, 14, 0, 131, 129, 1, 0, 0, 0, 132, 135, 1, 
	0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 27, 1, 0, 0, 
	0, 135, 133, 1, 0, 0, 0, 136, 141, 3, 30, 15, 0, 137, 138, 5, 47, 0, 0, 
	138, 140, 3, 30, 15, 0, 139, 137, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 
	139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 29, 1, 0, 0, 0, 143, 141, 1, 
	0, 0, 0, 144, 149, 3, 32, 16, 0, 145, 146, 7, 2, 0, 0, 146, 148, 3, 32, 
	16, 0, 147, 145, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 
	149, 150, 1, 0, 0, 0, 150, 31, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 162, 
	3, 34, 17, 0, 153, 158, 3, 36, 18, 0, 154, 155, 7, 3, 0, 0, 155, 157, 3, 
	36, 18, 0, 156, 154, 1, 0, 0, 0, 157, 160, 1, 0, 0, 0, 158, 156, 1, 0, 
	0, 0, 158, 159, 1, 0, 0, 0, 159, 162, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 
	161, 152, 1, 0, 0, 0, 161, 153, 1, 0, 0, 0, 162, 33, 1, 0, 0, 0, 163, 164, 
	5, 22, 0, 0, 164, 165, 5, 36, 0, 0, 165, 166, 5, 33, 0, 0, 166, 35, 1, 
	0, 0, 0, 167, 172, 3, 38, 19, 0, 168, 169, 7, 4, 0, 0, 169, 171, 3, 38, 
	19, 0, 170, 168, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 
	172, 173, 1, 0, 0, 0, 173, 37, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 175, 180, 
	3, 40, 20, 0, 176, 177, 7, 5, 0, 0, 177, 179, 3, 40, 20, 0, 178, 176, 1, 
	0, 0, 0, 179, 182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 
	0, 181, 39, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 185, 7, 6, 0, 0, 184, 
	183, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 197, 1, 0, 0, 0, 186, 198, 
	3, 10, 5, 0, 187, 198, 5, 22, 0, 0, 188, 198, 5, 51, 0, 0, 189, 198, 5, 
	50, 0, 0, 190, 198, 5, 34, 0, 0, 191, 198, 5, 33, 0, 0, 192, 193, 5, 31, 
	0, 0, 193, 194, 3, 18, 9, 0, 194, 195, 5, 32, 0, 0, 195, 198, 1, 0, 0, 
	0, 196, 198, 3, 42, 21, 0, 197, 186, 1, 0, 0, 0, 197, 187, 1, 0, 0, 0, 
	197, 188, 1, 0, 0, 0, 197, 189, 1, 0, 0, 0, 197, 190, 1, 0, 0, 0, 197, 
	191, 1, 0, 0, 0, 197, 192, 1, 0, 0, 0, 197, 196, 1, 0, 0, 0, 198, 41, 1, 
	0, 0, 0, 199, 200, 3, 8, 4, 0, 200, 201, 5, 31, 0, 0, 201, 206, 3, 18, 
	9, 0, 202, 203, 5, 27, 0, 0, 203, 205, 3, 18, 9, 0, 204, 202, 1, 0, 0, 
	0, 205, 208, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 
	209, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 209, 210, 5, 32, 0, 0, 210, 43, 
	1, 0, 0, 0, 19, 47, 49, 61, 66, 84, 96, 102, 118, 126, 133, 141, 149, 158, 
	161, 172, 180, 184, 197, 206,
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

// ExprParserInit initializes any static state used to implement ExprParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewExprParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExprParserInit() {
  staticData := &ExprParserParserStaticData
  staticData.once.Do(exprparserParserInit)
}

// NewExprParser produces a new parser instance for the optional input antlr.TokenStream.
func NewExprParser(input antlr.TokenStream) *ExprParser {
	ExprParserInit()
	this := new(ExprParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &ExprParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ExprParser.g4"

	return this
}


// ExprParser tokens.
const (
	ExprParserEOF = antlr.TokenEOF
	ExprParserAPPLY = 1
	ExprParserTO = 2
	ExprParserWHEN = 3
	ExprParserEXISTS = 4
	ExprParserEXPLAIN = 5
	ExprParserMAX = 6
	ExprParserMIN = 7
	ExprParserROUND = 8
	ExprParserIN = 9
	ExprParserCONST_NOTHING = 10
	ExprParserCONST_FRAUD = 11
	ExprParserCONST_CANCELLED = 12
	ExprParserCONST_INVALID = 13
	ExprParserCONST_ALERT = 14
	ExprParserCOSNT_CONFLICT = 15
	ExprParserCONST_UNGRADED = 16
	ExprParserCOSNT_OBSCURED = 17
	ExprParserCONST_ABSENT = 18
	ExprParserCONST_PRESENT = 19
	ExprParserCONST_EXCUSED = 20
	ExprParserCONST_TOOLOW = 21
	ExprParserIDENTIFIER = 22
	ExprParserLOP = 23
	ExprParserASK = 24
	ExprParserCOLON = 25
	ExprParserSEMICOLON = 26
	ExprParserCOMMA = 27
	ExprParserASSIGN = 28
	ExprParserLBRACK = 29
	ExprParserRBRACK = 30
	ExprParserLPAREN = 31
	ExprParserRPAREN = 32
	ExprParserSTRING = 33
	ExprParserREGEX = 34
	ExprParserNM = 35
	ExprParserMA = 36
	ExprParserADD = 37
	ExprParserSUB = 38
	ExprParserMUL = 39
	ExprParserDIV = 40
	ExprParserGT = 41
	ExprParserGE = 42
	ExprParserNE = 43
	ExprParserEQ = 44
	ExprParserLE = 45
	ExprParserLT = 46
	ExprParserAND = 47
	ExprParserOR = 48
	ExprParserNOT = 49
	ExprParserFLOAT = 50
	ExprParserINT = 51
	ExprParserSPACES = 52
)

// ExprParser rules.
const (
	ExprParserRULE_algorithm = 0
	ExprParserRULE_apply = 1
	ExprParserRULE_normal_operation = 2
	ExprParserRULE_when_operation = 3
	ExprParserRULE_function = 4
	ExprParserRULE_constant = 5
	ExprParserRULE_constantList = 6
	ExprParserRULE_lambda = 7
	ExprParserRULE_list = 8
	ExprParserRULE_expression = 9
	ExprParserRULE_assignmentExpression = 10
	ExprParserRULE_iterateExpression = 11
	ExprParserRULE_conditionalExpression = 12
	ExprParserRULE_logicalOrExpression = 13
	ExprParserRULE_logicalAndExpression = 14
	ExprParserRULE_equalityExpression = 15
	ExprParserRULE_relationalExpression = 16
	ExprParserRULE_regexExpression = 17
	ExprParserRULE_additiveExpression = 18
	ExprParserRULE_multiplicativeExpression = 19
	ExprParserRULE_unaryExpression = 20
	ExprParserRULE_functionCall = 21
)

// IAlgorithmContext is an interface to support dynamic dispatch.
type IAlgorithmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllApply() []IApplyContext
	Apply(i int) IApplyContext
	AllNormal_operation() []INormal_operationContext
	Normal_operation(i int) INormal_operationContext
	AllWhen_operation() []IWhen_operationContext
	When_operation(i int) IWhen_operationContext

	// IsAlgorithmContext differentiates from other interfaces.
	IsAlgorithmContext()
}

type AlgorithmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlgorithmContext() *AlgorithmContext {
	var p = new(AlgorithmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_algorithm
	return p
}

func InitEmptyAlgorithmContext(p *AlgorithmContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_algorithm
}

func (*AlgorithmContext) IsAlgorithmContext() {}

func NewAlgorithmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlgorithmContext {
	var p = new(AlgorithmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_algorithm

	return p
}

func (s *AlgorithmContext) GetParser() antlr.Parser { return s.parser }

func (s *AlgorithmContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExprParserEOF, 0)
}

func (s *AlgorithmContext) AllApply() []IApplyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IApplyContext); ok {
			len++
		}
	}

	tst := make([]IApplyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IApplyContext); ok {
			tst[i] = t.(IApplyContext)
			i++
		}
	}

	return tst
}

func (s *AlgorithmContext) Apply(i int) IApplyContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApplyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApplyContext)
}

func (s *AlgorithmContext) AllNormal_operation() []INormal_operationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INormal_operationContext); ok {
			len++
		}
	}

	tst := make([]INormal_operationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INormal_operationContext); ok {
			tst[i] = t.(INormal_operationContext)
			i++
		}
	}

	return tst
}

func (s *AlgorithmContext) Normal_operation(i int) INormal_operationContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INormal_operationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INormal_operationContext)
}

func (s *AlgorithmContext) AllWhen_operation() []IWhen_operationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWhen_operationContext); ok {
			len++
		}
	}

	tst := make([]IWhen_operationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWhen_operationContext); ok {
			tst[i] = t.(IWhen_operationContext)
			i++
		}
	}

	return tst
}

func (s *AlgorithmContext) When_operation(i int) IWhen_operationContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhen_operationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhen_operationContext)
}

func (s *AlgorithmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlgorithmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AlgorithmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterAlgorithm(s)
	}
}

func (s *AlgorithmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitAlgorithm(s)
	}
}

func (s *AlgorithmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitAlgorithm(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) Algorithm() (localctx IAlgorithmContext) {
	localctx = NewAlgorithmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExprParserRULE_algorithm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 3940952477531626) != 0) {
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ExprParserAPPLY:
			{
				p.SetState(44)
				p.Apply()
			}


		case ExprParserEXPLAIN, ExprParserMAX, ExprParserMIN, ExprParserROUND, ExprParserCONST_NOTHING, ExprParserCONST_FRAUD, ExprParserCONST_CANCELLED, ExprParserCONST_INVALID, ExprParserCONST_ALERT, ExprParserCOSNT_CONFLICT, ExprParserCONST_UNGRADED, ExprParserCOSNT_OBSCURED, ExprParserCONST_ABSENT, ExprParserCONST_PRESENT, ExprParserCONST_EXCUSED, ExprParserCONST_TOOLOW, ExprParserIDENTIFIER, ExprParserLPAREN, ExprParserSTRING, ExprParserREGEX, ExprParserSUB, ExprParserNOT, ExprParserFLOAT, ExprParserINT:
			{
				p.SetState(45)
				p.Normal_operation()
			}


		case ExprParserWHEN:
			{
				p.SetState(46)
				p.When_operation()
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(52)
		p.Match(ExprParserEOF)
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


// IApplyContext is an interface to support dynamic dispatch.
type IApplyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	APPLY() antlr.TerminalNode
	Lambda() ILambdaContext
	TO() antlr.TerminalNode
	List() IListContext
	SEMICOLON() antlr.TerminalNode

	// IsApplyContext differentiates from other interfaces.
	IsApplyContext()
}

type ApplyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApplyContext() *ApplyContext {
	var p = new(ApplyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_apply
	return p
}

func InitEmptyApplyContext(p *ApplyContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_apply
}

func (*ApplyContext) IsApplyContext() {}

func NewApplyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApplyContext {
	var p = new(ApplyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_apply

	return p
}

func (s *ApplyContext) GetParser() antlr.Parser { return s.parser }

func (s *ApplyContext) APPLY() antlr.TerminalNode {
	return s.GetToken(ExprParserAPPLY, 0)
}

func (s *ApplyContext) Lambda() ILambdaContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambdaContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambdaContext)
}

func (s *ApplyContext) TO() antlr.TerminalNode {
	return s.GetToken(ExprParserTO, 0)
}

func (s *ApplyContext) List() IListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ApplyContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ExprParserSEMICOLON, 0)
}

func (s *ApplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ApplyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterApply(s)
	}
}

func (s *ApplyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitApply(s)
	}
}

func (s *ApplyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitApply(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) Apply() (localctx IApplyContext) {
	localctx = NewApplyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExprParserRULE_apply)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(ExprParserAPPLY)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Lambda()
	}
	{
		p.SetState(56)
		p.Match(ExprParserTO)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(57)
		p.List()
	}
	{
		p.SetState(58)
		p.Match(ExprParserSEMICOLON)
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


// INormal_operationContext is an interface to support dynamic dispatch.
type INormal_operationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	SEMICOLON() antlr.TerminalNode
	EXPLAIN() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsNormal_operationContext differentiates from other interfaces.
	IsNormal_operationContext()
}

type Normal_operationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNormal_operationContext() *Normal_operationContext {
	var p = new(Normal_operationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_normal_operation
	return p
}

func InitEmptyNormal_operationContext(p *Normal_operationContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_normal_operation
}

func (*Normal_operationContext) IsNormal_operationContext() {}

func NewNormal_operationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Normal_operationContext {
	var p = new(Normal_operationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_normal_operation

	return p
}

func (s *Normal_operationContext) GetParser() antlr.Parser { return s.parser }

func (s *Normal_operationContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Normal_operationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ExprParserSEMICOLON, 0)
}

func (s *Normal_operationContext) EXPLAIN() antlr.TerminalNode {
	return s.GetToken(ExprParserEXPLAIN, 0)
}

func (s *Normal_operationContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ExprParserCOMMA, 0)
}

func (s *Normal_operationContext) STRING() antlr.TerminalNode {
	return s.GetToken(ExprParserSTRING, 0)
}

func (s *Normal_operationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Normal_operationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Normal_operationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterNormal_operation(s)
	}
}

func (s *Normal_operationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitNormal_operation(s)
	}
}

func (s *Normal_operationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitNormal_operation(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) Normal_operation() (localctx INormal_operationContext) {
	localctx = NewNormal_operationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ExprParserRULE_normal_operation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == ExprParserEXPLAIN {
		{
			p.SetState(60)
			p.Match(ExprParserEXPLAIN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(63)
		p.Expression()
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == ExprParserCOMMA {
		{
			p.SetState(64)
			p.Match(ExprParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Match(ExprParserSTRING)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(68)
		p.Match(ExprParserSEMICOLON)
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


// IWhen_operationContext is an interface to support dynamic dispatch.
type IWhen_operationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHEN() antlr.TerminalNode
	List() IListContext
	EXISTS() antlr.TerminalNode
	Normal_operation() INormal_operationContext

	// IsWhen_operationContext differentiates from other interfaces.
	IsWhen_operationContext()
}

type When_operationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhen_operationContext() *When_operationContext {
	var p = new(When_operationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_when_operation
	return p
}

func InitEmptyWhen_operationContext(p *When_operationContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_when_operation
}

func (*When_operationContext) IsWhen_operationContext() {}

func NewWhen_operationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *When_operationContext {
	var p = new(When_operationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_when_operation

	return p
}

func (s *When_operationContext) GetParser() antlr.Parser { return s.parser }

func (s *When_operationContext) WHEN() antlr.TerminalNode {
	return s.GetToken(ExprParserWHEN, 0)
}

func (s *When_operationContext) List() IListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *When_operationContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(ExprParserEXISTS, 0)
}

func (s *When_operationContext) Normal_operation() INormal_operationContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INormal_operationContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INormal_operationContext)
}

func (s *When_operationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *When_operationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *When_operationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterWhen_operation(s)
	}
}

func (s *When_operationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitWhen_operation(s)
	}
}

func (s *When_operationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitWhen_operation(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) When_operation() (localctx IWhen_operationContext) {
	localctx = NewWhen_operationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ExprParserRULE_when_operation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(ExprParserWHEN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(71)
		p.List()
	}
	{
		p.SetState(72)
		p.Match(ExprParserEXISTS)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Normal_operation()
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


// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAX() antlr.TerminalNode
	MIN() antlr.TerminalNode
	ROUND() antlr.TerminalNode

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) MAX() antlr.TerminalNode {
	return s.GetToken(ExprParserMAX, 0)
}

func (s *FunctionContext) MIN() antlr.TerminalNode {
	return s.GetToken(ExprParserMIN, 0)
}

func (s *FunctionContext) ROUND() antlr.TerminalNode {
	return s.GetToken(ExprParserROUND, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ExprParserRULE_function)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 448) != 0)) {
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


// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST_ABSENT() antlr.TerminalNode
	CONST_PRESENT() antlr.TerminalNode
	CONST_EXCUSED() antlr.TerminalNode
	CONST_NOTHING() antlr.TerminalNode
	CONST_FRAUD() antlr.TerminalNode
	CONST_CANCELLED() antlr.TerminalNode
	CONST_INVALID() antlr.TerminalNode
	CONST_ALERT() antlr.TerminalNode
	COSNT_CONFLICT() antlr.TerminalNode
	CONST_UNGRADED() antlr.TerminalNode
	COSNT_OBSCURED() antlr.TerminalNode
	CONST_TOOLOW() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) CONST_ABSENT() antlr.TerminalNode {
	return s.GetToken(ExprParserCONST_ABSENT, 0)
}

func (s *ConstantContext) CONST_PRESENT() antlr.TerminalNode {
	return s.GetToken(ExprParserCONST_PRESENT, 0)
}

func (s *ConstantContext) CONST_EXCUSED() antlr.TerminalNode {
	return s.GetToken(ExprParserCONST_EXCUSED, 0)
}

func (s *ConstantContext) CONST_NOTHING() antlr.TerminalNode {
	return s.GetToken(ExprParserCONST_NOTHING, 0)
}

func (s *ConstantContext) CONST_FRAUD() antlr.TerminalNode {
	return s.GetToken(ExprParserCONST_FRAUD, 0)
}

func (s *ConstantContext) CONST_CANCELLED() antlr.TerminalNode {
	return s.GetToken(ExprParserCONST_CANCELLED, 0)
}

func (s *ConstantContext) CONST_INVALID() antlr.TerminalNode {
	return s.GetToken(ExprParserCONST_INVALID, 0)
}

func (s *ConstantContext) CONST_ALERT() antlr.TerminalNode {
	return s.GetToken(ExprParserCONST_ALERT, 0)
}

func (s *ConstantContext) COSNT_CONFLICT() antlr.TerminalNode {
	return s.GetToken(ExprParserCOSNT_CONFLICT, 0)
}

func (s *ConstantContext) CONST_UNGRADED() antlr.TerminalNode {
	return s.GetToken(ExprParserCONST_UNGRADED, 0)
}

func (s *ConstantContext) COSNT_OBSCURED() antlr.TerminalNode {
	return s.GetToken(ExprParserCOSNT_OBSCURED, 0)
}

func (s *ConstantContext) CONST_TOOLOW() antlr.TerminalNode {
	return s.GetToken(ExprParserCONST_TOOLOW, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ExprParserRULE_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 4193280) != 0)) {
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


// IConstantListContext is an interface to support dynamic dispatch.
type IConstantListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConstant() []IConstantContext
	Constant(i int) IConstantContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsConstantListContext differentiates from other interfaces.
	IsConstantListContext()
}

type ConstantListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantListContext() *ConstantListContext {
	var p = new(ConstantListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_constantList
	return p
}

func InitEmptyConstantListContext(p *ConstantListContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_constantList
}

func (*ConstantListContext) IsConstantListContext() {}

func NewConstantListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantListContext {
	var p = new(ConstantListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_constantList

	return p
}

func (s *ConstantListContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantListContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *ConstantListContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ConstantListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ExprParserCOMMA)
}

func (s *ConstantListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserCOMMA, i)
}

func (s *ConstantListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConstantListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterConstantList(s)
	}
}

func (s *ConstantListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitConstantList(s)
	}
}

func (s *ConstantListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitConstantList(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) ConstantList() (localctx IConstantListContext) {
	localctx = NewConstantListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ExprParserRULE_constantList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Constant()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == ExprParserCOMMA {
		{
			p.SetState(80)
			p.Match(ExprParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Constant()
		}


		p.SetState(86)
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


// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LOP() antlr.TerminalNode
	ConditionalExpression() IConditionalExpressionContext

	// IsLambdaContext differentiates from other interfaces.
	IsLambdaContext()
}

type LambdaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaContext() *LambdaContext {
	var p = new(LambdaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_lambda
	return p
}

func InitEmptyLambdaContext(p *LambdaContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_lambda
}

func (*LambdaContext) IsLambdaContext() {}

func NewLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaContext {
	var p = new(LambdaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_lambda

	return p
}

func (s *LambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, 0)
}

func (s *LambdaContext) LOP() antlr.TerminalNode {
	return s.GetToken(ExprParserLOP, 0)
}

func (s *LambdaContext) ConditionalExpression() IConditionalExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExpressionContext)
}

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterLambda(s)
	}
}

func (s *LambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitLambda(s)
	}
}

func (s *LambdaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitLambda(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) Lambda() (localctx ILambdaContext) {
	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ExprParserRULE_lambda)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Match(ExprParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(ExprParserLOP)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(89)
		p.ConditionalExpression()
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


// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ExprParserIDENTIFIER)
}

func (s *ListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, i)
}

func (s *ListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ExprParserCOMMA)
}

func (s *ListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserCOMMA, i)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ExprParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(ExprParserIDENTIFIER)
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


	for _la == ExprParserCOMMA {
		{
			p.SetState(92)
			p.Match(ExprParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Match(ExprParserIDENTIFIER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(98)
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


// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IterateExpression() IIterateExpressionContext
	ConditionalExpression() IConditionalExpressionContext
	AssignmentExpression() IAssignmentExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) IterateExpression() IIterateExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIterateExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIterateExpressionContext)
}

func (s *ExpressionContext) ConditionalExpression() IConditionalExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExpressionContext)
}

func (s *ExpressionContext) AssignmentExpression() IAssignmentExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ExprParserRULE_expression)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.IterateExpression()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.ConditionalExpression()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(101)
			p.AssignmentExpression()
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


// IAssignmentExpressionContext is an interface to support dynamic dispatch.
type IAssignmentExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignmentExpressionContext differentiates from other interfaces.
	IsAssignmentExpressionContext()
}

type AssignmentExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentExpressionContext() *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_assignmentExpression
	return p
}

func InitEmptyAssignmentExpressionContext(p *AssignmentExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_assignmentExpression
}

func (*AssignmentExpressionContext) IsAssignmentExpressionContext() {}

func NewAssignmentExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_assignmentExpression

	return p
}

func (s *AssignmentExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, 0)
}

func (s *AssignmentExpressionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ExprParserASSIGN, 0)
}

func (s *AssignmentExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AssignmentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitAssignmentExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) AssignmentExpression() (localctx IAssignmentExpressionContext) {
	localctx = NewAssignmentExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ExprParserRULE_assignmentExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(ExprParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(ExprParserASSIGN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Expression()
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


// IIterateExpressionContext is an interface to support dynamic dispatch.
type IIterateExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	IN() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	ConstantList() IConstantListContext
	RBRACK() antlr.TerminalNode
	ASK() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COLON() antlr.TerminalNode

	// IsIterateExpressionContext differentiates from other interfaces.
	IsIterateExpressionContext()
}

type IterateExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterateExpressionContext() *IterateExpressionContext {
	var p = new(IterateExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_iterateExpression
	return p
}

func InitEmptyIterateExpressionContext(p *IterateExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_iterateExpression
}

func (*IterateExpressionContext) IsIterateExpressionContext() {}

func NewIterateExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterateExpressionContext {
	var p = new(IterateExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_iterateExpression

	return p
}

func (s *IterateExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *IterateExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, 0)
}

func (s *IterateExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(ExprParserIN, 0)
}

func (s *IterateExpressionContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ExprParserLBRACK, 0)
}

func (s *IterateExpressionContext) ConstantList() IConstantListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantListContext)
}

func (s *IterateExpressionContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ExprParserRBRACK, 0)
}

func (s *IterateExpressionContext) ASK() antlr.TerminalNode {
	return s.GetToken(ExprParserASK, 0)
}

func (s *IterateExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IterateExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IterateExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(ExprParserCOLON, 0)
}

func (s *IterateExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterateExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IterateExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterIterateExpression(s)
	}
}

func (s *IterateExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitIterateExpression(s)
	}
}

func (s *IterateExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitIterateExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) IterateExpression() (localctx IIterateExpressionContext) {
	localctx = NewIterateExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ExprParserRULE_iterateExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(ExprParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(ExprParserIN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Match(ExprParserLBRACK)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(111)
		p.ConstantList()
	}
	{
		p.SetState(112)
		p.Match(ExprParserRBRACK)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == ExprParserASK {
		{
			p.SetState(113)
			p.Match(ExprParserASK)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Expression()
		}
		{
			p.SetState(115)
			p.Match(ExprParserCOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Expression()
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


// IConditionalExpressionContext is an interface to support dynamic dispatch.
type IConditionalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalOrExpression() ILogicalOrExpressionContext
	ASK() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COLON() antlr.TerminalNode

	// IsConditionalExpressionContext differentiates from other interfaces.
	IsConditionalExpressionContext()
}

type ConditionalExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalExpressionContext() *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_conditionalExpression
	return p
}

func InitEmptyConditionalExpressionContext(p *ConditionalExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_conditionalExpression
}

func (*ConditionalExpressionContext) IsConditionalExpressionContext() {}

func NewConditionalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_conditionalExpression

	return p
}

func (s *ConditionalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalExpressionContext) LogicalOrExpression() ILogicalOrExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOrExpressionContext)
}

func (s *ConditionalExpressionContext) ASK() antlr.TerminalNode {
	return s.GetToken(ExprParserASK, 0)
}

func (s *ConditionalExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(ExprParserCOLON, 0)
}

func (s *ConditionalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConditionalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterConditionalExpression(s)
	}
}

func (s *ConditionalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitConditionalExpression(s)
	}
}

func (s *ConditionalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitConditionalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) ConditionalExpression() (localctx IConditionalExpressionContext) {
	localctx = NewConditionalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ExprParserRULE_conditionalExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.LogicalOrExpression()
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == ExprParserASK {
		{
			p.SetState(121)
			p.Match(ExprParserASK)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Expression()
		}
		{
			p.SetState(123)
			p.Match(ExprParserCOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Expression()
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


// ILogicalOrExpressionContext is an interface to support dynamic dispatch.
type ILogicalOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLogicalAndExpression() []ILogicalAndExpressionContext
	LogicalAndExpression(i int) ILogicalAndExpressionContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsLogicalOrExpressionContext differentiates from other interfaces.
	IsLogicalOrExpressionContext()
}

type LogicalOrExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOrExpressionContext() *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_logicalOrExpression
	return p
}

func InitEmptyLogicalOrExpressionContext(p *LogicalOrExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_logicalOrExpression
}

func (*LogicalOrExpressionContext) IsLogicalOrExpressionContext() {}

func NewLogicalOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_logicalOrExpression

	return p
}

func (s *LogicalOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrExpressionContext) AllLogicalAndExpression() []ILogicalAndExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalAndExpressionContext); ok {
			len++
		}
	}

	tst := make([]ILogicalAndExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalAndExpressionContext); ok {
			tst[i] = t.(ILogicalAndExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) LogicalAndExpression(i int) ILogicalAndExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalAndExpressionContext)
}

func (s *LogicalOrExpressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(ExprParserOR)
}

func (s *LogicalOrExpressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserOR, i)
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitLogicalOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) LogicalOrExpression() (localctx ILogicalOrExpressionContext) {
	localctx = NewLogicalOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ExprParserRULE_logicalOrExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.LogicalAndExpression()
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == ExprParserOR {
		{
			p.SetState(129)
			p.Match(ExprParserOR)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(130)
			p.LogicalAndExpression()
		}


		p.SetState(135)
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


// ILogicalAndExpressionContext is an interface to support dynamic dispatch.
type ILogicalAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEqualityExpression() []IEqualityExpressionContext
	EqualityExpression(i int) IEqualityExpressionContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsLogicalAndExpressionContext differentiates from other interfaces.
	IsLogicalAndExpressionContext()
}

type LogicalAndExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalAndExpressionContext() *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_logicalAndExpression
	return p
}

func InitEmptyLogicalAndExpressionContext(p *LogicalAndExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_logicalAndExpression
}

func (*LogicalAndExpressionContext) IsLogicalAndExpressionContext() {}

func NewLogicalAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_logicalAndExpression

	return p
}

func (s *LogicalAndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndExpressionContext) AllEqualityExpression() []IEqualityExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			len++
		}
	}

	tst := make([]IEqualityExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualityExpressionContext); ok {
			tst[i] = t.(IEqualityExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) EqualityExpression(i int) IEqualityExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *LogicalAndExpressionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(ExprParserAND)
}

func (s *LogicalAndExpressionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserAND, i)
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitLogicalAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) LogicalAndExpression() (localctx ILogicalAndExpressionContext) {
	localctx = NewLogicalAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ExprParserRULE_logicalAndExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.EqualityExpression()
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == ExprParserAND {
		{
			p.SetState(137)
			p.Match(ExprParserAND)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(138)
			p.EqualityExpression()
		}


		p.SetState(143)
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


// IEqualityExpressionContext is an interface to support dynamic dispatch.
type IEqualityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRelationalExpression() []IRelationalExpressionContext
	RelationalExpression(i int) IRelationalExpressionContext
	AllEQ() []antlr.TerminalNode
	EQ(i int) antlr.TerminalNode
	AllNE() []antlr.TerminalNode
	NE(i int) antlr.TerminalNode

	// IsEqualityExpressionContext differentiates from other interfaces.
	IsEqualityExpressionContext()
}

type EqualityExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExpressionContext() *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_equalityExpression
	return p
}

func InitEmptyEqualityExpressionContext(p *EqualityExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_equalityExpression
}

func (*EqualityExpressionContext) IsEqualityExpressionContext() {}

func NewEqualityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_equalityExpression

	return p
}

func (s *EqualityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExpressionContext) AllRelationalExpression() []IRelationalExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationalExpressionContext); ok {
			len++
		}
	}

	tst := make([]IRelationalExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationalExpressionContext); ok {
			tst[i] = t.(IRelationalExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExpressionContext) RelationalExpression(i int) IRelationalExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *EqualityExpressionContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(ExprParserEQ)
}

func (s *EqualityExpressionContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserEQ, i)
}

func (s *EqualityExpressionContext) AllNE() []antlr.TerminalNode {
	return s.GetTokens(ExprParserNE)
}

func (s *EqualityExpressionContext) NE(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserNE, i)
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) EqualityExpression() (localctx IEqualityExpressionContext) {
	localctx = NewEqualityExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ExprParserRULE_equalityExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.RelationalExpression()
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == ExprParserNE || _la == ExprParserEQ {
		{
			p.SetState(145)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ExprParserNE || _la == ExprParserEQ) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(146)
			p.RelationalExpression()
		}


		p.SetState(151)
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


// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RegexExpression() IRegexExpressionContext
	AllAdditiveExpression() []IAdditiveExpressionContext
	AdditiveExpression(i int) IAdditiveExpressionContext
	AllGT() []antlr.TerminalNode
	GT(i int) antlr.TerminalNode
	AllGE() []antlr.TerminalNode
	GE(i int) antlr.TerminalNode
	AllLT() []antlr.TerminalNode
	LT(i int) antlr.TerminalNode
	AllLE() []antlr.TerminalNode
	LE(i int) antlr.TerminalNode
	AllMA() []antlr.TerminalNode
	MA(i int) antlr.TerminalNode
	AllNM() []antlr.TerminalNode
	NM(i int) antlr.TerminalNode

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_relationalExpression
	return p
}

func InitEmptyRelationalExpressionContext(p *RelationalExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_relationalExpression
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) RegexExpression() IRegexExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexExpressionContext)
}

func (s *RelationalExpressionContext) AllAdditiveExpression() []IAdditiveExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveExpressionContext); ok {
			tst[i] = t.(IAdditiveExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RelationalExpressionContext) AdditiveExpression(i int) IAdditiveExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *RelationalExpressionContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(ExprParserGT)
}

func (s *RelationalExpressionContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserGT, i)
}

func (s *RelationalExpressionContext) AllGE() []antlr.TerminalNode {
	return s.GetTokens(ExprParserGE)
}

func (s *RelationalExpressionContext) GE(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserGE, i)
}

func (s *RelationalExpressionContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(ExprParserLT)
}

func (s *RelationalExpressionContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserLT, i)
}

func (s *RelationalExpressionContext) AllLE() []antlr.TerminalNode {
	return s.GetTokens(ExprParserLE)
}

func (s *RelationalExpressionContext) LE(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserLE, i)
}

func (s *RelationalExpressionContext) AllMA() []antlr.TerminalNode {
	return s.GetTokens(ExprParserMA)
}

func (s *RelationalExpressionContext) MA(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserMA, i)
}

func (s *RelationalExpressionContext) AllNM() []antlr.TerminalNode {
	return s.GetTokens(ExprParserNM)
}

func (s *RelationalExpressionContext) NM(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserNM, i)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitRelationalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ExprParserRULE_relationalExpression)
	var _la int

	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.RegexExpression()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.AdditiveExpression()
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 112253265248256) != 0) {
			{
				p.SetState(154)
				_la = p.GetTokenStream().LA(1)

				if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 112253265248256) != 0)) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(155)
				p.AdditiveExpression()
			}


			p.SetState(160)
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


// IRegexExpressionContext is an interface to support dynamic dispatch.
type IRegexExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	MA() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsRegexExpressionContext differentiates from other interfaces.
	IsRegexExpressionContext()
}

type RegexExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexExpressionContext() *RegexExpressionContext {
	var p = new(RegexExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_regexExpression
	return p
}

func InitEmptyRegexExpressionContext(p *RegexExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_regexExpression
}

func (*RegexExpressionContext) IsRegexExpressionContext() {}

func NewRegexExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexExpressionContext {
	var p = new(RegexExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_regexExpression

	return p
}

func (s *RegexExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, 0)
}

func (s *RegexExpressionContext) MA() antlr.TerminalNode {
	return s.GetToken(ExprParserMA, 0)
}

func (s *RegexExpressionContext) STRING() antlr.TerminalNode {
	return s.GetToken(ExprParserSTRING, 0)
}

func (s *RegexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RegexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterRegexExpression(s)
	}
}

func (s *RegexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitRegexExpression(s)
	}
}

func (s *RegexExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitRegexExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) RegexExpression() (localctx IRegexExpressionContext) {
	localctx = NewRegexExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ExprParserRULE_regexExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(ExprParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(164)
		p.Match(ExprParserMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Match(ExprParserSTRING)
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


// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMultiplicativeExpression() []IMultiplicativeExpressionContext
	MultiplicativeExpression(i int) IMultiplicativeExpressionContext
	AllADD() []antlr.TerminalNode
	ADD(i int) antlr.TerminalNode
	AllSUB() []antlr.TerminalNode
	SUB(i int) antlr.TerminalNode

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_additiveExpression
	return p
}

func InitEmptyAdditiveExpressionContext(p *AdditiveExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_additiveExpression
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) AllMultiplicativeExpression() []IMultiplicativeExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeExpressionContext); ok {
			tst[i] = t.(IMultiplicativeExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) MultiplicativeExpression(i int) IMultiplicativeExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AllADD() []antlr.TerminalNode {
	return s.GetTokens(ExprParserADD)
}

func (s *AdditiveExpressionContext) ADD(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserADD, i)
}

func (s *AdditiveExpressionContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(ExprParserSUB)
}

func (s *AdditiveExpressionContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserSUB, i)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ExprParserRULE_additiveExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.MultiplicativeExpression()
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == ExprParserADD || _la == ExprParserSUB {
		{
			p.SetState(168)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ExprParserADD || _la == ExprParserSUB) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(169)
			p.MultiplicativeExpression()
		}


		p.SetState(174)
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


// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnaryExpression() []IUnaryExpressionContext
	UnaryExpression(i int) IUnaryExpressionContext
	AllMUL() []antlr.TerminalNode
	MUL(i int) antlr.TerminalNode
	AllDIV() []antlr.TerminalNode
	DIV(i int) antlr.TerminalNode

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_multiplicativeExpression
	return p
}

func InitEmptyMultiplicativeExpressionContext(p *MultiplicativeExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_multiplicativeExpression
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) AllUnaryExpression() []IUnaryExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			len++
		}
	}

	tst := make([]IUnaryExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryExpressionContext); ok {
			tst[i] = t.(IUnaryExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) UnaryExpression(i int) IUnaryExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *MultiplicativeExpressionContext) AllMUL() []antlr.TerminalNode {
	return s.GetTokens(ExprParserMUL)
}

func (s *MultiplicativeExpressionContext) MUL(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserMUL, i)
}

func (s *MultiplicativeExpressionContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(ExprParserDIV)
}

func (s *MultiplicativeExpressionContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserDIV, i)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ExprParserRULE_multiplicativeExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.UnaryExpression()
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == ExprParserMUL || _la == ExprParserDIV {
		{
			p.SetState(176)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ExprParserMUL || _la == ExprParserDIV) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(177)
			p.UnaryExpression()
		}


		p.SetState(182)
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


// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Constant() IConstantContext
	IDENTIFIER() antlr.TerminalNode
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	REGEX() antlr.TerminalNode
	STRING() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	FunctionCall() IFunctionCallContext
	NOT() antlr.TerminalNode
	SUB() antlr.TerminalNode

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_unaryExpression
	return p
}

func InitEmptyUnaryExpressionContext(p *UnaryExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_unaryExpression
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) Constant() IConstantContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *UnaryExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, 0)
}

func (s *UnaryExpressionContext) INT() antlr.TerminalNode {
	return s.GetToken(ExprParserINT, 0)
}

func (s *UnaryExpressionContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ExprParserFLOAT, 0)
}

func (s *UnaryExpressionContext) REGEX() antlr.TerminalNode {
	return s.GetToken(ExprParserREGEX, 0)
}

func (s *UnaryExpressionContext) STRING() antlr.TerminalNode {
	return s.GetToken(ExprParserSTRING, 0)
}

func (s *UnaryExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserLPAREN, 0)
}

func (s *UnaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserRPAREN, 0)
}

func (s *UnaryExpressionContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *UnaryExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(ExprParserNOT, 0)
}

func (s *UnaryExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(ExprParserSUB, 0)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ExprParserRULE_unaryExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == ExprParserSUB || _la == ExprParserNOT {
		{
			p.SetState(183)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ExprParserSUB || _la == ExprParserNOT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ExprParserCONST_NOTHING, ExprParserCONST_FRAUD, ExprParserCONST_CANCELLED, ExprParserCONST_INVALID, ExprParserCONST_ALERT, ExprParserCOSNT_CONFLICT, ExprParserCONST_UNGRADED, ExprParserCOSNT_OBSCURED, ExprParserCONST_ABSENT, ExprParserCONST_PRESENT, ExprParserCONST_EXCUSED, ExprParserCONST_TOOLOW:
		{
			p.SetState(186)
			p.Constant()
		}


	case ExprParserIDENTIFIER:
		{
			p.SetState(187)
			p.Match(ExprParserIDENTIFIER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case ExprParserINT:
		{
			p.SetState(188)
			p.Match(ExprParserINT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case ExprParserFLOAT:
		{
			p.SetState(189)
			p.Match(ExprParserFLOAT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case ExprParserREGEX:
		{
			p.SetState(190)
			p.Match(ExprParserREGEX)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case ExprParserSTRING:
		{
			p.SetState(191)
			p.Match(ExprParserSTRING)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case ExprParserLPAREN:
		{
			p.SetState(192)
			p.Match(ExprParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Expression()
		}
		{
			p.SetState(194)
			p.Match(ExprParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case ExprParserMAX, ExprParserMIN, ExprParserROUND:
		{
			p.SetState(196)
			p.FunctionCall()
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


// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function() IFunctionContext
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) Function() IFunctionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserLPAREN, 0)
}

func (s *FunctionCallContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *FunctionCallContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserRPAREN, 0)
}

func (s *FunctionCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ExprParserCOMMA)
}

func (s *FunctionCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserCOMMA, i)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprParserVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExprParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ExprParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Function()
	}
	{
		p.SetState(200)
		p.Match(ExprParserLPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Expression()
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == ExprParserCOMMA {
		{
			p.SetState(202)
			p.Match(ExprParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Expression()
		}


		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(209)
		p.Match(ExprParserRPAREN)
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


