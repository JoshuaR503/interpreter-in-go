package parser

import (
	"fmt"
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
	errors    []string     // A slice of strings to store any errors encountered during parsing.
}

// New creates and returns a new instance of Parser.
// It initializes the parser by reading two tokens, setting both curToken and peekToken.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{}, // Initialize the errors slice as an empty list.
	}
	// Read two tokens so that curToken and peekToken are both set before parsing begins.
	p.nextToken()
	p.nextToken()
	return p
}

// Errors returns the list of errors that the parser encountered during parsing.
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError adds an error message to the parser's errors slice when the expected
// token type does not match the actual peekToken type.
func (p *Parser) peekError(t token.TokenType) {
	// Create an error message indicating the expected token type and the actual token type.
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)

	// Append the error message to the errors slice.
	p.errors = append(p.errors, msg)
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
	// Create the root node of the AST, which will hold all the statements.
	program := &ast.Program{}

	// Initialize the Statements slice to hold the parsed statements.
	program.Statements = []ast.Statement{}

	// Iterate over the input tokens until we reach the end (EOF).
	for p.curToken.Type != token.EOF {
		// Parse the current statement.
		stmt := p.parseStatement()

		// If parseStatement returns a valid statement (not nil),
		// add it to the Statements slice of the root node.
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		// Move to the next token in the input.
		p.nextToken()
	}

	// Return the fully constructed AST root node.
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	// parseStatement checks the type of the current token to decide what kind of statement to parse.
	// If the current token is a 'LET' token, it delegates to parseLetStatement.
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil // If the token type is not recognized, return nil.
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// Create a new LetStatement node using the current token, which should be a 'LET' token.
	stmt := &ast.LetStatement{Token: p.curToken}

	// Expect the next token to be an identifier (the variable name).
	// If it's not, return nil and skip further parsing for this statement.
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// Create an Identifier node using the current token and set it as the Name of the LetStatement.
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Expect the next token to be an equal sign ('=').
	// If it's not, return nil and skip further parsing for this statement.
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: Currently, we're skipping the expressions that follow the '=' sign
	// until we encounter a semicolon. This will be replaced with actual expression parsing later.
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	// Return the constructed LetStatement node.
	return stmt
}

// curTokenIs checks if the current token matches the given token type.
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs checks if the next token (peekToken) matches the given token type.
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek checks if the next token is of the expected type.
// If it is, it advances to the next token and returns true.
// If not, it returns false, indicating an unexpected token.
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken() // Move to the next token if the type matches.
		return true
	} else {
		p.peekError(t)
		return false // Return false if the expected token type does not match.
	}
}
