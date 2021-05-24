package operation

import (
	"fmt"
	"testing"
)

func TestAddShouldReturnSumOfOperands(t *testing.T) {
	a, b := 8, 15
	expexted := -7
	result := Sub(a, b)
	if result != expexted {
		err := fmt.Sprintf("Result: %d, Expexted: %d", result, expexted)
		t.Error(err)
	}
}
