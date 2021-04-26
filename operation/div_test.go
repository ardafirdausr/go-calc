package operation

import (
	"fmt"
	"testing"
)

func TestDivShouldReturnDivisionOfOperands(t *testing.T) {
	a, b := 88, 8
	var expexted float32 = 11
	result, err := Div(a, b)
	if err != nil {
		t.Error(err)
	}

	if result != expexted {
		err := fmt.Sprintf("Result: %f, Expexted: %f", result, expexted)
		t.Error(err)
	}
}

func TestDivShouldReturnErrorWhenDividerIsZero(t *testing.T) {
	a, b := 88, 0
	_, err := Div(a, b)
	if err == nil {
		t.Error("Div function should return error")
	}
}
