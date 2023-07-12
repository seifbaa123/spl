package lexer

var symbolsTokens = []StaticToken{
	{Type: PLUS, Symbol: "+"},
	{Type: MINUS, Symbol: "-"},
	{Type: MULTIPLY, Symbol: "*"},
	{Type: DIVIDE, Symbol: "/"},
	{Type: MODULO, Symbol: "%"},

	{Type: OPEN_PAREN, Symbol: "("},
	{Type: CLOSE_PAREN, Symbol: ")"},

	{Type: EQUALS, Symbol: "="},

	{Type: DOT, Symbol: "."},
	{Type: COLON, Symbol: ":"},
	{Type: SEMI_COLON, Symbol: ";"},

	{Type: END_OF_LINE, Symbol: "\n"},
}

var keywordsTokens = []StaticToken{
	{Type: PRINT, Symbol: "print"},

	{Type: LET, Symbol: "let"},
	{Type: CONST, Symbol: "const"},

	{Type: AS, Symbol: "as"},
	{Type: IS, Symbol: "is"},

	{Type: OR, Symbol: "or"},
	{Type: AND, Symbol: "and"},
	{Type: XOR, Symbol: "xor"},
	{Type: NOT, Symbol: "not"},

	{Type: TRUE, Symbol: "true"},
	{Type: FALSE, Symbol: "false"},
}
