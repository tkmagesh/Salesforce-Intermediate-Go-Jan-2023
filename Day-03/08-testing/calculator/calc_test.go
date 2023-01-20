package calculator

import "testing"

func TestAdd(t *testing.T) {
	x, y := 10, 11
	expected := 21
	result := Add(x, y)
	if result != expected {
		t.Errorf("Add(%d, %d), expected = %d but got %d\n", x, y, expected, result)
	}
}
