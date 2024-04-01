package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" //signifies a token or character we don't know about
	EOF = "EOF" // tells the parser that we can stop

	IDENT = "IDENT"
	INT = "INT"

	ASSIGN = "="
	PLUS = "+"

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = "("
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
)

// we also need to map certain keywords to token types
var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok { //return the keyword token type if its in the keyword map
		return tok
	}
	return IDENT
}