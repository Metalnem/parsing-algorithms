// Package precedence implements top down operator precedence algorithm for parsing mathematical expressions in infix notation.
package precedence

import (
	"strconv"

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

type symbol struct {
	value string
	lbp   int
	nud   func(*scan.Scanner) (ast.Expr, error)
	led   func(*scan.Scanner, ast.Expr) (ast.Expr, error)
}

type parser struct {
	s *scan.Scanner
}

var symbols map[string]symbol

func init() {
	symbols = byValue([]symbol{
		op("+").infix(10, left).prefix(30),
		op("-").infix(10, left).prefix(30),
		op("*").infix(20, left),
		op("/").infix(20, left),
		op("^").infix(40, right),
	})
}

// New creates a new top down operator precedence parser.
func New() parse.Parser {
	return parser{}
}

func (parser) Parse(input string) (ast.Expr, error) {
	s := scan.NewScanner(input)
	expr, err := expression(s, 0)

	if err != nil {
		return nil, err
	}

	t := s.Next()

	if t.Type != scan.EOF {
		return nil, errors.Errorf("Expected EOF, got %s", t.Value)
	}

	return expr, nil
}

func next(s *scan.Scanner) symbol {
	t := s.Next()

	if t.Type == scan.LeftParen {
		return paren()
	}

	if t.Type == scan.Operator {
		return symbols[t.Value]
	}

	if t.Type == scan.Number {
		return literal(t.Value)
	}

	return symbol{}
}

func expression(s *scan.Scanner, bp int) (ast.Expr, error) {
	t := next(s)
	token := next(s)
	left, err := t.nud(s)

	if err != nil {
		return nil, err
	}

	for bp < token.lbp {
		t = token
		token = next(s)

		if left, err = t.led(s, left); err != nil {
			return nil, err
		}
	}

	return left, nil
}

func op(value string) symbol {
	return symbol{value: value}
}

func paren() symbol {
	var sym symbol

	sym.nud = func(s *scan.Scanner) (ast.Expr, error) {
		expr, err := expression(s, 0)

		if err != nil {
			return nil, err
		}

		t := s.Next()

		if t.Type != scan.RightParen {
			return nil, errors.Errorf("Expected right paren, got %s", t.Value)
		}

		return expr, nil
	}

	return sym
}

func literal(value string) symbol {
	sym := symbol{value: value}

	sym.nud = func(s *scan.Scanner) (ast.Expr, error) {
		val, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return nil, errors.Errorf("Expected number, got %s", value)
		}

		return &ast.Number{Value: val}, nil
	}

	return sym
}

func (sym symbol) infix(bp int, assoc assoc) symbol {
	sym.lbp = bp

	if assoc == right {
		bp = bp - 1
	}

	sym.led = func(s *scan.Scanner, left ast.Expr) (ast.Expr, error) {
		expr, err := expression(s, bp)

		if err != nil {
			return nil, err
		}

		return &ast.BinaryExpr{Op: sym.value, X: left, Y: expr}, nil
	}

	return sym
}

func (sym symbol) prefix(bp int) symbol {
	sym.nud = func(s *scan.Scanner) (ast.Expr, error) {
		expr, err := expression(s, bp)

		if err != nil {
			return nil, err
		}

		return &ast.UnaryExpr{Op: sym.value, X: expr}, nil
	}

	return sym
}

func byValue(symbols []symbol) map[string]symbol {
	m := make(map[string]symbol)

	for _, symbol := range symbols {
		m[symbol.value] = symbol
	}

	return m
}
