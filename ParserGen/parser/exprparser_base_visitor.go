// Code generated from parser/ExprParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ExprParser

import "github.com/antlr4-go/antlr/v4"

type BaseExprParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseExprParserVisitor) VisitAlgorithm(ctx *AlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitApply(ctx *ApplyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitExplainBasic(ctx *ExplainBasicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitExplainConditional(ctx *ExplainConditionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitExplainOverride(ctx *ExplainOverrideContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitConstantList(ctx *ConstantListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitLambda(ctx *LambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitIterateExpression(ctx *IterateExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitRegexExpression(ctx *RegexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}
