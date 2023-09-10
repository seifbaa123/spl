package parser

import (
	"spl/lexer"
)

type Parser struct {
	Tokens []lexer.Token
	Index  uint

	IsInParenthesis bool
}
