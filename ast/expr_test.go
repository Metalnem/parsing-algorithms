package ast

import (
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	expr := &BinaryExpr{
		"^",
		&BinaryExpr{
			"+",
			&BinaryExpr{
				"+",
				&Number{3.14},
				&BinaryExpr{
					"-",
					&Number{2},
					&Number{1},
				},
			},
			&BinaryExpr{
				"/",
				&UnaryExpr{
					"-",
					&Number{1},
				},
				&Number{2},
			},
		},
		&BinaryExpr{
			"*",
			&Number{2.5},
			&Number{0.1},
		},
	}

	tolerance := 0.00000000001
	expected := 1.38125971592
	actual := expr.Eval()

	if math.Abs(expected-actual) > tolerance {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}
