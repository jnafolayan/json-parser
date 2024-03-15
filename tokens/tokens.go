package tokens

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	EOF    = "EOF"
	LBRACE = "{"
	RBRACE = "}"
)
