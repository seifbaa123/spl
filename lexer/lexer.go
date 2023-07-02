package lexer

import (
	"fmt"
	"os"
	"spl/utils"
)

func NewLexer(src string, file string) *Lexer {
	return &Lexer{
		Src:    src,
		File:   file,
		Line:   0,
		Column: 0,
		Index:  0,
	}
}

func (lexer *Lexer) Tokenize() []Token {
	var tokens []Token

loop:
	for lexer.Index < uint(len(lexer.Src)) {
		// clean white space
		if isWhiteSpace(lexer.Src[lexer.Index]) {
			lexer.Index++
			continue
		}

		// lex symbols
		for _, token := range symbolsTokens {
			if utils.Substr(lexer.Src, int(lexer.Index), len(token.Symbol)) == token.Symbol {
				tokens = append(tokens, Token{
					Type:   token.Type,
					Symbol: token.Symbol,
					File:   lexer.File,
					Line:   lexer.Line,
					Column: lexer.Column,
				})
				lexer.Index++
				lexer.Column++
				continue loop
			}
		}

		// lex identifiers
		if isAlpha(lexer.Src[lexer.Index]) {
			tokens = append(tokens, lexer.lexIdentifier())
			continue
		}

		// lex numbers
		if isNumber(lexer.Src[lexer.Index]) {
			tokens = append(tokens, lexer.lexNumber())
			continue
		}

		// invalid token
		fmt.Fprintf(os.Stderr, "Invalid token %c\n", lexer.Src[lexer.Index])
		os.Exit(1)
	}

	// add eof token
	tokens = append(tokens, Token{
		Type:   EOF,
		Symbol: "EOF",
		File:   lexer.File,
		Line:   lexer.Line,
		Column: lexer.Column,
	})

	return tokens
}
