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

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // token.LET
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}