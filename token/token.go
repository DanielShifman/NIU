package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // Unknown token
	EOF		= "EOF" // End of File

	//Identifiers and Literals
	IDENT = "IDENT" // add, foo, x, y, ...
	INT	  = "INT"   // 1, 23, 456

	// Operators
	ASSIGN  = "="
	PLUS	= "+"

	// Delimeters
	COMMA	  = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET		 = "LET"
)

var keywords = map[string] TokenType {
	"fn":	FUNCTION,
	"let":	LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}