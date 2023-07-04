package lexer

var symbolsTokens = []StaticToken{
	{Type: PLUS, Symbol: "+"},
	{Type: MINUS, Symbol: "-"},
	{Type: MULTIPLY, Symbol: "*"},
	{Type: DIVIDE, Symbol: "/"},
	{Type: MODULO, Symbol: "%"},

	{Type: OPEN_PAREN, Symbol: "("},
	{Type: CLOSE_PAREN, Symbol: ")"},

	{Type: SEMI_COLON, Symbol: ";"},
	{Type: END_OF_LINE, Symbol: "\n"},
}

var keywordsTokens = []StaticToken{
	{Type: PRINT, Symbol: "print"},
}
