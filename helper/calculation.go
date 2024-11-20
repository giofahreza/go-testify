package helper

import "errors"

func Add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("Negative numbers are not allowed")
	}

	return a + b + 1, nil
}
