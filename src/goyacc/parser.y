%{
package goyacc

import (
	"ciri/src/token"
)

func setResult(l yyLexer, v []token.Token) {
  l.(*Lexer).Tokens = v
}
%}

%union{
  St string
  Fl float64
  In int
  Ch byte
}

%token<Fl>
        CTE_F
%token<In>
	CTE_I

%token<St>
	VAR
	IF
	ELSE
	ID
	CTE_STRING

	INT_TYPE
	FLOAT_TYPE

	PROGRAM
	PRINT

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies  precedence  for  unary  minus  */


%start programa

%%

programa: PROGRAM ID ':' vars bloque

vars: VAR allVars
    |
allVars: nextId ':' tipo ';' nextVar
nextId: ID
      | ID ',' nextId
nextVar: allVars
       |


bloque: '{' nextStatuto '}'
nextStatuto: estatuto nextStatuto
	   |

estatuto: assign
	| condition
	| print



condition: IF '(' expresion ')' bloque elseBlock ';'
elseBlock: ELSE bloque
	 |

assign: ID '=' expresion ';'

print: PRINT '(' nextPrintExp nextPrint ')' ';'
nextPrintExp: expresion
	   |  CTE_STRING ;
nextPrint: ',' nextPrintExp nextPrint
	 |

tipo: INT_TYPE
    | FLOAT_TYPE


varCte: ID
       | CTE_I
       | CTE_F



factor: '(' expresion ')'
      | cteExp
cteExp: varCte
      | '+' varCte
      | '-' varCte

termino: nextFactor
nextFactor: factor
	  | factor '/' termino
	  | factor '*' termino


exp: termino '+' otherTerm
	| termino '-' otherTerm

otherTerm: termino
	 |


expresion: exp
	 | nextExp

nextExp: exp '>' exp
       | exp '<' exp

