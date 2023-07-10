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

	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	MODULO

	OPEN_PAREN
	CLOSE_PAREN

	EQUALS

	COLON
	SEMI_COLON

	END_OF_LINE
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
