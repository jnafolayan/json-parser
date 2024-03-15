package lexer

import (
	"bufio"
	"bytes"
	"io"

	"github.com/jnafolayan/json-parser/tokens"
)

type Lexer struct {
	cursor     int
	char       byte
	scanner    *bufio.Scanner
	lineBuffer string
}

func NewLexer(r io.Reader) *Lexer {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	l := &Lexer{
		cursor:  0,
		scanner: s,
	}

	// Read the first character
	l.readCharacter()

	return l
}

// FromString creates a lexer object from an input string
func FromString(input string) *Lexer {
	r := bytes.NewReader([]byte(input))
	return NewLexer(r)
}

func (l *Lexer) NextToken() tokens.Token {
	l.skipWhitespace()

	tok := l.scanNextToken()
	return tok
}

func (l *Lexer) scanNextToken() tokens.Token {
	var tok tokens.Token

	switch l.char {
	case '{':
		tok = tokens.NewToken(tokens.LBRACE, string(l.char))
	case '}':
		tok = tokens.NewToken(tokens.RBRACE, string(l.char))
	case ':':
		tok = tokens.NewToken(tokens.COLON, string(l.char))
	case ',':
		tok = tokens.NewToken(tokens.COMMA, string(l.char))
	case 0:
		tok = tokens.NewToken(tokens.EOF, "")
	default:
		if l.char == '"' {
			// Scan a string
			tok = tokens.NewToken(tokens.STRING, l.scanString())
			break
		}
		tok = tokens.NewToken(tokens.ILLEGAL, string(l.char))
	}

	l.readCharacter()

	return tok
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.char) {
		l.readCharacter()
	}
}

func isWhitespace(b byte) bool {
	// https://www.json.org/json-en.html
	return b == '\u0020' || b == '\u000a' || b == '\u000d' || b == '\u0009'
}

// scanNextLine scans the next line from the source. Returns false if an error occured,
// including if we're at EOF
func (l *Lexer) scanNextLine() bool {
	if l.scanner.Scan() {
		l.lineBuffer = l.scanner.Text()
		// Reset the cursor
		l.cursor = 0
		return true
	}
	return false
}

func (l *Lexer) shouldScanNextLine() bool {
	return len(l.lineBuffer) == 0 || l.cursor >= len(l.lineBuffer)
}

// readCharacter is a helper function to scan the next source line if necessary and
// return the next character in the stream.
func (l *Lexer) readCharacter() byte {
	if l.shouldScanNextLine() {
		couldScan := l.scanNextLine()
		if !couldScan {
			l.char = 0
			return l.char
		}
	}

	l.char = l.lineBuffer[l.cursor]
	l.cursor += 1
	return l.char
}
