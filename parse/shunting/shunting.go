// Package shunting implements Shunting Yard algorithm for parsing mathematical expressions in infix notation.
package shunting

import (
	"github.com/metalnem/parsing-algorithms/ast"
	"github.com/metalnem/parsing-algorithms/parse"
	"github.com/metalnem/parsing-algorithms/scan"
	"github.com/pkg/errors"
)

type assoc int

const (
	left assoc = iota
	right
)

type kind int

const (
	unary kind = iota
	binary
)

type op struct {
	value string
	prec  int
	assoc assoc
	kind  kind
}

var ops = []op{
	{"+", 1, left, binary},
	{"-", 1, left, binary},
	{"*", 2, left, binary},
	{"/", 2, left, binary},
	{"+", 3, right, unary},
	{"-", 3, right, unary},
	{"^", 4, right, binary},
}

type parser struct {
}

type state struct {
	s     *scan.Scanner
	ops   []op
	exprs []ast.Expr
}

// New creates a new Shunting Yard parser.
func New() parse.Parser {
	return parser{}
}

func (parser) Parse(input string) (ast.Expr, error) {
	s := scan.NewScanner(input)

	state := &state{s: s}
	expr, err := state.parseExpr()

	if err != nil {
		return nil, err
	}

	if next := state.s.Next(); next.Type != scan.EOF {
		return nil, errors.Errorf("Expected EOF, got %s", next.Value)
	}

	return expr, nil
}

func (s *state) parseExpr() (ast.Expr, error) {
	return nil, nil
}
