package ast

import "strings"

type TokenType int

const (
	TokenKey TokenType = iota
	TokenPlus
	TokenPipe
	TokenSequence
	TokenLParen
	TokenRParen
	TokenQuestion
	TokenColon
	TokenEOF
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input string
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) AllTokens() []Token {
	var tokens []Token
	for {
		tok := l.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == TokenEOF {
			break
		}
	}
	return tokens
}

func (l *Lexer) NextToken() Token {
	for l.pos < len(l.input) {
		ch := l.input[l.pos]
		l.pos++

		switch ch {
		case '+':
			return Token{Type: TokenPlus, Value: "+"}
		case '|':
			return Token{Type: TokenPipe, Value: "|"}
		case '(':
			return Token{Type: TokenLParen, Value: "("}
		case ')':
			return Token{Type: TokenRParen, Value: ")"}
		case '?':
			return Token{Type: TokenQuestion, Value: "?"}
		case ':':
			return Token{Type: TokenColon, Value: ":"}
		case ' ':
			return Token{Type: TokenSequence, Value: " "}
		default:
			start := l.pos - 1
			for l.pos < len(l.input) && isAlphaNum(l.input[l.pos]) {
				l.pos++
			}
			value := l.input[start:l.pos]
			return Token{Type: TokenKey, Value: value}
		}
	}
	return Token{Type: TokenEOF}
}

func isAlphaNum(ch byte) bool {
	return strings.ContainsRune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_", rune(ch))
}
