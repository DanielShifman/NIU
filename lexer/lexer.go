package lexer

import (
	"Interpreter/token"
)

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isIDLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {

	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}

}

type Lexer struct {
	input		 string
	position 	 int  // Points to current char
	readPosition int  // Points to next char
	ch			 byte // Current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0							// Sets EOF at EOF
	} else {
		l.ch = l.input[l.readPosition]      // Sets next char if not at EOF
	}
	l.position = l.readPosition				// Increments position pointers
	l.readPosition++
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isIDLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isIDLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' ||
			ch <= 'A' && ch <= 'Z' ||
			ch == '_' // Treated as letter for identifier purposes
}