package lexer

var symbolsTokens = []StaticToken{
	{Type: PLUS, Symbol: "+"},
	{Type: MINUS, Symbol: "-"},
	{Type: MULTIPLY, Symbol: "*"},
	{Type: DIVIDE, Symbol: "/"},
	{Type: MODULO, Symbol: "%"},

	{Type: END_OF_LINE, Symbol: "\n"},
}
