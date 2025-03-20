package helpers

import "errors"

func Sum(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func Login(username, password string) (bool, error) {
	if username == "" || password == "" {
		return false, errors.New("Username and password are required")
	}

	if username != "admin" || password != "admin" {
		return false, errors.New("Invalid username or password")
	}

	return true, nil
}
