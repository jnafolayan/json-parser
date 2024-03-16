package lexer

import (
	"bufio"
	"bytes"
	"io"

	"github.com/jnafolayan/json-parser/tokens"
)

type Lexer struct {
	// cursor is 1 ahead of the current scanned character
	cursor      int
	char        byte
	scanner     *bufio.Scanner
	chunkBuffer string
	// number of lines to read into the buffer
	linesPerChunk int
}

func NewLexer(r io.Reader) *Lexer {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	l := &Lexer{
		cursor:        0,
		scanner:       s,
		linesPerChunk: 5,
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
	case '[':
		tok = tokens.NewToken(tokens.LBRACKET, string(l.char))
	case ']':
		tok = tokens.NewToken(tokens.RBRACKET, string(l.char))
	case ':':
		tok = tokens.NewToken(tokens.COLON, string(l.char))
	case ',':
		tok = tokens.NewToken(tokens.COMMA, string(l.char))
	case '/':
		tok = tokens.NewToken(tokens.SOLIDUS, string(l.char))
	case '\\':
		tok = tokens.NewToken(tokens.RSOLIDUS, string(l.char))
	case 0:
		tok = tokens.NewToken(tokens.EOF, "")
	case '"':
		// Scan a string
		tok = tokens.NewToken(tokens.STRING, l.scanString())
	default:
		if isLetter(l.char) {
			tok.Literal = l.scanKeyword()
			tok.Type = tokens.LookupKeyword(tok.Literal)
			// Return to prevent a readCharacter() call since scanKeyword() stops
			// at the following character
			return tok
		} else if isDigit(l.char) || (l.char == '-' && isDigit(l.peekCharacter())) {
			uMinus := false
			if l.char == '-' {
				uMinus = true
				l.readCharacter()
			}
			num, err := l.scanNumber()
			if uMinus && num != "" {
				num = "-" + num
			}
			if err != nil {
				tok = tokens.NewToken(tokens.ILLEGAL, num)
			} else {
				tok.Type = tokens.NUMBER
				tok.Literal = num
			}
			// Return to prevent a readCharacter() call since scanNumber() stops
			// at the following character
			return tok
		} else {
			tok = tokens.NewToken(tokens.ILLEGAL, string(l.char))
		}
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

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// scanNextChunk scans the next chunk from the source. Returns false if an error occured,
// including if we're at EOF
func (l *Lexer) scanNextChunk() bool {
	l.chunkBuffer = ""
	firstScan := true
	for l.scanner.Scan() {
		if !firstScan {
			l.chunkBuffer += "\n"
		}
		firstScan = false
		l.chunkBuffer += l.scanner.Text()
	}

	l.cursor = 0
	return l.chunkBuffer != ""
}

func (l *Lexer) shouldScanNextLine() bool {
	return len(l.chunkBuffer) == 0 || l.cursor >= len(l.chunkBuffer)
}

// readCharacter is a helper function to scan the next input chunk if necessary and
// return the next character in the stream.
func (l *Lexer) readCharacter() byte {
	if l.shouldScanNextLine() {
		couldScan := l.scanNextChunk()
		if !couldScan {
			l.char = 0
			return l.char
		}
	}

	l.char = l.chunkBuffer[l.cursor]
	l.cursor += 1
	return l.char
}

func (l *Lexer) peekCharacter() byte {
	if l.shouldScanNextLine() {
		prev := l.chunkBuffer
		cursor := l.cursor
		l.scanNextChunk()
		// Restore any past or unused state
		l.chunkBuffer = prev + l.chunkBuffer
		l.cursor = cursor
	}

	if l.cursor < len(l.chunkBuffer) {
		// Remember cursor is a lookahead
		return l.chunkBuffer[l.cursor]
	}

	return 0
}
