package lexer

func isWhiteSpace(c byte) bool {
	return c == ' ' || c == '\t'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'Z' && c <= 'Z') || c == '_'
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func (lexer *Lexer) lexIdentifier() Token {
	var identifier []byte
	for isAlpha(lexer.Src[lexer.Index]) || isNumber(lexer.Src[lexer.Index]) {
		identifier = append(identifier, lexer.Src[lexer.Index])
		lexer.Index++
		lexer.Column++
	}

	return Token{
		Type:   IDENTIFIER,
		Symbol: string(identifier),
		File:   lexer.File,
		Line:   lexer.Line,
		Column: lexer.Column,
	}
}

func (lexer *Lexer) lexNumber() Token {
	var number []byte
	for isNumber(lexer.Src[lexer.Index]) {
		number = append(number, lexer.Src[lexer.Index])
		lexer.Index++
		lexer.Column++
	}

	return Token{
		Type:   NUMBER,
		Symbol: string(number),
		File:   lexer.File,
		Line:   lexer.Line,
		Column: lexer.Column,
	}
}
