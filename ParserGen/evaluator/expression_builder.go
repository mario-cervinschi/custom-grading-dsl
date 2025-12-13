package evaluator

import (
	"ParserGen/parser"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type ExpressionBuilder struct {
	Memory Variables
}

func (eb *ExpressionBuilder) BuildSubstituted(ctx antlr.ParseTree) string {
	if ctx == nil {
		return ""
	}

	switch node := ctx.(type) {

	case *parser.ExpressionContext:
		if node.AssignmentExpression() != nil {
			return eb.BuildSubstituted(node.AssignmentExpression())
		}
		if node.ConditionalExpression() != nil {
			return eb.BuildSubstituted(node.ConditionalExpression())
		}
		if node.IterateExpression() != nil {
			return eb.BuildSubstituted(node.IterateExpression())
		}
		return ""

	case *parser.AssignmentExpressionContext:
		varName := node.IDENTIFIER().GetText()
		expr := eb.BuildSubstituted(node.Expression())
		return fmt.Sprintf("%s=%s", varName, expr)

	case *parser.ConditionalExpressionContext:
		cond := eb.BuildSubstituted(node.LogicalOrExpression())
		if node.Expression(0) != nil && node.Expression(1) != nil {
			trueExpr := eb.BuildSubstituted(node.Expression(0))
			falseExpr := eb.BuildSubstituted(node.Expression(1))
			return fmt.Sprintf("%s?%s:%s", cond, trueExpr, falseExpr)
		}
		return cond

	case *parser.LogicalOrExpressionContext:
		result := eb.BuildSubstituted(node.LogicalAndExpression(0))
		for i := 1; i < len(node.AllLogicalAndExpression()); i++ {
			right := eb.BuildSubstituted(node.LogicalAndExpression(i))
			result = fmt.Sprintf("%s||%s", result, right)
		}
		return result

	case *parser.LogicalAndExpressionContext:
		result := eb.BuildSubstituted(node.EqualityExpression(0))
		for i := 1; i < len(node.AllEqualityExpression()); i++ {
			right := eb.BuildSubstituted(node.EqualityExpression(i))
			result = fmt.Sprintf("%s&&%s", result, right)
		}
		return result

	case *parser.EqualityExpressionContext:
		result := eb.BuildSubstituted(node.RelationalExpression(0))
		for i := 1; i < len(node.AllRelationalExpression()); i++ {
			op := node.GetChild(2*i - 1).(antlr.TerminalNode).GetText()
			right := eb.BuildSubstituted(node.RelationalExpression(i))
			result = fmt.Sprintf("%s%s%s", result, op, right)
		}
		return result

	case *parser.RelationalExpressionContext:
		if node.RegexExpression() != nil {
			return eb.BuildSubstituted(node.RegexExpression())
		}
		if node.AdditiveExpression(0) != nil {
			result := eb.BuildSubstituted(node.AdditiveExpression(0))
			for i := 1; i < len(node.AllAdditiveExpression()); i++ {
				op := node.GetChild(2*i - 1).(antlr.TerminalNode).GetText()
				right := eb.BuildSubstituted(node.AdditiveExpression(i))
				result = fmt.Sprintf("%s%s%s", result, op, right)
			}
			return result
		}
		return ""

	case *parser.AdditiveExpressionContext:
		result := eb.BuildSubstituted(node.MultiplicativeExpression(0))
		for i := 1; i < len(node.AllMultiplicativeExpression()); i++ {
			op := node.GetChild(2*i - 1).(antlr.TerminalNode).GetText()
			right := eb.BuildSubstituted(node.MultiplicativeExpression(i))
			result = fmt.Sprintf("%s%s%s", result, op, right)
		}
		return result

	case *parser.MultiplicativeExpressionContext:
		result := eb.BuildSubstituted(node.UnaryExpression(0))
		for i := 1; i < len(node.AllUnaryExpression()); i++ {
			op := node.GetChild(2*i - 1).(antlr.TerminalNode).GetText()
			right := eb.BuildSubstituted(node.UnaryExpression(i))
			result = fmt.Sprintf("%s%s%s", result, op, right)
		}
		return result

	case *parser.UnaryExpressionContext:
		hasNot := node.NOT() != nil
		hasNeg := node.SUB() != nil

		var inner string

		if node.LPAREN() != nil {
			inner = eb.BuildSubstituted(node.Expression())
			inner = fmt.Sprintf("(%s)", inner)
		} else if node.IDENTIFIER() != nil {
			varName := node.IDENTIFIER().GetText()
			if val, exists := eb.Memory[varName]; exists {
				switch v := val.(type) {
				case float64:
					inner = fmt.Sprintf("%.2f", v)
				case bool:
					inner = fmt.Sprintf("%v", v)
				case string:
					inner = fmt.Sprintf("%s", v)
				default:
					inner = fmt.Sprintf("%v", v)
				}
			} else {
				inner = varName
			}
		} else if node.INT() != nil || node.FLOAT() != nil {
			inner = node.GetText()
		} else if node.STRING() != nil {
			inner = node.GetText()
		} else if node.Constant() != nil {
			inner = node.Constant().GetText()
		} else if node.FunctionCall() != nil {
			inner = eb.BuildSubstituted(node.FunctionCall())
		} else {
			inner = node.GetText()
		}

		if hasNeg {
			inner = "-" + inner
		}
		if hasNot {
			inner = "!" + inner
		}

		return inner

	case *parser.FunctionCallContext:
		funcName := node.Function().GetText()
		args := node.AllExpression()
		var argStrs []string
		for _, arg := range args {
			argStrs = append(argStrs, eb.BuildSubstituted(arg))
		}
		return fmt.Sprintf("%s(%s)", funcName, strings.Join(argStrs, ", "))

	case *parser.RegexExpressionContext:
		varName := node.IDENTIFIER().GetText()
		pattern := node.STRING().GetText()

		if val, exists := eb.Memory[varName]; exists {
			var valStr string
			switch v := val.(type) {
			case string:
				valStr = fmt.Sprintf("%s", v)
			default:
				valStr = fmt.Sprintf("%v", v)
			}
			return fmt.Sprintf("%s ~%s", valStr, pattern)
		}
		return fmt.Sprintf("%s ~%s", varName, pattern)

	case *parser.ConstantContext:
		return node.GetText()

	default:
		if ctx != nil {
			return ctx.GetText()
		}
		return ""
	}
}
