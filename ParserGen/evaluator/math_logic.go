// contains mathematical and logical operations

package evaluator

import (
	"ParserGen/parser"

	"github.com/antlr4-go/antlr/v4"
)

func (v *Evaluator) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	// first mul from additive
	result := v.Visit(ctx.MultiplicativeExpression(0))

	if _, ok := result.(string); ok {
		return result
	}

	if _, ok := result.(bool); ok {
		return result
	}

	resultFloat := result.(float64)

	for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
		// upcoming variables...
		right := v.Visit(ctx.MultiplicativeExpression(i)).(float64)

		op := ctx.GetChild(2*i - 1).(antlr.TerminalNode).GetText()

		if op == "+" {
			resultFloat += right
		} else if op == "-" {
			resultFloat -= right
		}
	}
	return resultFloat
}

func (v *Evaluator) VisitConstant(ctx *parser.ConstantContext) interface{} {
	text := ctx.GetText()
	return text
}

func (v *Evaluator) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	// visiting the first unary expr
	result := v.Visit(ctx.UnaryExpression(0))

	if _, ok := result.(string); ok {
		return result
	}

	if _, ok := result.(bool); ok {
		return result
	}

	resultFloat := result.(float64)

	for i := 1; i < len(ctx.AllUnaryExpression()); i++ {
		right := v.Visit(ctx.UnaryExpression(i)).(float64)
		op := ctx.GetChild(2*i - 1).(antlr.TerminalNode).GetText()

		if op == "*" {
			resultFloat *= right
		} else if op == "/" {
			resultFloat /= right
		}
	}
	return resultFloat
}

func (v *Evaluator) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} {
	result := v.Visit(ctx.LogicalAndExpression(0))

	if len(ctx.AllLogicalAndExpression()) == 1 {
		return result
	}

	resultBool, ok := result.(bool)
	if !ok {
		if val, isFloat := result.(float64); isFloat {
			resultBool = val != 0.0
		}
	}

	for i := 1; i < len(ctx.AllLogicalAndExpression()); i++ {
		// upcoming variables...
		rightVal := v.Visit(ctx.LogicalAndExpression(i))
		rightBool, ok := rightVal.(bool)
		if !ok {
			if val, isFloat := rightVal.(float64); isFloat {
				rightBool = val != 0.0
			}
		}

		resultBool = resultBool || rightBool
	}
	return resultBool
}

func (v *Evaluator) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} {
	result := v.Visit(ctx.EqualityExpression(0))

	if len(ctx.AllEqualityExpression()) == 1 {
		return result
	}

	resultBool, ok := result.(bool)
	if !ok {
		if val, isFloat := result.(float64); isFloat {
			resultBool = val != 0.0
		}
	}

	for i := 1; i < len(ctx.AllEqualityExpression()); i++ {
		// upcoming variables...
		rightVal := v.Visit(ctx.EqualityExpression(i))
		rightBool, ok := rightVal.(bool)
		if !ok {
			if val, isFloat := rightVal.(float64); isFloat {
				rightBool = val != 0.0
			}
		}

		resultBool = resultBool && rightBool
	}
	return resultBool
}

func (v *Evaluator) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	result := v.Visit(ctx.RelationalExpression(0))

	if len(ctx.AllRelationalExpression()) == 1 {
		return result
	}

	for i := 1; i < len(ctx.AllRelationalExpression()); i++ {
		rightVal := v.Visit(ctx.RelationalExpression(i))
		op := ctx.GetChild(2*i - 1).(antlr.TerminalNode).GetText()

		lStr, lIsStr := result.(string)
		rStr, rIsStr := rightVal.(string)

		if lIsStr && rIsStr {
			if op == "==" {
				result = lStr == rStr
			} else if op == "!=" {
				result = lStr != rStr
			}
			continue
		}

		lFloat, lIsFloat := result.(float64)
		rFloat, rIsFloat := rightVal.(float64)

		if lIsFloat && rIsFloat {
			if op == "==" {
				result = lFloat == rFloat
			} else if op == "!=" {
				result = lFloat != rFloat
			}
			continue
		}

		lBool, lIsBool := result.(bool)
		rBool, rIsBool := rightVal.(bool)

		if lIsBool && rIsBool {
			if op == "==" {
				result = lBool == rBool
			} else if op == "!=" {
				result = lBool != rBool
			}
			continue
		}

		if op == "==" {
			result = false
		} else if op == "!=" {
			result = true
		}
	}

	return result
}

func (v *Evaluator) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} {
	if ctx.RegexExpression() != nil {
		return v.Visit(ctx.RegexExpression())
	}

	if ctx.AdditiveExpression(0) != nil {
		left := v.Visit(ctx.AdditiveExpression(0))

		if _, ok := left.(string); ok {
			return left
		}

		if _, ok := left.(bool); ok {
			return left
		}

		leftFloat := left.(float64)

		if len(ctx.AllAdditiveExpression()) == 1 {
			return leftFloat
		}

		for i := 1; i < len(ctx.AllAdditiveExpression()); i++ {
			// upcoming variables...
			right := v.Visit(ctx.AdditiveExpression(i)).(float64)

			op := ctx.GetChild(2*i - 1).(antlr.TerminalNode).GetText()

			var result bool

			if op == ">" {
				result = leftFloat > right
			} else if op == ">=" {
				result = leftFloat >= right
			} else if op == "<" {
				result = leftFloat < right
			} else if op == "<=" {
				result = leftFloat <= right
			} else if op == "~" {
				// regex match - TODO
				result = false
			} else if op == "!~" {
				// regex no match - TODO
				result = false
			}

			if !result {
				return false
			}
			leftFloat = right
		}
		return true
	}

	return false
}
