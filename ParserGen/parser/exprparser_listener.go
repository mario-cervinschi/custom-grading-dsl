// Code generated from parser/ExprParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ExprParser
import "github.com/antlr4-go/antlr/v4"

// ExprParserListener is a complete listener for a parse tree produced by ExprParser.
type ExprParserListener interface {
	antlr.ParseTreeListener

	// EnterAlgorithm is called when entering the algorithm production.
	EnterAlgorithm(c *AlgorithmContext)

	// EnterApply is called when entering the apply production.
	EnterApply(c *ApplyContext)

	// EnterNormal_operation is called when entering the normal_operation production.
	EnterNormal_operation(c *Normal_operationContext)

	// EnterWhen_operation is called when entering the when_operation production.
	EnterWhen_operation(c *When_operationContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterConstantList is called when entering the constantList production.
	EnterConstantList(c *ConstantListContext)

	// EnterLambda is called when entering the lambda production.
	EnterLambda(c *LambdaContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterIterateExpression is called when entering the iterateExpression production.
	EnterIterateExpression(c *IterateExpressionContext)

	// EnterConditionalExpression is called when entering the conditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterLogicalOrExpression is called when entering the logicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterLogicalAndExpression is called when entering the logicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterRegexExpression is called when entering the regexExpression production.
	EnterRegexExpression(c *RegexExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// ExitAlgorithm is called when exiting the algorithm production.
	ExitAlgorithm(c *AlgorithmContext)

	// ExitApply is called when exiting the apply production.
	ExitApply(c *ApplyContext)

	// ExitNormal_operation is called when exiting the normal_operation production.
	ExitNormal_operation(c *Normal_operationContext)

	// ExitWhen_operation is called when exiting the when_operation production.
	ExitWhen_operation(c *When_operationContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitConstantList is called when exiting the constantList production.
	ExitConstantList(c *ConstantListContext)

	// ExitLambda is called when exiting the lambda production.
	ExitLambda(c *LambdaContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitIterateExpression is called when exiting the iterateExpression production.
	ExitIterateExpression(c *IterateExpressionContext)

	// ExitConditionalExpression is called when exiting the conditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitLogicalOrExpression is called when exiting the logicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitLogicalAndExpression is called when exiting the logicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitRegexExpression is called when exiting the regexExpression production.
	ExitRegexExpression(c *RegexExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)
}
