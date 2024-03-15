package parser

import (
	"testing"

	"github.com/jnafolayan/json-parser/lexer"
)

func TestParseShouldPass(t *testing.T) {
	input := "{}"
	l := lexer.FromString(input)
	p := NewParser(l)
	if p.Parse() != nil {
		t.Fatalf("expected parsing to pass")
	}
}

func TestParseShouldFail(t *testing.T) {
	input := "{"
	l := lexer.FromString(input)
	p := NewParser(l)
	if p.Parse() == nil {
		t.Fatalf("expected parsing to fail")
	}
}
