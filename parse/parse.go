// Package parse defines common interface for parsing mathematical expressions in infix notation.
package parse

import "github.com/metalnem/parsing-algorithms/ast"

// Parser is an interface that wraps the basic Parse method.
type Parser interface {
	Parse(input string) (ast.Expr, error)
}
