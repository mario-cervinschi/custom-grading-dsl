lexer grammar ExprLexer;

@header {
    package parser
}

APPLY: A P P L Y;
TO: T O;

WHEN: W H E N;
EXISTS: E X I S T S;
EXPLAIN: E X P L A I N;

MAX: M A X;
MIN: M I N;
ROUND: R O U N D;

IN: I N;

CONST_NOTHING: 'nothing';
CONST_FRAUD: 'fraud';
CONST_CANCELLED: 'cancelled';
CONST_INVALID: 'invalid';
CONST_ALERT: 'alert';
COSNT_CONFLICT: 'conflict';
CONST_UNGRADED: 'ungraded';
COSNT_OBSCURED: 'obscured';
CONST_ABSENT: 'absent';
CONST_PRESENT: 'present';
CONST_EXCUSED: 'excused';
CONST_TOOLOW: 'toolow';

IDENTIFIER: [A-Za-z][A-Za-z0-9_]*;

LOP: '=>';
ASK: '?';
COLON: ':';
SEMICOLON: ';';
COMMA: ',';
ASSIGN: '=';

LBRACK: '[';
RBRACK: ']';
LPAREN: '(';
RPAREN: ')';

STRING: '"'  (~('"' | '\r' | '\n') | '""' | NEWLINE)* '"';
REGEX: (MA | NM) STRING;

NM: '!~';
MA: '~';

ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';

GT: '>';
GE: '>=';
NE: '!=';
EQ: '==';
LE: '<=';
LT: '<';

AND: '&&';
OR: '||';
NOT: '!';

FLOAT: [0-9]+ '.' [0-9]*;
INT: [0-9]+;
SPACES: [ \t\r\n]+ -> channel(HIDDEN);

fragment NEWLINE: '\r'? '\n';

fragment A:[aA];
fragment B:[bB];
fragment C:[cC];
fragment D:[dD];
fragment E:[eE];
fragment F:[fF];
fragment G:[gG];
fragment H:[hH];
fragment I:[iI];
fragment J:[jJ];
fragment K:[kK];
fragment L:[lL];
fragment M:[mM];
fragment N:[nN];
fragment O:[oO];
fragment P:[pP];
fragment Q:[qQ];
fragment R:[rR];
fragment S:[sS];
fragment T:[tT];
fragment U:[uU];
fragment V:[vV];
fragment W:[wW];
fragment X:[xX];
fragment Y:[yY];
fragment Z:[zZ];