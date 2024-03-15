package lexer

import "strings"

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
