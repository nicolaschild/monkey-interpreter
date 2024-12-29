package token

// Our token types
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// Possible tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// Takes an identifier and returns its token
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	// If we match our map for our special identifiers, return
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	// Otherwise, return it as a regular identifier (variable, function names)
	return IDENT
}
