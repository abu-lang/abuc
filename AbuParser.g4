parser grammar AbuParser;

options {tokenVocab=AbuLexer;}

program
    : typeDecl* device ecarule*
    ;

typeDecl
    : DEFINE ID AS CL_BRACKET resField+ CR_BRACKET
    ;

resField
    : ID COLON PHYSICAL OUTPUT type
    | ID COLON PHYSICAL INPUT type
    | ID COLON LOGICAL type
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
    | compResDecl
    ;

compResDecl
    : ID ID (EQUALSIGN RL_BRACKET ID EQUALSIGN expression (COMMA ID EQUALSIGN expression)* RR_BRACKET)?
    ;

type
    : BOOLEAN | INTEGER | DECIMAL | STRING
    ;

ecarule
    : RULE ID ON resource+ task+
    ;

task
    : FOR ALL? expression DO actions
    ;


actions
    : assignment (SEMICOLON assignment)*
    ;

assignment
    : resource EQUALSIGN expression
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
    : simpleResource
    | nestedResource
    ;

simpleResource
    : THIS? ID
    | EXT ID
    ;

nestedResource
    : THIS? ID SL_BRACKET ID SR_BRACKET
    | EXT ID SL_BRACKET ID SR_BRACKET
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
