package lexer

import "interpreter/token";

type Lexer struct {
	input string
	position int // current position in input (points to current char)
	readPosition int // current reading position in in∑put (after current char) 
	ch byte // current char under examination
}

func New(input string) *Lexer { 
	l := &Lexer{input: input}
	l.readChar()
	return l
}


// Te purpose of readChar is to give us the next character and advance our position in the input string.
// TODO: Fully support Unicode
func (l *Lexer) readChar() {
	// Check whether we have reached the end of input.
	if l.readPosition >= len(l.input) {
		// sets l.ch to 0, which is the ASCII code for the "NUL" character and signifies either 
		// “we haven’t read anything yet” or “end of file” for us.
		l.ch = 0
	} else {
		// sets l.ch to the next character by accessing l.input[l.readPosition].
		l.ch = l.input[l.readPosition]
	}

	// l.position is updated to the just used l.readPosition and l.readPosition is incremented by one. 
	// l.readPosition always points to the next position
	l.position = l.readPosition
	l.readPosition += 1
}

// Look at the current character under examination (l.ch) and return a token depending on which character it is.
// 
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}