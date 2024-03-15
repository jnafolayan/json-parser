package lexer

import (
	"bufio"
	"bytes"
	"io"

	"github.com/jnafolayan/json-parser/tokens"
)

type Lexer struct {
	cursor     uint32
	scanner    *bufio.Scanner
	lineBuffer string
}

func NewLexer(r io.Reader) *Lexer {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	return &Lexer{
		cursor:  0,
		scanner: s,
	}
}

// FromString creates a lexer object from an input string
func FromString(input string) *Lexer {
	r := bytes.NewReader([]byte(input))
	return NewLexer(r)
}

func (l *Lexer) NextToken() tokens.Token {
	if len(l.lineBuffer) == 0 {
		couldRead := l.readNextLine()
		if !couldRead {
			return tokens.Token{
				Type:    tokens.EOF,
				Literal: tokens.EOF,
			}
		}
	}

	return tokens.Token{}
}

// readNextLine scans the next line from the source. Returns false if an error occured,
// including if we're at EOF
func (l *Lexer) readNextLine() bool {
	if l.scanner.Scan() {
		l.lineBuffer = l.scanner.Text()
		return true
	}
	return false
}
