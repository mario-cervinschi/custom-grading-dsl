// contains variables, assignments, functions, regex..

package evaluator

import (
	"ParserGen/parser"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func (v *Evaluator) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	if ctx.IterateExpression() != nil {
		return v.Visit(ctx.IterateExpression())
	}

	if ctx.ConditionalExpression() != nil {
		return v.Visit(ctx.ConditionalExpression())
	}

	if ctx.AssignmentExpression() != nil {
		return v.Visit(ctx.AssignmentExpression())
	}

	return 0.0
}

func (v *Evaluator) VisitAssignmentExpression(ctx *parser.AssignmentExpressionContext) interface{} {
	varName := ctx.IDENTIFIER().GetText()

	val := v.Visit(ctx.Expression())
	v.CurrentExplanation.Variable = varName

	v.Memory[varName] = val
	//switch v := val.(type) {
	//case float64:
	//	fmt.Printf("Assigned %s = %.2f\n", varName, v)
	//case bool:
	//	fmt.Printf("Assigned %s = %v\n", varName, v)
	//case string:
	//	fmt.Printf("Assigned %s = %s\n", varName, v)
	//default:
	//	fmt.Printf("Assigned %s = %v\n", varName, v)
	//}
	return val
}

func (v *Evaluator) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	hasNot := false
	hasNeg := false

	if ctx.NOT() != nil {
		hasNot = true
	}
	if ctx.SUB() != nil {
		hasNeg = true
	}

	var result interface{}

	if ctx.LPAREN() != nil {
		return v.Visit(ctx.Expression())
	} else if ctx.Constant() != nil {
		return v.Visit(ctx.Constant())
	} else if ctx.INT() != nil {
		val, _ := strconv.ParseFloat(ctx.INT().GetText(), 64)
		return val
	} else if ctx.FLOAT() != nil {
		val, _ := strconv.ParseFloat(ctx.FLOAT().GetText(), 64)
		return val
	} else if ctx.IDENTIFIER() != nil {
		varName := ctx.IDENTIFIER().GetText()
		val, exists := v.Memory[varName]
		if !exists {
			// fmt.Printf("Variable %s not found. using 0.0\n", varName)
			return 0.0
		} else {
			return val
		}
	} else if ctx.STRING() != nil {
		val, _ := strconv.Unquote(ctx.STRING().GetText())
		return val
	} else if ctx.FunctionCall() != nil {
		result = v.Visit(ctx.FunctionCall())
	} else {
		result = 0.0
	}

	if hasNeg {
		if val, ok := result.(float64); ok {
			result = -val
		}
	}

	if hasNot {
		result = !result.(bool)
	}

	return result
}

func (v *Evaluator) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	funcName := ctx.Function().GetText()
	args := ctx.AllExpression()

	var values []float64
	for _, arg := range args {
		val := v.Visit(arg)
		if fval, ok := val.(float64); ok {
			values = append(values, fval)
		}
	}

	var result float64

	switch funcName {
	case "max":
		if len(values) == 0 {
			return 0.0
		}
		result = values[0]
		for _, v := range values[1:] {
			if v > result {
				result = v
			}
		}
	case "min":
		if len(values) == 0 {
			return 0.0
		}
		result = values[0]
		for _, v := range values[1:] {
			if v < result {
				result = v
			}
		}
	case "round":
		if len(values) == 0 {
			return 0.0
		}
		result = math.Round(values[0])
	default:
		fmt.Printf("Unknown function: %s\n", funcName)
		return 0.0
	}

	return result
}

func (v *Evaluator) VisitRegexExpression(ctx *parser.RegexExpressionContext) interface{} {
	varName := ctx.IDENTIFIER().GetText()
	pattern, _ := strconv.Unquote(ctx.STRING().GetText())

	varVal, exists := v.Memory[varName]
	if !exists {
		return false
	}

	strVal, ok := varVal.(string)
	if !ok {
		strVal = fmt.Sprintf("%v", varVal)
	}

	matched, err := regexp.MatchString(pattern, strVal)
	if err != nil {
		fmt.Printf("Regex error: %v\n", err)
		return false
	}

	return matched
}
