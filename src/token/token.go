package token

import (
	"regexp"
)

type Type string

type Token struct {
	Type            Type
	LineNumber      uint32
	CharacterNumber uint32
	Literal         string
	IsKeyword       bool
}

type Keyword struct {
	Regex string
	Type  Type
}

var IDENT = Keyword{Regex: "[a-zA-Z](a-zA-Z]|[0-9])*", Type: ID}
var FLOAT_IDENT = Keyword{Regex: "[+-]?([0-9]*[.])[0-9]+", Type: FLOAT}
var INT_IDENT = Keyword{Regex: "[-+]?[0-9]+", Type: INT}

var simpleKeywords = map[string]Keyword{
	"var":     Keyword{Type: VAR},
	"int":     Keyword{Type: INT_TYPE},
	"float":   Keyword{Type: FLOAT_TYPE},
	"<>":      Keyword{Type: LESS_THEN_GREAT},
	"program": Keyword{Type: PROGRAM},
	"true":    Keyword{Type: TRUE},
	"false":   Keyword{Type: FALSE},
	"if":      Keyword{Type: IF},
	"else":    Keyword{Type: ELSE},
	"print":   Keyword{Type: PRINT},
}

const (
	TRUE  = "TRUE"
	FALSE = "FALSE"
	IF    = "IF"
	ELSE  = "ELSE"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"

	GREATER_THAN       = ">"
	LESS_THAN          = "<"
	LESS_THEN_GREAT    = "<>"
	OPEN_PARENTHESIS   = "("
	CLOSED_PARENTHESIS = ")"
	OPEN_BRACE         = "{"
	CLOSED_BRACE       = "}"

	COMMA     = ","
	SEMICOLON = ";"
	ASSIGN    = "="
	COLON     = ":"

	PROGRAM = "program"
	VAR     = "var"
	ID      = "ID"
	PRINT   = "print"
	STRING  = "STRING"

	INT_TYPE   = "INT_TYPE"
	FLOAT_TYPE = "FLOAT_TYPE"
	INT        = "INT"
	FLOAT      = "FLOAT"
)

func LookupIdentifier(keyword Keyword, potentialKeyword string) Type {
	valid, err := regexp.MatchString(keyword.Regex, potentialKeyword)
	if valid && err == nil {
		return keyword.Type
	}
	return ILLEGAL

}

func LookupSimpleKeyword(ident string) Type {
	if tok, ok := simpleKeywords[ident]; ok {
		return tok.Type
	}
	return ILLEGAL
}
