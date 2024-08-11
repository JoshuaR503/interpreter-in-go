package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/token"
)

// Parser represents the parser with fields to manage the current and next tokens,
// as well as a reference to the lexer that provides the tokens.
type Parser struct {
	l         *lexer.Lexer // l is a pointer to an instance of the lexer, used to get the next token.
	curToken  token.Token  // curToken is the current token under examination.
	peekToken token.Token  // peekToken is the next token, used to help decide what to do after curToken.
}

// New creates and returns a new instance of Parser.
// It initializes the parser by reading two tokens, setting both curToken and peekToken.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// Read two tokens so that curToken and peekToken are both set before parsing begins.
	p.nextToken()
	p.nextToken()
	return p
}

// nextToken is a helper method that advances the parser's current and next tokens.
// It moves peekToken to curToken and fetches the next token from the lexer to update peekToken.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram is the entry point for parsing a Monkey program.
// For now, it's not implemented and simply returns nil.
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
