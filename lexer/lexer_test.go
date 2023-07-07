package lexer

import (
	"go/token"
	"testing"

	"github.com/crbroughton/monkey-interpreter/token"
)

func TextNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	line := New(input)

	for i, tokenType := range tests {
		tok := line.NextToken()

		if tok.Type != tokenType.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tokenType.expectedType, tok.Type)
		}
		if tok.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tokenType.expectedLiteral, tok.Literal)
		}

	}
}
