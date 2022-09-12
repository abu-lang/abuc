lexer grammar AbuLexer;

AND             : 'and' ;
OR              : 'or'  ;
NOT             : 'not' ;

ABS             : 'abs' ;
DOT             : '.'   ;

PLUS            : '+' ;
MINUS           : '-' ;
DIV             : '/' ;
MUL             : '*' ;
MOD             : '%' ;

EQUALS          : '==';
GT              : '>' ;
LT              : '<' ;
GTE             : '>=';
LTE             : '<=';
NOTEQUALS       : '!=';

DOUBLECOLON     : '::';

COLON           : ':' ;
SEMICOLON       : ';' ;
COLONEQUAL      : ':=';
EQUALSIGN       : '=' ;

RL_BRACKET      : '(' ;
RR_BRACKET      : ')' ;
CL_BRACKET      : '{' ;
CR_BRACKET      : '}' ;

HAS             : 'has'      ;
WHERE           : 'where'    ;
PHYSICAL        : 'physical' ;
LOGICAL         : 'logical'  ;
INPUT           : 'input'    ;
OUTPUT          : 'output'   ;
RULE            : 'rule'     ;
THIS            : 'this' DOT ;
EXT             : 'ext' DOT  ;
ON              : 'on'       ;
FOR             : 'for'      ;
ALL             : 'all'      ;
DO              : 'do'       ;

BOOLEAN         : 'boolean' ;
INTEGER         : 'integer' ;
DECIMAL         : 'decimal' ;
STRING          : 'string'  ;

TRUE            : 'true'  ;
FALSE           : 'false' ;

ID              : [a-zA-Z][a-zA-Z0-9]* ;

QUOTED_STRING   :  '"' ~('"')* '"'       ;
DEC_LITERAL     : INT_LITERAL DOT [0-9]* ;
INT_LITERAL     : '0'
                | [1-9][0-9]*
                ;

// IGNORED TOKENS
WS              : [ \t\r\n]+    -> channel(HIDDEN) ;
COMMENT         : '\\@' .*? '@\\' -> skip          ;
LINE_COMMENT    : '#' ~[\r\n]* -> skip             ;