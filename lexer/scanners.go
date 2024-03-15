package lexer

import (
	"errors"
	"strings"
)

func (l *Lexer) scanString() string {
	// Move l.char to the position after the opening `"`
	l.readCharacter()

	// Iterate until we hit a closing quote or EOF
	var res strings.Builder
	for l.char != '"' && l.char != 0 {
		res.WriteByte(l.char)
		l.readCharacter()
	}

	return res.String()
}

func (l *Lexer) scanKeyword() string {
	var res strings.Builder
	for isLetter(l.char) {
		res.WriteByte(l.char)
		l.readCharacter()
	}

	return res.String()
}

// scanNumber parses grammar:
// number := integer fraction exponent
// integer := digit | onenine digits | '-' digit | '-' onenine digits
func (l *Lexer) scanNumber() (string, error) {
	var number strings.Builder
	l.scanDigitsIntoBuilder(&number)

	if l.char == '.' {
		number.WriteByte(l.char)
		l.readCharacter()
		if l.scanDigitsIntoBuilder(&number) == 0 {
			return number.String(), errors.New("expected a digit after")
		}
	}

	if l.char == 'E' || l.char == 'e' {
		number.WriteByte('e')
		l.readCharacter()
		if l.char == '+' || l.char == '-' {
			number.WriteByte(l.char)
			l.readCharacter()
		}
		if l.scanDigitsIntoBuilder(&number) == 0 {
			return number.String(), errors.New("expected a digit after")
		}
	}

	num := number.String()
	if len(num) > 1 {
		// Is it a valid number?
		if num[0] == '0' {
			return num, errors.New("invalid number")
		}
	}

	return num, nil
}

func (l *Lexer) scanDigitsIntoBuilder(builder *strings.Builder) (count int) {
	for isDigit(l.char) {
		builder.WriteByte(l.char)
		count += 1
		l.readCharacter()
	}
	return
}
