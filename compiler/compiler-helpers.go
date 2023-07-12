package compiler

import (
	"fmt"
	"spl/lexer"
)

func prefixToken(s string, t lexer.Token) string {
	return fmt.Sprintf("%s_%d_%d", s, t.Line, t.Column)
}
