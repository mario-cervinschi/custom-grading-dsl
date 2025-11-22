// Code generated from parser/ExprParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ExprParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ExprParser.
type ExprParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExprParser#algorithm.
	VisitAlgorithm(ctx *AlgorithmContext) interface{}

	// Visit a parse tree produced by ExprParser#apply.
	VisitApply(ctx *ApplyContext) interface{}

	// Visit a parse tree produced by ExprParser#ExplainBasic.
	VisitExplainBasic(ctx *ExplainBasicContext) interface{}

	// Visit a parse tree produced by ExprParser#ExplainConditional.
	VisitExplainConditional(ctx *ExplainConditionalContext) interface{}

	// Visit a parse tree produced by ExprParser#ExplainOverride.
	VisitExplainOverride(ctx *ExplainOverrideContext) interface{}

	// Visit a parse tree produced by ExprParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by ExprParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by ExprParser#constantList.
	VisitConstantList(ctx *ConstantListContext) interface{}

	// Visit a parse tree produced by ExprParser#lambda.
	VisitLambda(ctx *LambdaContext) interface{}

	// Visit a parse tree produced by ExprParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by ExprParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by ExprParser#assignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by ExprParser#iterateExpression.
	VisitIterateExpression(ctx *IterateExpressionContext) interface{}

	// Visit a parse tree produced by ExprParser#conditionalExpression.
	VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{}

	// Visit a parse tree produced by ExprParser#logicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by ExprParser#logicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by ExprParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by ExprParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by ExprParser#regexExpression.
	VisitRegexExpression(ctx *RegexExpressionContext) interface{}

	// Visit a parse tree produced by ExprParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by ExprParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by ExprParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by ExprParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

}