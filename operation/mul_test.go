package operation

import (
	"fmt"
	"testing"
)

func TestMulShouldReturnMultiplicationOfOperands(t *testing.T) {
	a, b := 11, 8
	expexted := 88
	result := Mul(a, b)
	if result != expexted {
		err := fmt.Sprintf("Result: %d, Expexted: %d", result, expexted)
		t.Error(err)
	}
}
