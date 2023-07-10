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
		Type:   INT,
		Symbol: string(number),
		File:   lexer.File,
		Line:   lexer.Line,
		Column: lexer.Column,
	}
}

func (lexer *Lexer) lexString() Token {
	var str []byte

	// advance opening "
	lexer.Index++
	lexer.Column++

	for lexer.Index < uint(len(lexer.Src)) && lexer.Src[lexer.Index] != '"' {
		str = append(str, lexer.Src[lexer.Index])

		if lexer.Src[lexer.Index] == '\n' {
			lexer.Line++
			lexer.Column = 0
		}

		lexer.Index++
		lexer.Column++
	}

	if lexer.Src[lexer.Index] != '"' {
		lexer.error("Syntax Error: Expected closing \"")
	}

	// advance closing "
	lexer.Index++
	lexer.Column++

	return Token{
		Type:   STR,
		Symbol: string(str),
		File:   lexer.File,
		Line:   lexer.Line,
		Column: lexer.Column,
	}
}

func (lexer *Lexer) lexChar() Token {
	lexer.Index++
	if len(lexer.Src) == int(lexer.Index) || lexer.Src[lexer.Index] == '\'' {
		lexer.error("Syntax Error: expected char")
	}

	char := lexer.Src[lexer.Index]
	lexer.Index++

	if lexer.Src[lexer.Index] != '\'' {
		lexer.error("Syntax Error: expected closing '")
	}
	lexer.Index++

	return Token{
		Type:   CHAR,
		Symbol: string(char),
		File:   lexer.File,
		Line:   lexer.Line,
		Column: lexer.Column,
	}
}
