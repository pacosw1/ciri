package goyacc

import (
	"ciri/src/token"
	"errors"
	"fmt"
)

type Lexer struct {
	input         string
	position      int
	nextPosition  int
	current       byte
	lineNumber    uint32
	Tokens        []token.Token
	lastReadToken token.Token
	err           error
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input, lineNumber: 0, Tokens: make([]token.Token, 0)}
	lexer.readChar()
	return lexer
}

//NextToken reads the next token if valid
func (l *Lexer) NextToken() token.Token {
	var t token.Token
	l.ignoreWhitespaces()

	switch l.current {
	case '=':
		t = l.newToken(token.ASSIGN)
	case '<':
		t = l.newToken(token.LESS_THAN)
	case '>':
		t = l.newToken(token.GREATER_THAN)
	case ';':
		t = l.newToken(token.SEMICOLON)
	case ':':
		t = l.newToken(token.COLON)
	case '(':
		t = l.newToken(token.OPEN_PARENTHESIS)
	case ')':
		t = l.newToken(token.CLOSED_PARENTHESIS)
	case '{':
		t = l.newToken(token.OPEN_BRACE)
	case '}':
		t = l.newToken(token.CLOSED_BRACE)
	case '+':
		t = l.newToken(token.PLUS)
	case '-':
		t = l.newToken(token.MINUS)
	case '/':
		t = l.newToken(token.DIVIDE)
	case ',':
		t = l.newToken(token.COMMA)
	case '*':
		t = l.newToken(token.MULTIPLY)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:

		if isLetter(l.current) {
			t = l.lookupKeyword()
		} else if isDigit(l.current) {
			t = l.lookupNumerics()
		} else if isStringStart(l.current) {
			t = l.lookupString()
		} else {
			t = l.newToken(token.ILLEGAL)
		}
	}
	if !t.IsKeyword {
		l.readChar()
	}
	l.Tokens = append(l.Tokens, t)
	l.lastReadToken = t
	return t
}

// Helpers

func (l *Lexer) ignoreWhitespaces() {
	for l.current == ' ' || l.current == '\t' || l.current == '\n' || l.current == '\r' {
		if l.current == '\n' {
			l.lineNumber += 1
		}
		l.readChar()
	}
}

func isStringStart(ch byte) bool {
	return ch == '"'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

//Yacc interface

func (l *Lexer) GetError() error {
	return l.err
}

func (l *Lexer) Error(s string) {
	s = fmt.Sprintf(" \n %s \n near  %s", s, l.formatError())
	l.err = errors.New(s)
}

func (l *Lexer) formatError() string {
	//show last 5 Tokens
	tokens := l.Tokens
	output := ""
	var tok token.Token
	pos := len(tokens) - 5
	if pos < 0 {
		pos = 0
	}

	for pos < len(tokens) {
		tok = tokens[pos]
		output += tok.Literal + " "

		tokens = tokens[:len(tokens)-1]
		pos += 1
	}

	return fmt.Sprintf("%s %s <--  \n in line: %d", l.lastReadToken.Literal, output, tok.LineNumber)
}

// Lookups

func (l *Lexer) lookupKeyword() token.Token {
	lexCopy := l.makeCopy()

	potentialKeyword := readWord(l)
	tokenType := token.LookupSimpleKeyword(potentialKeyword)
	if tokenType != token.ILLEGAL {
		return l.newKeywordToken(tokenType, potentialKeyword)
	}

	l.resetLookup(lexCopy)
	potentialId := readAlphaNumeric(l)
	tokenType = token.LookupIdentifier(token.IDENT, potentialId)
	if tokenType != token.ILLEGAL {
		return l.newKeywordToken(tokenType, potentialId)
	}

	return l.newToken(token.ILLEGAL)
}

func (l *Lexer) resetLookup(old *Lexer) {
	l.position = old.position
	l.nextPosition = old.nextPosition
	l.lineNumber = old.lineNumber

	l.current = l.input[l.position]
}

func (l *Lexer) lookupString() token.Token {
	raw, unclosed := l.readRawString()
	if unclosed {
		return l.newToken(token.ILLEGAL)
	}
	return l.newKeywordToken(token.STRING, raw)
}

func (l *Lexer) lookupNumerics() token.Token {
	lexCopy := l.makeCopy()
	potentialFloat := readFloat(l)

	tokenType := token.LookupIdentifier(token.FLOAT_IDENT, potentialFloat)
	if tokenType != token.ILLEGAL {
		return l.newKeywordToken(tokenType, potentialFloat)
	}

	l.resetLookup(lexCopy)
	potentialInt := readInt(l)
	tokenType = token.LookupIdentifier(token.INT_IDENT, potentialInt)
	if tokenType != token.ILLEGAL {
		return l.newKeywordToken(tokenType, potentialInt)
	}

	return l.newToken(token.ILLEGAL)
}

// Lookup helpers

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.current = 0
		return
	}
	l.current = l.input[l.nextPosition]
	l.position = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) makeCopy() *Lexer {
	return &Lexer{
		input:        l.input,
		position:     l.position,
		nextPosition: l.nextPosition,
		current:      l.current,
		lineNumber:   l.lineNumber,
	}
}

func readFloat(l *Lexer) string {
	keyword := ""

	for l.current == '.' || isDigit(l.current) {
		keyword += string(l.current)
		l.readChar()
	}
	return keyword
}

func readInt(l *Lexer) string {
	keyword := ""

	for isDigit(l.current) {
		keyword += string(l.current)
		l.readChar()
	}
	return keyword
}

func readWord(l *Lexer) string {
	keyword := ""

	for isLetter(l.current) {
		keyword += string(l.current)
		l.readChar()
	}
	return keyword
}

func readAlphaNumeric(l *Lexer) string {
	keyword := ""

	for isLetter(l.current) || isDigit(l.current) {
		keyword += string(l.current)
		l.readChar()
	}
	return keyword
}

// Token Builders

func (l *Lexer) newKeywordToken(tokenType token.Type, keyword string) token.Token {
	return token.Token{
		Type:            tokenType,
		LineNumber:      l.lineNumber,
		CharacterNumber: uint32(l.position - len(keyword) - 1),
		Literal:         keyword,
		IsKeyword:       true,
	}
}

func (l *Lexer) newToken(tokeType token.Type) token.Token {
	return token.Token{
		Type:            tokeType,
		Literal:         string(l.current),
		LineNumber:      l.lineNumber,
		CharacterNumber: uint32(l.position),
		IsKeyword:       false,
	}
}

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

func (l *Lexer) readRawString() (string, bool) {
	unclosedError := false
	output := string(l.current)
	for {
		l.readChar()
		if l.current == '"' || l.current == 0 {
			if l.current == 0 {
				unclosedError = true
			} else {
				output += string(l.current)
				l.readChar()
			}

			break
		}
		output += string(l.current)
	}

	return output, unclosedError
}

func (l *Lexer) Lex(parserVal *yySymType) int {
	tok := l.NextToken()

	switch tok.Type {
	case token.VAR:
		parserVal.St = tok.Literal
		return VAR
	case token.ID:
		parserVal.St = tok.Literal
		return ID
	case token.FLOAT_TYPE:
		parserVal.St = tok.Literal
		return FLOAT_TYPE
	case token.INT_TYPE:
		parserVal.St = tok.Literal
		return INT_TYPE
	case token.IF:
		parserVal.St = tok.Literal
		return IF
	case token.ELSE:
		parserVal.St = tok.Literal
		return ELSE
	case token.PROGRAM:
		parserVal.St = tok.Literal
		return PROGRAM
	case token.PRINT:
		parserVal.St = tok.Literal
		return PRINT
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
		return CTE_I
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
		return CTE_STRING
	case token.FLOAT:
		parserVal.St = tok.Literal
		return CTE_F
	default:
		return int(tok.LineNumber)
	}

}
