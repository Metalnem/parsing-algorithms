package test

import (
	"fmt"
	"math"
	"testing"

	"github.com/metalnem/parsing-algorithms/parse"
	"github.com/metalnem/parsing-algorithms/parse/climbing"
	"github.com/metalnem/parsing-algorithms/parse/shunting"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input string
		value float64
	}{
		{"1 + 2 * 3", 7},
		{"7 - 9 * (2 - 3)", 16},
		{"2 * 3 * 4", 24},
		{"2 ^ 3 ^ 4", math.Pow(2, math.Pow(3, 4))},
		{"(2 ^ 3) ^ 4", 4096},
		{"5", 5},
		{"4 + 2", 6},
		{"9 - 8 - 7", -6},
		{"9 - (8 - 7)", 8},
		{"(9 - 8) - 7", -6},
		{"2 + 3 ^ 2 * 3 + 4", 33},
		{"-11", -11},
		{"+12", 12},
		{"-(3 * -5)", 15},
		{"+(-4 * +6)", -24},
		{"-3 ^ -4", -0.01234567901},
		{"50 - -100 - -10", 160},
	}

	parsers := []struct {
		name string
		p    parse.Parser
	}{
		{"Precedence climbing", climbing.New()},
		{"Shunting Yard", shunting.New()},
	}

	tolerance := 0.00000000001

	for _, parser := range parsers {
		for _, test := range tests {
			p := parser.p
			test := test

			t.Run(fmt.Sprintf("%s - %s", parser.name, test.input), func(t *testing.T) {
				t.Parallel()

				expr, err := p.Parse(test.input)

				if err != nil {
					t.Fatal(err)
				}

				expected := test.value
				actual := expr.Eval()

				if math.Abs(expected-actual) > tolerance {
					t.Errorf("Expected %f, got %f", expected, actual)
				}
			})
		}
	}
}

func TestParseFail(t *testing.T) {
	tests := []string{
		"1 + 2 abc",
		"3 * 7 ,",
		"((27 - 11) + 3",
		"2 + x",
		"",
		"1+2(3*4)",
	}

	parsers := []struct {
		name string
		p    parse.Parser
	}{
		{"Precedence climbing", climbing.New()},
		{"Shunting Yard", shunting.New()},
	}

	for _, parser := range parsers {
		for _, test := range tests {
			p := parser.p
			test := test

			t.Run(fmt.Sprintf("%s - %s", parser.name, test), func(t *testing.T) {
				t.Parallel()

				_, err := p.Parse(test)

				if err == nil {
					t.Fatal("Expected error")
				}
			})
		}
	}
}
