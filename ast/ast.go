package ast

import (
	"bytes"

	"kdnakt/writing-an-interpreter-in-go/token"
)

// Node is an AST node
type Node interface {
	TokenLiteral() string
	String() string
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

// String returns statements as string
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

// String returns the content of let statement
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement is `return ...`
type ReturnStatement struct {
	Token		token.Token
	ReturnValue	Expression
}

func (rs *ReturnStatement) statementNode() {}
// TokenLiteral returns literal value
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
// String returns `return ...`
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

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
func (i *Identifier) String() string {
	return i.Value
}

// ExpressionStatement is an expression
type ExpressionStatement struct {
	Token		token.Token // first token in the expression
	Expression	Expression
}

func (es *ExpressionStatement) statementNode() {}
// TokenLiteral returns literal value
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

type PrefixExpression struct {
	Token		token.Token
	Operator	string
	Right		Expression
}

func (pe *PrefixExpression) expressionNode() {}
// TokenLiteral returns literal value
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}


type InfixExpression struct {
	Token		token.Token
	Left		Expression
	Operator	string
	Right		Expression
}

func (ie *InfixExpression) expressionNode() {}
// TokenLiteral returns literal value
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string { return b.Token.Literal }

type IfExpression struct {
	Token		token.Token
	Condition	Expression
	Consequence	*BlockStatement
	Alternative	*BlockStatement
}

func (ie *IfExpression) expressionNode() {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

type BlockStatement struct {
	Token		token.Token
	Statements	[]Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}