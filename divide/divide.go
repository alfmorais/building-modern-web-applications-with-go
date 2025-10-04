package divide

import (
	"errors"
)

func Divide(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}

	var result float32 = x / y
	return result, nil
}
