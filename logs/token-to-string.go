package logs

import "spl/lexer"

func TokenToString(token lexer.Token) string {
	if token.Type == lexer.END_OF_LINE {
		return "\\n"
	}

	if token.Type == lexer.EOF {
		return "EOF"
	}

	return token.Symbol
}
