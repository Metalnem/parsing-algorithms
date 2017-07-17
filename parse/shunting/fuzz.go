// +build gofuzz

package shunting

func Fuzz(data []byte) int {
	if _, err := New().Parse(string(data)); err != nil {
		return 0
	}

	return 1
}
