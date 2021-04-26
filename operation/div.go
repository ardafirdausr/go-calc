package operation

import "errors"

func Div(a, b int) (float32, error) {
	if b == 0 {
		return float32(0), errors.New("Divider cannot be 0")
	}
	return float32(a) / float32(b), nil
}
