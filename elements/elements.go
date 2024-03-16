package elements

import "github.com/jnafolayan/json-parser/tokens"

type Element interface {
	element()
}
type String struct {
	Value tokens.Token
}

func (s *String) element() {}

type Number struct {
	Token tokens.Token
	Value float64
}

func (n *Number) element() {}

type Keyword struct {
	Value tokens.Token
}

func (k *Keyword) element() {}

type Object struct {
	Pairs []*ObjectPair
}

func (o *Object) element() {}

type ObjectPair struct {
	Key   tokens.Token
	Value Element
}

type Array struct {
	Elements []Element
}

func (a *Array) element() {}
