package lexer

import (
	"testing"

	"github.com/jnafolayan/json-parser/tokens"
)

func TestNextToken(t *testing.T) {
	input := "{}}"
	expected := []struct {
		expType    tokens.TokenType
		expLiteral string
	}{
		{tokens.LBRACE, "{"},
		{tokens.RBRACE, "}"},
		{tokens.RBRACE, "}"},
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
