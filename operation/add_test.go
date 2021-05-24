package operation

import (
	"fmt"
	"testing"
)

func TestSubShouldReturnSubtractionOfOperands(t *testing.T) {
	a, b := 4, 7
	expexted := 11
	result := Add(a, b)
	if result != expexted {
		err := fmt.Sprintf("Result: %d, Expexted: %d", result, expexted)
		t.Error(err)
	}
}
