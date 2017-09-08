// Package precedence implements top down operator precedence algorithm for parsing mathematical expressions in infix notation.
package precedence

import (
	"github.com/metalnem/parsing-algorithms/ast"
	"github.com/metalnem/parsing-algorithms/parse"
)

type parser struct {
}

// New creates a new top down operator precedence parser.
func New() parse.Parser {
	return parser{}
}

func (parser) Parse(input string) (ast.Expr, error) {
	return nil, nil
}
