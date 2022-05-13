package lexer

import (
	"my-lang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	// input := `=+(){},;`
	input := `^`;

	tests := []struct {
		expectedType string
		expectedLiteral string
	} {
		{ token.Illegal, "^"},
		{ token.Eof, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q",i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral  {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got =%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
