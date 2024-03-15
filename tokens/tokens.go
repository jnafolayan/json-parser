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

	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"
	DQUOTE   = "\""
	SQUOTE   = "'"
	COLON    = ":"
	COMMA    = ","

	NUMBER = "NUMBER"
	STRING = "STRING"

	TRUE  = "TRUE"
	FALSE = "FALSE"
	NULL  = "NULL"
)

var keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
	"null":  NULL,
}

func LookupKeyword(keyword string) TokenType {
	if tt, ok := keywords[keyword]; ok {
		return tt
	}
	return ILLEGAL
}
