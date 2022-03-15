package token

import (
	"testing"
)

func TestLookUpSimpleKeywords(t *testing.T) {
	tests := []struct {
		expectedType    Type
		expectedLiteral string
	}{
		{
			expectedType:    PROGRAM,
			expectedLiteral: "program",
		},
		{
			expectedType:    IF,
			expectedLiteral: "if",
		},
		{
			expectedType:    ELSE,
			expectedLiteral: "else",
		},
		{
			expectedType:    PRINT,
			expectedLiteral: "print",
		},
		{
			expectedType:    TRUE,
			expectedLiteral: "true",
		},

		{
			expectedType:    FALSE,
			expectedLiteral: "false",
		},
		{
			expectedType:    INT_TYPE,
			expectedLiteral: "int",
		},
		{
			expectedType:    FLOAT_TYPE,
			expectedLiteral: "float",
		},
		{
			expectedType:    VAR,
			expectedLiteral: "var",
		},
		{
			expectedType:    LESS_THEN_GREAT,
			expectedLiteral: "<>",
		},
	}

	for i, tt := range tests {
		tokenType := LookupSimpleKeyword(tt.expectedLiteral)

		if tokenType != tt.expectedType {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedType, tokenType)
		}

		if tokenType == ILLEGAL {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, ILLEGAL)
		}
	}
}
