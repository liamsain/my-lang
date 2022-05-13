package lexer
import (
	"my-lang/token"
)

type Lexer struct {
	input string
	position int //current pos in input (points to current char)
	readPosition int // current reading position in input (after current)
	ch byte // current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case 0:
		tok.Literal = ""
		tok.Type = token.Eof
	default: 
		tok.Type = token.Illegal
		tok.Literal = string(l.ch)
	}
	l.readChar()
	return tok
}

func newToken(tokenType string, ch byte) token.Token{
	return token.Token {Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	// len check on input on every char seems expensive!
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}