package lexer

import (
	"ciri/src/token"
	"testing"
)

func TestTokenizeFloatVar(t *testing.T) {

	input := `
		var x = 10.55
	`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.ID, "x"},
		{token.ASSIGN, "="},
		{token.FLOAT, "10.55"},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong at line=%d . expected=%q, got=%q",
				i, tok.LineNumber, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}

func TestTokenizeIntVar(t *testing.T) {

	input := `
		var x = 10
	`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.ID, "x"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong at line=%d . expected=%q, got=%q",
				i, tok.LineNumber, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}

func TestTokenizePrint(t *testing.T) {
	input := `
		print("hello");
	`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.PRINT, "print"},
		{token.OPEN_PARENTHESIS, "("},
		{token.STRING, `"hello"`},
		{token.CLOSED_PARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong at line=%d . expected=%q, got=%q",
				i, tok.LineNumber, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestTokenizeConditional(t *testing.T) {
	input := `
		if (x > 10) {
			var x2 = x + 10;
			print(10);
		} else if (4 < 10.534) {
			print(x2);
		    print(x);
		}
	`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.OPEN_PARENTHESIS, "("},
		{token.ID, "x"},
		{token.GREATER_THAN, ">"},
		{token.INT, "10"},
		{token.CLOSED_PARENTHESIS, ")"},
		{token.OPEN_BRACE, "{"},
		{token.VAR, "var"},
		{token.ID, "x2"},
		{token.ASSIGN, "="},
		{token.ID, "x"},
		{token.PLUS, "+"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.PRINT, "print"},
		{token.OPEN_PARENTHESIS, "("},
		{token.INT, "10"},
		{token.CLOSED_PARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
		{token.CLOSED_BRACE, "}"},
		{token.ELSE, "else"},
		{token.IF, "if"},
		{token.OPEN_PARENTHESIS, "("},
		{token.INT, "4"},
		{token.LESS_THAN, "<"},
		{token.FLOAT, "10.534"},
		{token.CLOSED_PARENTHESIS, ")"},
		{token.OPEN_BRACE, "{"},
		{token.PRINT, "print"},
		{token.OPEN_PARENTHESIS, "("},
		{token.ID, "x2"},
		{token.CLOSED_PARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
		{token.PRINT, "print"},
		{token.OPEN_PARENTHESIS, "("},
		{token.ID, "x"},
		{token.CLOSED_PARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong at line=%d . expected=%q, got=%q",
				i, tok.LineNumber, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestResetLexerLookup(t *testing.T) {
	input := `var x = 10;`
	l := New(input)
	copy := l.makeCopy()
	l.NextToken()

	l.resetLookup(copy)

	if l.position != 0 {
		t.Fatalf("prev position %d actual: %d  should be %d", copy.position, l.position, copy.position)
	}

}

func TestNextToken(t *testing.T) {
	input := `  
				program myProgram: int {

				var five = 5;
				var ten = 10.2;
				var eleven = 11;
		}
	`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.PROGRAM, "program"},
		{token.ID, "myProgram"},
		{token.COLON, ":"},
		{token.INT_TYPE, "int"},
		{token.OPEN_BRACE, "{"},
		{token.VAR, "var"},
		{token.ID, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.ID, "ten"},
		{token.ASSIGN, "="},
		{token.FLOAT, "10.2"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.ID, "eleven"},
		{token.ASSIGN, "="},
		{token.INT, "11"},
		{token.SEMICOLON, ";"},
		{token.CLOSED_BRACE, "}"},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong at line=%d . expected=%q, got=%q",
				i, tok.LineNumber, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
