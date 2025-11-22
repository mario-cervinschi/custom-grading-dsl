parser grammar ExprParser;

@header {
    package parser
}

options {
    tokenVocab = ExprLexer;
}

algorithm: (apply | explain)* EOF;

apply: APPLY lambda TO list SEMICOLON;
explain
    : EXPLAIN expression SEMICOLON                                          # ExplainBasic
    | EXPLAINQ expression SEMICOLON                                         # ExplainConditional
    | EXPLAINQ target=IDENTIFIER ASSIGN source=IDENTIFIER COMMA STRING SEMICOLON # ExplainOverride
    ;

function: MAX | MIN | ROUND;

constant:
    CONST_ABSENT
    | CONST_PRESENT
    | CONST_EXCUSED
    | CONST_NOTHING
    | CONST_FRAUD
    | CONST_CANCELLED
    | CONST_INVALID
    | CONST_ALERT
    | COSNT_CONFLICT
    | CONST_UNGRADED
    | COSNT_OBSCURED
    | CONST_TOOLOW;
constantList: constant ( COMMA constant )*;

lambda: IDENTIFIER LOP conditionalExpression ;
list: IDENTIFIER (COMMA IDENTIFIER)*;

expression:
    iterateExpression
    | conditionalExpression
    | assignmentExpression
    ;

//------------ assignment  -----vvvvvv-----

assignmentExpression:
    IDENTIFIER ASSIGN expression;

//---------------------------

//------------ list-iteration  -----vvvvvv-----

iterateExpression:
    IDENTIFIER IN LBRACK constantList RBRACK (ASK expression COLON expression)?
    ;

//---------------------------

//------------ conditional-expression  -----vvvvvv-----

conditionalExpression:
    logicalOrExpression (ASK expression COLON expression)?
    ;

//---------------------------

//------------ logical-expression  -----vvvvvv-----

logicalOrExpression:
    logicalAndExpression (OR logicalAndExpression)*;

logicalAndExpression:
    equalityExpression (AND equalityExpression)*;

equalityExpression:
    relationalExpression ((EQ | NE) relationalExpression)*;

relationalExpression:
    regexExpression | additiveExpression ((GT | GE | LT | LE | MA | NM) additiveExpression)*;

//---------------------------

//------------ regex-expression  -----vvvvvv-----

regexExpression:
    IDENTIFIER MA STRING;

//---------------------------

//------------ arithmetic-expression  -----vvvvvv-----

additiveExpression:
    multiplicativeExpression ((ADD | SUB) multiplicativeExpression)*;

multiplicativeExpression:
    unaryExpression ((MUL | DIV) unaryExpression)*;

unaryExpression:
    (NOT | SUB)? (
        constant
        | IDENTIFIER
        | INT
        | FLOAT
        | REGEX
        | STRING
        | LPAREN expression RPAREN
        | functionCall
    );


//---------------------------

//------------ function-call  -----vvvvvv-----

functionCall:
    function LPAREN expression (COMMA expression)* RPAREN;

//---------------------------