package lexer

type TokenType uint

const (
	IDENTIFIER TokenType = iota
	INT
	STR
	CHAR

	LET
	CONST
	PRINT

	TRUE
	FALSE

	AS
	IS

	OR
	AND
	XOR
	NOT

	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	MODULO

	OPEN_PAREN
	CLOSE_PAREN

	OPEN_BRACKET
	CLOSE_BRACKET

	EQUALS

	EQUALS_TO
	NOT_EQUALS_TO
	GREATER
	GREATER_OR_EQUALS
	LESS
	LESS_OR_EQUALS

	DOT
	COLON
	COMMA
	SEMI_COLON

	NEW_LINE
	EOF
)

type StaticToken struct {
	Type   TokenType
	Symbol string
}

type Token struct {
	Type   TokenType
	Symbol string
	File   string
	Line   uint
	Column uint
}

type Lexer struct {
	Src    string
	File   string
	Line   uint
	Column uint
	Index  uint
}
