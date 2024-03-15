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

func TestParseJSONObject1(t *testing.T) {
	input := `{
		"key1": true,
		"key2": false,
		"key3": null,
		"key4": "value",
		"key5": 101,
		"key6": [
			1, 2, 4,
			{
				"a": 20,
				"b": 2.3e5
			},
			["40"]
		]
	}`
	l := lexer.FromString(input)
	p := NewParser(l)
	if p.Parse() != nil {
		t.Fatalf("expected json to be valid")
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
