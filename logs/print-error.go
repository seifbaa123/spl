package logs

import (
	"fmt"
	"os"
	"spl/lexer"
)

func PrintError(token lexer.Token, message string) {
	fmt.Fprintf(os.Stderr, "%s:%d:%d: %s\n", token.File, token.Line, token.Column, message)
}
