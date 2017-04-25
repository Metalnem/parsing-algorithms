package scan

import "testing"

func TestScanner(t *testing.T) {
	s := NewScanner("3.14 + 2^10 - 13 * (27 - 81 * 23 ^ (7-11))/3 - 2.718")

	expected := []Token{
		{Number, "3.14"},
		{Operator, "+"},
		{Number, "2"},
		{Operator, "^"},
		{Number, "10"},
		{Operator, "-"},
		{Number, "13"},
		{Operator, "*"},
		{LeftParen, "("},
		{Number, "27"},
		{Operator, "-"},
		{Number, "81"},
		{Operator, "*"},
		{Number, "23"},
		{Operator, "^"},
		{LeftParen, "("},
		{Number, "7"},
		{Operator, "-"},
		{Number, "11"},
		{RightParen, ")"},
		{RightParen, ")"},
		{Operator, "/"},
		{Number, "3"},
		{Operator, "-"},
		{Number, "2.718"},
	}

	var actual []Token

	for token := s.Next(); token.Type != EOF; token = s.Next() {
		actual = append(actual, token)
	}

	if len(expected) != len(actual) {
		t.Errorf("Expected %d, got %d", len(expected), len(actual))
	}

	for i, token := range expected {
		if token != actual[i] {
			t.Fatalf("Expected %v, got %v", token, actual[i])
		}
	}
}
