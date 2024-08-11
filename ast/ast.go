package ast

import "go/token"

// Node is the base interface that all nodes in our AST (Abstract Syntax Tree) must implement.
// It requires a TokenLiteral() method that returns the literal value of the token associated with the node.
// This method is primarily used for debugging and testing purposes.
type Node interface {
	TokenLiteral() string
}

// Statement is an interface that extends the Node interface.
// It represents a statement in our AST and includes a dummy method statementNode().
// The purpose of this method is to help the Go compiler distinguish between Statements and Expressions,
// potentially causing it to throw errors if they are used incorrectly.
type Statement interface {
	Node
	statementNode()
}

// Expression is an interface that also extends the Node interface.
// It represents an expression in our AST and includes a dummy method expressionNode().
// Similar to Statement, this method helps the Go compiler differentiate between Expressions and Statements,
// ensuring that they are used correctly in the AST.
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement represents a 'let' statement in our AST (Abstract Syntax Tree).
// It has two fields:
// - Name: This holds the identifier (or name) of the variable being declared.
// - Value: This holds the expression that assigns a value to the variable.
// The two methods statementNode and TokenLiteral satisfy the Statement and Node interfaces, respectively.

type LetStatement struct {
	Token token.Token // The token.LET token, representing the 'let' keyword.
	Name  *Identifier // The identifier (variable name) in the 'let' statement, e.g., 'x' in 'let x = 5;'.
	Value Expression  // The expression that provides the value to be assigned to the identifier.
}

// statementNode is a dummy method that helps the Go compiler recognize this as a Statement node.
func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the literal value of the 'let' token, which is used mainly for debugging.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier represents an identifier (variable name) in our AST.
// It implements the Expression interface, allowing it to be used in different parts of the program,
// even though in the context of a 'let' statement, it doesn't produce a value.
// This simplifies our code by reusing the Identifier type in multiple places.
type Identifier struct {
	Token token.Token // The token.IDENT token, representing the identifier.
	Value string      // The name of the identifier, e.g., 'x' in 'let x = 5;'.
}

// expressionNode is a dummy method that helps the Go compiler recognize this as an Expression node.
func (i *Identifier) expressionNode() {}

// TokenLiteral returns the literal value of the identifier's token, which is used mainly for debugging.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
