package lexer

import (
	"Interpreter/token"
)

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.eatWhiteSpace()

	switch l.ch {
	case '=':
		tok = l.make2CharToken(newToken(token.ASSIGN, l.ch))
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
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = l.make2CharToken(newToken(token.BANG, l.ch))
	case '/':
		tok = newToken(token.FSLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '?':
		tok = newToken(token.TERNARY, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isIDLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {

	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}

}

type Lexer struct {
	input        string
	position     int  // Points to current char
	readPosition int  // Points to next char
	ch           byte // Current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // Sets EOF at EOF
	} else {
		l.ch = l.input[l.readPosition] // Sets next char if not at EOF
	}
	l.position = l.readPosition // Increments position pointers
	l.readPosition++
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isIDLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) eatWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isIDLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' ||
		ch >= 'A' && ch <= 'Z' ||
		ch == '_' // Treated as letter for identifier purposes
}
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) make2CharToken(defToken token.Token) token.Token {
	t := defToken
	if l.peekChar() == '=' {
		ch := l.ch
		l.readChar()
		if ch == '=' {
			t = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else if ch == '!' {
			t = token.Token{Type: token.NEQ, Literal: string(ch) + string(l.ch)}
		}
	}
	return t
}
