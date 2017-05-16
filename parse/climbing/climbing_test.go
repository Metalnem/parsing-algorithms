package climbing

import (
	"math"
	"testing"
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

	p := New()
	tolerance := 0.00000000001

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
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

func TestParseFail(t *testing.T) {
	tests := []struct {
		input string
		err   string
	}{
		{"1 + 2 abc", "Expected operator, got a"},
		{"3 * 7 ,", "Expected operator, got ,"},
		{"((27 - 11) + 3", "Expected right paren, got EOF"},
		{"2 + x", "Expected number, got x"},
		{"", "Expected number, got EOF"},
		{"1+2(3*4)", "Expected operator, got ("},
	}

	p := New()

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			_, err := p.Parse(test.input)

			if err == nil {
				t.Fatal("Expected error")
			}

			expected := test.err
			actual := err.Error()

			if expected != actual {
				t.Errorf("Expected %s, got %s", expected, actual)
			}
		})
	}
}
