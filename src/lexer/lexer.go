package lexer

import (
	"ciri/src/token"
)

type Lexer struct {
	input        string
	position     int
	nextPosition int
	current      byte
	lineNumber   uint32
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input, lineNumber: 0}
	lexer.readChar()
	return lexer
}

func (l *Lexer) ignoreWhitespaces() {
	for l.current == ' ' || l.current == '\t' || l.current == '\n' || l.current == '\r' {
		if l.current == '\n' {
			l.lineNumber += 1
		}
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.current = 0
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
	case '*':
		t = l.newToken(token.MULTIPLY)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.current) {
			return l.lookupKeyword()
		} else if isDigit(l.current) {
			return l.lookupNumerics()
		} else {
			return l.newToken(token.ILLEGAL)
		}
	}
	l.readChar()
	return t
}

func (l *Lexer) lookupKeyword() token.Token {
	copy := l.makeCopy()

	potentialKeyword := getAllLetters(l)
	tokenType := token.LookupSimpleKeyword(potentialKeyword)
	if tokenType != token.ILLEGAL {
		return l.newKeywordToken(tokenType, potentialKeyword)
	}

	l.resetLookup(copy)
	potentialId := getAllChars(l)
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

func (l *Lexer) lookupNumerics() token.Token {
	copy := l.makeCopy()
	potentialFloat := getAllFloats(l)

	tokenType := token.LookupIdentifier(token.FLOAT_IDENT, potentialFloat)
	if tokenType != token.ILLEGAL {
		return l.newKeywordToken(tokenType, potentialFloat)
	}

	l.resetLookup(copy)
	potentialInt := getAllNumbers(l)
	tokenType = token.LookupIdentifier(token.INT_IDENT, potentialInt)
	if tokenType != token.ILLEGAL {
		return l.newKeywordToken(tokenType, potentialInt)
	}

	return l.newToken(token.ILLEGAL)
}

func getAllFloats(l *Lexer) string {
	keyword := ""

	for l.current == '.' || isDigit(l.current) {
		keyword += string(l.current)
		l.readChar()
	}
	return keyword
}

func getAllNumbers(l *Lexer) string {
	keyword := ""

	for isDigit(l.current) {
		keyword += string(l.current)
		l.readChar()
	}
	return keyword
}
func getAllLetters(l *Lexer) string {
	keyword := ""

	for isLetter(l.current) {
		keyword += string(l.current)
		l.readChar()
	}
	return keyword
}

func getAllChars(l *Lexer) string {
	keyword := ""

	for isLetter(l.current) || isDigit(l.current) {
		keyword += string(l.current)
		l.readChar()
	}
	return keyword
}

func (l *Lexer) newKeywordToken(tokenType token.Type, keyword string) token.Token {
	return token.Token{
		Type:            tokenType,
		LineNumber:      l.lineNumber,
		CharacterNumber: uint32(l.position - len(keyword) - 1),
		Literal:         keyword,
	}
}

func (l *Lexer) newToken(tokeType token.Type) token.Token {
	return token.Token{
		Type:            tokeType,
		Literal:         string(l.current),
		LineNumber:      l.lineNumber,
		CharacterNumber: uint32(l.position),
	}
}

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.current) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.current) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.current == '"' || l.current == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
