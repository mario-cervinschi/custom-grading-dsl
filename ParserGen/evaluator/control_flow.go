// contains loops, ifs, apply......

package evaluator

import (
	"ParserGen/parser"
	"fmt"
	"strings"

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

	fmt.Println("Applying lambda to list:")
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

func (v *Evaluator) VisitExplainBasic(ctx *parser.ExplainBasicContext) interface{} {
	result := v.Visit(ctx.Expression())

	if val, ok := result.(float64); ok {
		fmt.Printf("Result: %.2f\n", val)
	} else if val, ok := result.(string); ok {
		fmt.Printf("Result: %s\n", val)
	} else if val, ok := result.(bool); ok {
		fmt.Printf("Result: %v\n", val)
	}

	return result
}

func (v *Evaluator) VisitExplainConditional(ctx *parser.ExplainConditionalContext) interface{} {
	return v.Visit(ctx.Expression())
}

func (v *Evaluator) VisitExplainOverride(ctx *parser.ExplainOverrideContext) interface{} {
	targetName := ctx.GetTarget().GetText()

	sourceName := ctx.GetSource().GetText()

	msgRaw := ctx.STRING().GetText()
	msg := strings.Trim(msgRaw, "\"")

	if val, exists := v.Memory[sourceName]; exists {
		v.Memory[targetName] = val

		fmt.Printf("\nOVERRIDE: %s = %v\n", targetName, val)
		fmt.Printf("    mesaj: %s\n", msg)
		return val
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
	fmt.Printf("Iteration results: %v\n", results)
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
