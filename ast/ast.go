package ast

import (
	"kdnakt/writing-an-interpreter-in-go/token"
)

// Node is an AST node
type Node interface {
	TokenLiteral() string
}

// Statement is a node
type Statement interface {
	Node
	statementNode()
}

// Expression is a node
type Expression interface {
	Node
	expressionNode()
}

// Program is a root node that has statements
type Program struct {
	Statements []Statement
}

// TokenLiteral returns literal value
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement is `let ...`
type LetStatement struct {
	Token token.Token // token.LET
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
// TokenLiteral returns literal value
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// ReturnStatement is `return ...`
type ReturnStatement struct {
	Token		token.Token
	ReturnValue	Expression
}

func (rs *ReturnStatement) statementNode() {}
// TokenLiteral returns literal value
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// Identifier is a symbol
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
// TokenLiteral returns literal value
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}