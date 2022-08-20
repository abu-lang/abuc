parser grammar SugaredAbuParser;

options {tokenVocab=SugaredAbuLexer;}

program
    : device+ ecarule*
    ;

device
    : ID COLON stringLiteral CL_BRACKET resList (WHERE expression)? CR_BRACKET (HAS ID+)?
    ;

resList
    : resDecl+
    ;

resDecl
    : PHYSICAL OUTPUT type ID EQUALSIGN expression
    | PHYSICAL INPUT type ID
    | LOGICAL type ID EQUALSIGN expression
    ;

type
    : BOOLEAN | INTEGER | DECIMAL | STRING
    ;

ecarule
    : RULE ID ON ID+ (LET letDecl IN)? task+
    | RULE ID ON ID+ DEFAULT actions (LET letDecl IN)? task*
    ;

letDecl
    : ID COLONEQUAL expression (SEMICOLON ID COLONEQUAL expression)*
    ;

task
    : FOR ALL? expression DO actions
    | FOR ALL? expression DO actions OWISE actions
    ;


actions
    : assignment (SEMICOLON assignment)*
    ;

assignment
    : THIS? ID EQUALSIGN expression
    | EXT ID EQUALSIGN expression
    ;

expression
    : ABS expression
    | NOT expression
    | expression mulDivModOperator expression
    | expression plusMinusOperator expression
    | expression comparisonOperator expression
    | expression DOUBLECOLON expression
    | expression AND expression
    | expression OR expression
    | RL_BRACKET expression RR_BRACKET
    | term
    ;

mulDivModOperator
    : MUL | DIV | MOD
    ;

plusMinusOperator
    : PLUS | MINUS
    ;

comparisonOperator
    : EQUALS | NOTEQUALS | LT | LTE | GT | GTE
    ;

term
    : value
    | resource
    ;

value
    : stringLiteral
    | integerLiteral
    | decimalLiteral
    | booleanLiteral
    ;

resource
    : THIS? ID
    | EXT ID
    ;

decimalLiteral
    : MINUS? DEC_LITERAL
    ;

integerLiteral
    : MINUS? INT_LITERAL
    ;

stringLiteral
    : QUOTED_STRING
    ;

booleanLiteral
    : TRUE
    | FALSE
    ;
