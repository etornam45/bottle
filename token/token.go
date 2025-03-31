package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

 const (
	ILLEGAL = "ILLEGAL"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"
  STR   = "STR"
	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	MOD      = "%"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="
	LT_EQ  = "<="
	GT_EQ  = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	SWITCH   = "SWITCH"
	CASE     = "CASE"
	DEFAULT  = "DEFAULT"

	// END OF FILE
	EOF = "EOF"
)

var keywords = map[string]TokenType{
	"let": LET,
	"fn":  FUNCTION,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
	"switch": SWITCH,
	"case": CASE,
	"default": DEFAULT,
	"true": TRUE,
	"false": FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
