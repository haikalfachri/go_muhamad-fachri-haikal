package calculator

import (
	"errors"
)

func Addition(a, b float64) float64{
	return a + b
}

func Subtraction(a, b float64) float64 {
	return a - b
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func Multiplication(a, b float64) float64 {
	return a * b
}