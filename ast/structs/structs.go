package structs

import (
	"log"

	"github.com/golu360/lox/parser"
)

type Expr interface {
}

type Binary struct {
	left     Expr
	operator parser.Token
	right    Expr
}

func (b *Binary) GetExpr() {
	log.Printf("%s %s %s\n", b.left, b.operator.Lexeme, b.right)
}
func NewBinaryExpr(left Expr, operator parser.Token, right Expr) *Binary {
	return &Binary{
		left:     left,
		right:    right,
		operator: operator,
	}
}
