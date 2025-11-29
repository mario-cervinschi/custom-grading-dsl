// Code generated from parser/ExprParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ExprParser
import "github.com/antlr4-go/antlr/v4"

// BaseExprParserListener is a complete listener for a parse tree produced by ExprParser.
type BaseExprParserListener struct{}

var _ ExprParserListener = &BaseExprParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExprParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExprParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExprParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExprParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAlgorithm is called when production algorithm is entered.
func (s *BaseExprParserListener) EnterAlgorithm(ctx *AlgorithmContext) {}

// ExitAlgorithm is called when production algorithm is exited.
func (s *BaseExprParserListener) ExitAlgorithm(ctx *AlgorithmContext) {}

// EnterApply is called when production apply is entered.
func (s *BaseExprParserListener) EnterApply(ctx *ApplyContext) {}

// ExitApply is called when production apply is exited.
func (s *BaseExprParserListener) ExitApply(ctx *ApplyContext) {}

// EnterNormal_operation is called when production normal_operation is entered.
func (s *BaseExprParserListener) EnterNormal_operation(ctx *Normal_operationContext) {}

// ExitNormal_operation is called when production normal_operation is exited.
func (s *BaseExprParserListener) ExitNormal_operation(ctx *Normal_operationContext) {}

// EnterWhen_operation is called when production when_operation is entered.
func (s *BaseExprParserListener) EnterWhen_operation(ctx *When_operationContext) {}

// ExitWhen_operation is called when production when_operation is exited.
func (s *BaseExprParserListener) ExitWhen_operation(ctx *When_operationContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseExprParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseExprParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseExprParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseExprParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterConstantList is called when production constantList is entered.
func (s *BaseExprParserListener) EnterConstantList(ctx *ConstantListContext) {}

// ExitConstantList is called when production constantList is exited.
func (s *BaseExprParserListener) ExitConstantList(ctx *ConstantListContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BaseExprParserListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BaseExprParserListener) ExitLambda(ctx *LambdaContext) {}

// EnterList is called when production list is entered.
func (s *BaseExprParserListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseExprParserListener) ExitList(ctx *ListContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseExprParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseExprParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BaseExprParserListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BaseExprParserListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterIterateExpression is called when production iterateExpression is entered.
func (s *BaseExprParserListener) EnterIterateExpression(ctx *IterateExpressionContext) {}

// ExitIterateExpression is called when production iterateExpression is exited.
func (s *BaseExprParserListener) ExitIterateExpression(ctx *IterateExpressionContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BaseExprParserListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BaseExprParserListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *BaseExprParserListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *BaseExprParserListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *BaseExprParserListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *BaseExprParserListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseExprParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseExprParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseExprParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseExprParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterRegexExpression is called when production regexExpression is entered.
func (s *BaseExprParserListener) EnterRegexExpression(ctx *RegexExpressionContext) {}

// ExitRegexExpression is called when production regexExpression is exited.
func (s *BaseExprParserListener) ExitRegexExpression(ctx *RegexExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseExprParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseExprParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseExprParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseExprParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseExprParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseExprParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseExprParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseExprParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}
