package expr

import (
	"github.com/golu360/lox/parser"
)

type Expr[R any] interface {
	Accept(visitor Visitor[R]) R
}

type Visitor[Expr any] interface {
	VisitBinaryExpr(*Binary)
	VisitGroupingExpr(*Grouping)
	VisitLiteralExpr(*Literal)
	VisitUnaryExpr(*Unary)
}
type Binary struct {
	Left     Expr[Binary]
	Operator parser.Token
	Right    Expr[Binary]
}

func (t *Binary) Accept(v Visitor[Binary]) {
	v.VisitBinaryExpr(t)
}

type Grouping struct {
	Expression Expr[Grouping]
}

func (t *Grouping) Accept(v Visitor[Grouping]) {
	v.VisitGroupingExpr(t)
}

type Literal struct {
	Value interface{}
}

func (t *Literal) Accept(v Visitor[Literal]) {
	v.VisitLiteralExpr(t)
}

type Unary struct {
	Operator parser.Token
	Right    Expr[Unary]
}

func (t *Unary) Accept(v Visitor[Unary]) {
	v.VisitUnaryExpr(t)
}
