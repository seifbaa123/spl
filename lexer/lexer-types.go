package lexer

type TokenType uint

const (
	IDENTIFIER TokenType = iota
	NUMBER

	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	MODULO

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
