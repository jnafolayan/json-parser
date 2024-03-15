package tokens

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(t TokenType, literal string) Token {
	return Token{t, literal}
}

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	LBRACE = "{"
	RBRACE = "}"
)
