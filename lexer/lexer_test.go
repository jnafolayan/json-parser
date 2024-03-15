package lexer

import (
	"testing"

	"github.com/jnafolayan/json-parser/tokens"
)

func TestNextToken(t *testing.T) {
	input := `{}}
			"hello"
			true false null
			0 01 10 1.1 1e4 1.2e5 1.2e-5 1.2e`
	expected := []struct {
		expType    tokens.TokenType
		expLiteral string
	}{
		{tokens.LBRACE, "{"},
		{tokens.RBRACE, "}"},
		{tokens.RBRACE, "}"},
		{tokens.STRING, "hello"},
		{tokens.TRUE, "true"},
		{tokens.FALSE, "false"},
		{tokens.NULL, "null"},
		{tokens.NUMBER, "0"},
		{tokens.ILLEGAL, "01"},
		{tokens.NUMBER, "10"},
		{tokens.NUMBER, "1.1"},
		{tokens.NUMBER, "1e4"},
		{tokens.NUMBER, "1.2e5"},
		{tokens.NUMBER, "1.2e-5"},
		{tokens.ILLEGAL, "1.2e"},
		{tokens.EOF, ""},
	}

	l := FromString(input)
	for _, tt := range expected {
		tok := l.NextToken()
		if tok.Type != tt.expType {
			t.Fatalf("expected type %q, got %q", tt.expType, tok.Type)
		}
		if tok.Literal != tt.expLiteral {
			t.Fatalf("expected literal %q, got %q", tt.expLiteral, tok.Literal)
		}
	}
}

func TestEmptySource(t *testing.T) {
	input := ""
	l := FromString(input)
	tok := l.NextToken()
	if tok.Type != tokens.EOF {
		t.Fatalf("expected type %q, got %q", tokens.EOF, tok.Type)
	}
	if tok.Literal != "" {
		t.Fatalf("expected literal %q, got %q", "", tok.Literal)
	}
}
