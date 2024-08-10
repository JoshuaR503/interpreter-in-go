package lexer

import "interpreter/token"

// The Lexer struct holds the input string and information about the current position in that string.
type Lexer struct {
	input        string // the input string that we are going to analyze
	position     int    // current position in the input string (points to the current character)
	readPosition int    // the next reading position in the input string (one character ahead)
	ch           byte   // the current character being analyzed
}

// New creates a new Lexer instance and initializes it with the input string.
func New(input string) *Lexer {
	// Create a new Lexer and set the input string
	l := &Lexer{input: input}
	// Read the first character to initialize the lexer
	l.readChar()
	return l
}

// readChar reads the next character in the input string and advances the lexer’s position.
func (l *Lexer) readChar() {
	// Check if we've reached the end of the input string
	if l.readPosition >= len(l.input) {
		// If yes, set the current character to 0 (NUL), which signifies the end of the input
		l.ch = 0
	} else {
		// Otherwise, set the current character to the next one in the input string
		l.ch = l.input[l.readPosition]
	}
	// Move the current position to the read position
	l.position = l.readPosition
	// Move the read position one step forward to prepare for the next character
	l.readPosition += 1
}

// NextToken identifies and returns the next token (a meaningful element like a word, symbol, or number) from the input string.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Skip any whitespace (like spaces or tabs) so we can focus on meaningful characters
	l.skipWhitespace()

	// Check what the current character is and decide what type of token it represents
	switch l.ch {
	case '=':
		// Handle '==' as a comparison operator
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			// Handle '=' as an assignment operator
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch) // Handle the '+' operator
	case '-':
		tok = newToken(token.MINUS, l.ch) // Handle the '-' operator
	case '!':
		// Handle '!=' as a "not equal" operator
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			// Handle '!' as a "bang" operator (logical NOT)
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch) // Handle the '/' operator
	case '*':
		tok = newToken(token.ASTERISK, l.ch) // Handle the '*' operator
	case '<':
		tok = newToken(token.LT, l.ch) // Handle the '<' (less than) operator
	case '>':
		tok = newToken(token.GT, l.ch) // Handle the '>' (greater than) operator
	case ';':
		tok = newToken(token.SEMICOLON, l.ch) // Handle the ';' (semicolon)
	case ',':
		tok = newToken(token.COMMA, l.ch) // Handle the ',' (comma)
	case '{':
		tok = newToken(token.LBRACE, l.ch) // Handle the '{' (left brace)
	case '}':
		tok = newToken(token.RBRACE, l.ch) // Handle the '}' (right brace)
	case '(':
		tok = newToken(token.LPAREN, l.ch) // Handle the '(' (left parenthesis)
	case ')':
		tok = newToken(token.RPAREN, l.ch) // Handle the ')' (right parenthesis)
	case 0:
		// If we've reached the end of the input, return an EOF (End Of File) token
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// Handle identifiers (like variable names) or numbers
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()          // Read the full identifier
			tok.Type = token.LookupIdent(tok.Literal) // Determine the type of identifier
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT         // Set the type to an integer
			tok.Literal = l.readNumber() // Read the full number
			return tok
		} else {
			// If the character is unrecognized, mark it as an illegal token
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// Move to the next character for further analysis
	l.readChar()
	return tok
}

// readIdentifier reads an identifier (like a variable name) until a non-letter character is found.
func (l *Lexer) readIdentifier() string {
	position := l.position
	// Continue reading characters as long as they are letters
	for isLetter(l.ch) {
		l.readChar()
	}
	// Return the full identifier
	return l.input[position:l.position]
}

// newToken creates a new token with the specified type and character value.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// isLetter checks if a character is a letter (a-z, A-Z, or '_').
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// skipWhitespace skips over any spaces, tabs, or newlines in the input string.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// readNumber reads a number (a sequence of digits) until a non-digit character is found.
func (l *Lexer) readNumber() string {
	position := l.position
	// Continue reading characters as long as they are digits
	for isDigit(l.ch) {
		l.readChar()
	}
	// Return the full number
	return l.input[position:l.position]
}

// isDigit checks if a character is a digit (0-9).
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// peekChar allows us to look at the next character in the input string without moving the current position.
func (l *Lexer) peekChar() byte {
	// Check if we've reached the end of the input string
	if l.readPosition >= len(l.input) {
		return 0 // Return 0 (NUL) if at the end
	} else {
		// Return the next character in the input string
		return l.input[l.readPosition]
	}
}
