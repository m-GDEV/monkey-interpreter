// First code written for this project

package token

type TokenType string

// Struct defining what information each token needs
type Token struct {
	Type    TokenType // int, str, keyword, etc
	Literal string    // 'x', '10', etc
}

// Multiline const statement??
const (
	ILLEGAL = "ILLEGAL" // to represent token/character we don't recognize
	EOF     = "EOF"

	// Identifier + literals
	IDENT = "IDENT" // add, x, y, etc
	INT   = "INT"   // 1, 45, etc

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
