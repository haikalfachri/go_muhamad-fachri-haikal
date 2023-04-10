package calculator

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(2, 3)
	expected := 5.0
	if result != 5 {
		t.Errorf("Addition(2, 3) = %f; expected %f", result, expected)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(2, 3)
	expected := -1.0
	if result != -1 {
		t.Errorf("Subtraction(2, 3) = %f; expected %f", result, expected)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(2, 3)
	expected := 6.0
	if result != 6 {
		t.Errorf("Multiplication(2, 3) = %f; expected %f", result, expected)
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(6, 3)
	expected := 2.0
	if result != expected|| err != nil {
		t.Errorf("Division(6, 3) = (%f, %v); expected (%f, nil)", result, err, expected)
	}

	result, err = Division(6, 0)
	if err == nil {
		t.Errorf("Division(6, 0) is not divided by zero")
	}
}