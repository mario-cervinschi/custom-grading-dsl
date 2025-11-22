// contains main structure and constructor

package evaluator

import (
	"ParserGen/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Variables map[string]interface{}

type Evaluator struct {
	*parser.BaseExprParserVisitor
	Memory Variables
}

func NewEvaluator(vars Variables) *Evaluator {
	return &Evaluator{
		BaseExprParserVisitor: &parser.BaseExprParserVisitor{},
		Memory:                vars,
	}
}

func (v *Evaluator) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}
