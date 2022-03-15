package goyacc

import "ciri/src/token"

// Parse parses the input and returns the result.
func Parse(input string) ([]token.Token, error) {
	l := New(input)
	_ = yyParse(l)
	return l.Tokens, l.GetError()
}
