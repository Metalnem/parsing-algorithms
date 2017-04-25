// Package ast declares the types used to represent syntax trees for simple expressions.
package ast

import "math"

// Expr represents a node in the syntax tree.
type Expr interface {
	Eval() float64
}

// UnaryExpr represents a unary expression.
type UnaryExpr struct {
	Op string
	X  Expr
}

// BinaryExpr represents a binary expression.
type BinaryExpr struct {
	Op string
	X  Expr
	Y  Expr
}

// Number represents a number.
type Number struct {
	Value float64
}

// Eval calculates the value of a unary expression.
func (expr *UnaryExpr) Eval() float64 {
	switch expr.Op {
	case "-":
		return -expr.X.Eval()
	default:
		panic("Invalid unary operator.")
	}
}

// Eval calculates the value of a binary expression.
func (expr *BinaryExpr) Eval() float64 {
	switch expr.Op {
	case "+":
		return expr.X.Eval() + expr.Y.Eval()
	case "-":
		return expr.X.Eval() - expr.Y.Eval()
	case "*":
		return expr.X.Eval() * expr.Y.Eval()
	case "/":
		return expr.X.Eval() / expr.Y.Eval()
	case "^":
		return math.Pow(expr.X.Eval(), expr.Y.Eval())
	default:
		panic("Invalid binary operator.")
	}
}

// Eval returns the value of a number.
func (expr *Number) Eval() float64 {
	return expr.Value
}
