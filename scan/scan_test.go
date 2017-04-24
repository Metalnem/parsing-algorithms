package scan

import "testing"

func TestScanner(t *testing.T) {
	s := New("3.14 + 2^10 - 13 * (27 - 81 * 23 ^ (7-11))/3 - 2.718")

	expected := []Token{
		{Type: Operand, Value: "3.14"},
		{Type: Operator, Value: "+"},
		{Type: Operand, Value: "2"},
		{Type: Operator, Value: "^"},
		{Type: Operand, Value: "10"},
		{Type: Operator, Value: "-"},
		{Type: Operand, Value: "13"},
		{Type: Operator, Value: "*"},
		{Type: LeftParen, Value: "("},
		{Type: Operand, Value: "27"},
		{Type: Operator, Value: "-"},
		{Type: Operand, Value: "81"},
		{Type: Operator, Value: "*"},
		{Type: Operand, Value: "23"},
		{Type: Operator, Value: "^"},
		{Type: LeftParen, Value: "("},
		{Type: Operand, Value: "7"},
		{Type: Operator, Value: "-"},
		{Type: Operand, Value: "11"},
		{Type: RightParen, Value: ")"},
		{Type: RightParen, Value: ")"},
		{Type: Operator, Value: "/"},
		{Type: Operand, Value: "3"},
		{Type: Operator, Value: "-"},
		{Type: Operand, Value: "2.718"},
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
