package parser

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jnafolayan/json-parser/elements"
	"github.com/jnafolayan/json-parser/lexer"
	"github.com/jnafolayan/json-parser/tokens"
	"github.com/jnafolayan/json-parser/validations"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken tokens.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}
	p.nextToken()
	return p
}

// Parse returns `true` is the JSON is valid and `false` otherwise.
// Parses the grammar found at https://www.json.org/json-en.html
func (p *Parser) Parse() error {
	_, err := p.parseElement()
	p.nextToken()
	if err == nil && p.currentToken.Type != tokens.EOF {
		return errors.New("dangling elements found")
	}
	return err
}

// RestrictedParse is similar to Parse, except a valid JSON text is
// a serialized object or array.
// https://www.ietf.org/rfc/rfc4627.txt
func (p *Parser) RestrictedParse() error {
	ele, err := p.parseElement()
	if err != nil {
		return err
	}

	p.nextToken()
	if p.currentToken.Type != tokens.EOF {
		return errors.New("dangling elements found")
	}

	if ele.ElementType() != elements.OBJECT && ele.ElementType() != elements.ARRAY {
		return errors.New("JSON text must be a serialized object or array")
	}

	return nil
}

func (p *Parser) nextToken() tokens.Token {
	p.currentToken = p.lexer.NextToken()
	return p.currentToken
}

func (p *Parser) parseElement() (elements.Element, error) {
	switch p.currentToken.Type {
	case tokens.LBRACE:
		return p.parseObject()
	case tokens.LBRACKET:
		return p.parseArray()
	case tokens.STRING:
		return p.parseString()
	case tokens.NUMBER:
		return p.parseNumber()
	case tokens.TRUE, tokens.FALSE, tokens.NULL:
		return p.parseKeyword()
	default:
		return nil, fmt.Errorf("illegal token %q", p.currentToken.Literal)
	}
}

func (p *Parser) parseObject() (*elements.Object, error) {
	obj := &elements.Object{}

	p.nextToken()
	lastTokenWasComma := false
	for p.currentToken.Type == tokens.STRING {
		lastTokenWasComma = false

		key := p.currentToken
		p.nextToken()
		if p.currentToken.Type != tokens.COLON {
			return nil, fmt.Errorf("expected %q, got %q", tokens.COLON, p.currentToken.Literal)
		}

		p.nextToken()
		element, err := p.parseElement()
		if err != nil {
			return nil, err
		}

		obj.Pairs = append(obj.Pairs, &elements.ObjectPair{
			Key:   key,
			Value: element,
		})

		p.nextToken()
		if p.currentToken.Type == tokens.COMMA {
			p.nextToken()
			lastTokenWasComma = true
		}
	}

	if lastTokenWasComma {
		return nil, errors.New("no trailing comma allowed")
	}

	if p.currentToken.Type != tokens.RBRACE {
		return nil, fmt.Errorf("expected %q, found %q", tokens.RBRACE, p.currentToken.Literal)
	}

	return obj, nil
}

func (p *Parser) parseArray() (*elements.Array, error) {
	arr := &elements.Array{}

	p.nextToken()
	lastTokenWasComma := false
	for p.currentToken.Type != tokens.RBRACKET && p.currentToken.Type != tokens.EOF {
		lastTokenWasComma = false

		element, err := p.parseElement()
		if err != nil {
			return nil, err
		}

		arr.Elements = append(arr.Elements, element)
		p.nextToken()

		if p.currentToken.Type == tokens.COMMA {
			p.nextToken()
			lastTokenWasComma = true
		}
	}

	if lastTokenWasComma {
		return nil, errors.New("no trailing comma allowed")
	}

	if p.currentToken.Type != tokens.RBRACKET {
		return nil, fmt.Errorf("expected %q, found %q", tokens.RBRACKET, p.currentToken.Literal)
	}

	return arr, nil
}

func (p *Parser) parseString() (*elements.String, error) {
	err := validations.ValidateString(p.currentToken.Literal)
	if err != nil {
		return nil, err
	}
	return &elements.String{Value: p.currentToken}, nil
}

func (p *Parser) parseNumber() (*elements.Number, error) {
	parts := strings.Split(p.currentToken.Literal, "e")
	i, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid number %s: %w", p.currentToken.Literal, err)
	}

	if len(parts) == 2 {
		i2, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number %s: %w", p.currentToken.Literal, err)
		}

		i = i * math.Pow(10, i2)
	}

	return &elements.Number{Token: p.currentToken, Value: i}, nil
}

func (p *Parser) parseKeyword() (*elements.Keyword, error) {
	return &elements.Keyword{Value: p.currentToken}, nil

}
