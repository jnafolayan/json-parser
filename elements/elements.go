package elements

import "github.com/jnafolayan/json-parser/tokens"

const (
	STRING  = "STRING"
	NUMBER  = "NUMBER"
	KEYWORD = "KEYWORD"
	OBJECT  = "OBJECT"
	ARRAY   = "ARRAY"
)

type Element interface {
	ElementType() string
}
type String struct {
	Value tokens.Token
}

func (s *String) ElementType() string {
	return STRING
}

type Number struct {
	Token tokens.Token
	Value float64
}

func (n *Number) ElementType() string {
	return NUMBER
}

type Keyword struct {
	Value tokens.Token
}

func (k *Keyword) ElementType() string {
	return KEYWORD
}

type Object struct {
	Pairs []*ObjectPair
}

func (o *Object) ElementType() string {
	return OBJECT
}

type ObjectPair struct {
	Key   tokens.Token
	Value Element
}

type Array struct {
	Elements []Element
}

func (a *Array) ElementType() string {
	return ARRAY
}
