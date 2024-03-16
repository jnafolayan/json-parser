package validations

import (
	"bytes"
	"fmt"
	"unicode"
)

func ValidateString(str string) error {
	validEscapeSequences := []byte{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'}

	for i := 0; i < len(str); i++ {
		if str[i] == '\\' {
			i++
			if bytes.IndexByte(validEscapeSequences, str[i]) == -1 {
				return fmt.Errorf("illegal escape sequence: \\%s", string(str[i]))
			}
			continue
		}

		if unicode.IsControl(rune(str[i])) {
			return fmt.Errorf("unescaped control character in string: %s", str)
		}
	}

	return nil
}
