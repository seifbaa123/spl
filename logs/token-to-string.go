package logs

import "spl/lexer"

func TokenToString(token lexer.Token) string {
	if token.Type == lexer.STR {
		return "\"" + token.Symbol + "\""
	}

	if token.Type == lexer.NEW_LINE {
		return "\\n"
	}

	if token.Type == lexer.EOF {
		return "EOF"
	}

	return token.Symbol
}
