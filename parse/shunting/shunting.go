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
	state := &state{
		s:   scan.NewScanner(input),
		ops: []op{op{}},
	}

	if err := state.parseExpr(); err != nil {
		return nil, err
	}

	if next := state.s.Next(); next.Type != scan.EOF {
		return nil, errors.Errorf("Expected EOF, got %s", next.Value)
	}

	return state.exprs[0], nil
}

func (s *state) parseExpr() error {
	return nil
}

func (s *state) parsePrimary() error {
	return nil
}

func (s *state) push(op op) {
	for greater(s.ops[len(s.ops)-1], op) {
		s.pop()
	}

	s.ops = append(s.ops, op)
}

func (s *state) pop() {
	op := s.ops[len(s.ops)-1]
	s.ops = s.ops[:len(s.ops)-1]

	if op.kind == binary {
		y := s.exprs[len(s.exprs)-1]
		s.exprs = s.exprs[:len(s.exprs)-1]

		x := s.exprs[len(s.exprs)-1]
		s.exprs[len(s.exprs)-1] = &ast.BinaryExpr{Op: op.value, X: x, Y: y}
	} else {
		x := s.exprs[len(s.exprs)-1]
		s.exprs[len(s.exprs)-1] = &ast.UnaryExpr{Op: op.value, X: x}
	}
}

func makeUnary(s string) (op, bool) {
	return makeOp(s, unary)
}

func makeBinary(s string) (op, bool) {
	return makeOp(s, binary)
}

func makeOp(s string, kind kind) (op, bool) {
	for _, op := range ops {
		if op.value == s && op.kind == kind {
			return op, true
		}
	}

	return op{}, false
}

func greater(x, y op) bool {
	if x.kind == binary && y.kind == binary {
		return x.prec > y.prec || (x.assoc == left && x.prec == y.prec)
	}

	if x.kind == unary && y.kind == binary {
		return x.prec >= y.prec
	}

	return false
}
