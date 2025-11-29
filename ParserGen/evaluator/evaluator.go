// contains main structure and constructor

package evaluator

import (
	"ParserGen/parser"
	"github.com/antlr4-go/antlr/v4"
)

type Variables map[string]interface{}

type ExplanationData struct {
	OriginalExpression    string
	SubstitutedExpression string
	Result                string
	Description           string
}

type Evaluator struct {
	*parser.BaseExprParserVisitor
	Memory             Variables
	CurrentExplanation *ExplanationData
	IsExplaining       bool
	AllExplanations    []ExplanationData
}

func NewEvaluator(vars Variables) *Evaluator {
	return &Evaluator{
		BaseExprParserVisitor: &parser.BaseExprParserVisitor{},
		Memory:                vars,
		CurrentExplanation:    nil,
		IsExplaining:          false,
		AllExplanations:       []ExplanationData{},
	}
}

func (v *Evaluator) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}
