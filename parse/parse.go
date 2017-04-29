// Package parse implements various algorithms for parsing mathematical expressions in infix notation.
package parse

import (
	"strconv"

	"github.com/metalnem/parsing-algorithms/ast"
	"github.com/metalnem/parsing-algorithms/scan"
	"github.com/pkg/errors"
)

type assoc int

const (
	left assoc = iota
	right
)

type opInfo struct {
	prec  int
	assoc assoc
}

var binOps = map[string]opInfo{
	"+": {1, left},
	"-": {1, left},
	"*": {2, left},
	"/": {2, left},
	"^": {4, right},
}

var unOps = map[string]opInfo{
	"+": {3, right},
	"-": {3, right},
}

// Parser creates abstract syntax tree of input expression.
type Parser interface {
	Parse(input string) (ast.Expr, error)
}

type parser struct {
}

type state struct {
	s *scan.Scanner
	t scan.Token
}

// NewParser creates a new precedence climbing parser.
func NewParser() Parser {
	return parser{}
}

func (parser) Parse(input string) (ast.Expr, error) {
	s := scan.NewScanner(input)
	t := s.Next()

	state := &state{s: s, t: t}
	expr, err := state.parseExpr(1)

	if err != nil {
		return nil, err
	}

	if state.t.Type != scan.EOF {
		return nil, errors.Errorf("Expected operator, got %s", state.t.Value)
	}

	return expr, err
}

func (s *state) parseExpr(prec int) (ast.Expr, error) {
	lhs, err := s.parsePrimary()

	if err != nil {
		return nil, err
	}

	for {
		if s.t.Type == scan.EOF || s.t.Type != scan.Operator {
			break
		}

		val := s.t.Value
		op, ok := binOps[val]

		if !ok {
			return nil, errors.Errorf("Expected operator, got %s", s.t.Value)
		}

		if binOps[s.t.Value].prec < prec {
			break
		}

		nextPrec := op.prec

		if op.assoc == left {
			nextPrec++
		}

		s.t = s.s.Next()
		rhs, err := s.parseExpr(nextPrec)

		if err != nil {
			return nil, err
		}

		lhs = &ast.BinaryExpr{Op: val, X: lhs, Y: rhs}
	}

	return lhs, nil
}

func (s *state) parsePrimary() (ast.Expr, error) {
	if s.t.Type == scan.LeftParen {
		s.t = s.s.Next()
		expr, err := s.parseExpr(1)

		if err != nil {
			return nil, err
		}

		if s.t.Type != scan.RightParen {
			return nil, errors.Errorf("Expected right paren, got %s", s.t.Value)
		}

		s.t = s.s.Next()
		return expr, nil
	}

	if op, ok := unOps[s.t.Value]; ok {
		val := s.t.Value
		s.t = s.s.Next()
		expr, err := s.parseExpr(op.prec)

		if err != nil {
			return nil, err
		}

		return &ast.UnaryExpr{Op: val, X: expr}, nil
	}

	val, err := strconv.ParseFloat(s.t.Value, 64)

	if err != nil {
		return nil, errors.Errorf("Expected number, got %s", s.t.Value)
	}

	s.t = s.s.Next()
	return &ast.Number{Value: val}, nil
}
