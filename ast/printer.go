package ast

import (
	"fmt"

	expr "github.com/golu360/lox/out"
)

type AstPrinter struct {

}

func (printer *AstPrinter) Print(expr.Expr){


}

// VisitBinaryExpr implements Visitor.
func (e AstPrinter) VisitBinaryExpr(*expr.Binary) {
	panic("unimplemented")
}

// VisitGroupingExpr implements Visitor.
func (e AstPrinter) VisitGroupingExpr(*expr.Grouping) {
	panic("unimplemented")
}

// VisitLiteralExpr implements Visitor.
func (e AstPrinter) VisitLiteralExpr(*expr.Literal) {
	panic("unimplemented")
}

// VisitUnaryExpr implements Visitor.
func (e AstPrinter) VisitUnaryExpr(*expr.Unary) {
	panic("unimplemented")
}


