package elements

import "github.com/jnafolayan/json-parser/tokens"

type Element interface {
	element()
}

type String struct {
	Value tokens.Token
}

func (o *String) element() {}

type Keyword struct {
	Value tokens.Token
}

func (o *Keyword) element() {}

type Object struct {
	Pairs []*ObjectPair
}

type ObjectPair struct {
	Key   tokens.Token
	Value Element
}

func (o *Object) element() {}
