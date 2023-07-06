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
		Line:   1,
		Column: 1,
		Index:  0,
	}
}

func (lexer *Lexer) Tokenize() []Token {
	var tokens []Token

loop:
	for lexer.Index < uint(len(lexer.Src)) {
		// handle comments
		if lexer.Src[lexer.Index] == '#' {
			for len(lexer.Src) > int(lexer.Index) && lexer.Src[lexer.Index] != '\n' {
				lexer.Index++
			}

			continue loop
		}

		// clean white space
		if isWhiteSpace(lexer.Src[lexer.Index]) {
			lexer.Index++
			continue loop
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
			token := lexer.lexIdentifier()
			for _, t := range keywordsTokens {
				if t.Symbol == token.Symbol {
					token.Type = t.Type
					tokens = append(tokens, token)
					continue loop
				}
			}

			tokens = append(tokens, token)
			continue loop
		}

		// lex numbers
		if isNumber(lexer.Src[lexer.Index]) {
			tokens = append(tokens, lexer.lexNumber())
			continue loop
		}

		// lex char
		if lexer.Src[lexer.Index] == '\'' {
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

			tokens = append(tokens, Token{
				Type:   CHAR,
				Symbol: string(char),
				File:   lexer.File,
				Line:   lexer.Line,
				Column: lexer.Column,
			})

			continue loop
		}

		// invalid token
		lexer.error(fmt.Sprintf("Syntax Error: invalid token %c", lexer.Src[lexer.Index]))
	}

	// add eof token
	tokens = append(tokens, Token{Type: EOF})

	return tokens
}

func (lexer *Lexer) error(message string) {
	fmt.Fprintf(os.Stderr, "%s:%d:%d: %s\n", lexer.File, lexer.Line, lexer.Column, message)
	os.Exit(1)
}
