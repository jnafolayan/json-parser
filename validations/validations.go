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
			if i < len(str) {
				if str[i] == 'u' && isValidHexString(str, i+1, i+5) {
					i += 4
					continue
				}
				if bytes.IndexByte(validEscapeSequences, str[i]) == -1 {
					return fmt.Errorf("illegal escape sequence: \\%s", string(str[i]))
				}
			}
			continue
		}

		if unicode.IsControl(rune(str[i])) {
			return fmt.Errorf("unescaped control character in string: %s", str)
		}
	}

	return nil
}

func isValidHexString(str string, start, end int) bool {
	var ch byte
	for i := start; i < end; i++ {
		ch = str[i]
		if !(ch >= '0' && ch <= '9') && !(ch >= 'a' && ch <= 'f') && !(ch >= 'A' && ch <= 'F') {
			return false
		}
	}
	return true
}
