package lexer

import (
	"ciri/src/goyacc"
	"ciri/src/token"
)

//Lex interface for yacc, return generated types
func (l *Lexer) Lex(parserVal *goyacc.YySymType) int {
	tok := l.NextToken()

	switch tok.Type {
	case token.VAR:
		parserVal.St = tok.Literal
		return goyacc.VAR
	case token.ID:
		parserVal.St = tok.Literal
		return goyacc.ID
	case token.FLOAT_TYPE:
		parserVal.St = tok.Literal
		return goyacc.FLOAT_TYPE
	case token.INT_TYPE:
		parserVal.St = tok.Literal
		return goyacc.INT_TYPE
	case token.IF:
		parserVal.St = tok.Literal
		return goyacc.IF
	case token.ELSE:
		parserVal.St = tok.Literal
		return goyacc.ELSE
	case token.PROGRAM:
		parserVal.St = tok.Literal
		return goyacc.PROGRAM
	case token.PRINT:
		parserVal.St = tok.Literal
		return goyacc.PRINT
	case token.COLON:
		parserVal.St = tok.Literal
		return ':'
	case token.COMMA:
		parserVal.St = tok.Literal
		return ','
	case token.OPEN_BRACE:
		parserVal.St = tok.Literal
		return '{'
	case token.PLUS:
		parserVal.St = tok.Literal
		return '+'
	case token.MINUS:
		parserVal.St = tok.Literal
		return '-'
	case token.DIVIDE:
		parserVal.St = tok.Literal
		return '/'
	case token.MULTIPLY:
		parserVal.St = tok.Literal
		return '*'
	case token.CLOSED_BRACE:
		parserVal.St = tok.Literal
		return '}'
	case token.ASSIGN:
		parserVal.St = tok.Literal
		return '='
	case token.OPEN_PARENTHESIS:
		parserVal.St = tok.Literal
		return '('
	case token.CLOSED_PARENTHESIS:
		parserVal.St = tok.Literal
		return ')'
	case token.INT:
		parserVal.St = tok.Literal
		return goyacc.CTE_I
	case token.SEMICOLON:
		parserVal.St = tok.Literal
		return ';'
	case token.LESS_THAN:
		parserVal.St = tok.Literal
		return '<'
	case token.GREATER_THAN:
		parserVal.St = tok.Literal
		return '>'
	case token.STRING:
		parserVal.St = tok.Literal
		return goyacc.CTE_STRING
	case token.FLOAT:
		parserVal.St = tok.Literal
		return goyacc.CTE_F
	default:
		return int(tok.LineNumber)
	}

}
