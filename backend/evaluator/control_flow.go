// contains loops, ifs, apply......

package evaluator

import (
	"ParserGen/parser"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

func (v *Evaluator) VisitAlgorithm(ctx *parser.AlgorithmContext) interface{} {
	for _, c := range ctx.GetChildren() {
		if tree, ok := c.(antlr.ParseTree); ok {
			tree.Accept(v)
		}
	}
	return nil
}

func (v *Evaluator) VisitApply(ctx *parser.ApplyContext) interface{} {
	lambdaCtx := ctx.Lambda()
	lambdaVar := lambdaCtx.IDENTIFIER().GetText()
	lambdaExpr := lambdaCtx.ConditionalExpression()

	listCtx := ctx.List()
	identifiers := listCtx.AllIDENTIFIER()

	//fmt.Println("Applying lambda to list:")
	for _, id := range identifiers {
		varName := id.GetText()
		if val, exists := v.Memory[varName]; exists {
			v.Memory[lambdaVar] = val
			result := v.Visit(lambdaExpr)
			v.Memory[varName] = result
		}
	}

	delete(v.Memory, lambdaVar)
	return nil
}

func (v *Evaluator) VisitNormal_operation(ctx *parser.Normal_operationContext) interface{} {
	hasExplanation := ctx.EXPLAIN() != nil
	hasDescription := ctx.STRING() != nil

	v.IsExplaining = true

	if hasExplanation {
		originalExpr := ctx.Expression().GetText()

		builder := &ExpressionBuilder{Memory: v.Memory}
		substitutedExpr := builder.BuildSubstituted(ctx.Expression())

		v.CurrentExplanation = &ExplanationData{
			OriginalExpression:    originalExpr,
			SubstitutedExpression: substitutedExpr,
		}
	} else {
		v.CurrentExplanation = &ExplanationData{}
	}

	result := v.Visit(ctx.Expression())
	var resultStr string
	switch val := result.(type) {
	case float64:
		resultStr = fmt.Sprintf("%.2f", val)
	case bool:
		resultStr = fmt.Sprintf("%v", val)
	case string:
		resultStr = val
	default:
		resultStr = fmt.Sprintf("%v", val)
	}

	if hasDescription {
		stringDescription, err := strconv.Unquote(fmt.Sprintf("%v", ctx.STRING()))
		if err != nil {
			panic(err)
		}
		v.CurrentExplanation.Description = stringDescription
	}

	if hasExplanation {
		v.CurrentExplanation.Result = resultStr
		v.AllExplanations = append(v.AllExplanations, *v.CurrentExplanation)
	} else {
		v.CurrentExplanation.Result = resultStr
		v.AllExplanations = append(v.AllExplanations, *v.CurrentExplanation)
	}
	v.IsExplaining = false
	v.CurrentExplanation = nil

	return result
}

func (v *Evaluator) VisitWhen_operation(ctx *parser.When_operationContext) interface{} {
	identifierListCtx := ctx.List()
	identifierList := identifierListCtx.AllIDENTIFIER()

	inMemory := true

	for _, id := range identifierList {
		varName := id.GetText()
		if _, exists := v.Memory[varName]; !exists {
			inMemory = false
		}
	}

	if identifierList != nil && inMemory {
		result := v.Visit(ctx.Normal_operation())
		return result
	}
	return nil
}

func (v *Evaluator) VisitIterateExpression(ctx *parser.IterateExpressionContext) interface{} {
	loopVar := ctx.IDENTIFIER().GetText()
	constantListCtx := ctx.ConstantList()
	constants := constantListCtx.AllConstant()

	var results []interface{}

	inConstants := false

	for _, id := range constants {
		varName := id.GetText()
		if val, exists := v.Memory[loopVar]; exists {
			if val == varName {
				inConstants = true
			}
		}
	}

	if len(ctx.AllExpression()) == 2 {
		if inConstants {
			return v.Visit(ctx.Expression(0))
		} else {
			return v.Visit(ctx.Expression(1))
		}
	}
	return results
}

func (v *Evaluator) VisitConditionalExpression(ctx *parser.ConditionalExpressionContext) interface{} {
	result := v.Visit(ctx.LogicalOrExpression())

	if ctx.Expression(0) != nil && ctx.Expression(1) != nil {
		condition, ok := result.(bool)

		if !ok {
			if val, isFloat := result.(float64); isFloat {
				condition = val != 0.0
			} else {
				condition = false
			}
		}

		if condition {
			return v.Visit(ctx.Expression(0))
		} else {
			return v.Visit(ctx.Expression(1))
		}
	}

	return result
}
