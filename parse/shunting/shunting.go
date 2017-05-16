// Package shunting implements Shunting Yard algorithm for parsing mathematical expressions in infix notation.
package shunting

import (
	"github.com/metalnem/parsing-algorithms/ast"
	"github.com/metalnem/parsing-algorithms/parse"
	"github.com/metalnem/parsing-algorithms/scan"
	"github.com/pkg/errors"
)

type parser struct {
}

type state struct {
	s *scan.Scanner
	t scan.Token
}

// New creates a new Shunting Yard parser.
func New() parse.Parser {
	return parser{}
}

func (parser) Parse(input string) (ast.Expr, error) {
	s := scan.NewScanner(input)
	t := s.Next()

	state := &state{s: s, t: t}
	expr, err := state.parseExpr()

	if err != nil {
		return nil, err
	}

	if state.t.Type != scan.EOF {
		return nil, errors.Errorf("Expected EOF, got %s", state.t.Value)
	}

	return expr, nil
}

func (s *state) parseExpr() (ast.Expr, error) {
	return nil, nil
}
