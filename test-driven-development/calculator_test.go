package main

import "testing"

func TestAddition(t *testing.T) {
	calc := Calculator{}

	result := calc.Add(2, 3)

	if result != 5 {
		t.Errorf("Expected Add(2, 3) to return 5, got %d", result)
	}
}

func TestSubstraction(t *testing.T) {
	calc := Calculator{}
	result := calc.Subtract(5, 3)

	if result != 5 {
		t.Errorf("Expected Substract(5, 3) to return 2, got %d", result)
	}
}

func TestMultiplication(t *testing.T) {
	calc := Calculator{}
	result := calc.Multiply(4, 5)
	if result != 20 {
		t.Errorf("Expected Multiply(4, 5) to return 20, got %d", result)
	}
}

func TestDivision(t *testing.T) {
	calc := Calculator{}
	result := calc.Divide(10, 2)
	if result != 5 {
		t.Errorf("Expected Divide(10, 2) to return 5, got %d", result)
	}
}

func TestDivisionByZero(t *testing.T) {
	calc := Calculator{}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Divide(5, 0) to panic, but it didn't")
		}
	}()

	calc.Divide(5, 0)
}

func TestCalculatorOperations(t *testing.T) {
	calc := Calculator{}

	// Define test cases
	testCases := []struct {
		name      string
		a, b      int
		operation func(Calculator, int, int) int
		expected  int
	}{
		{"Addition", 2, 3, Calculator.Add, 5},
		{"Subtraction", 5, 3, Calculator.Subtract, 2},
		{"Multiplication", 4, 5, Calculator.Multiply, 20},
		{"Division", 10, 2, Calculator.Divide, 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.operation(calc, tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Expected %s(%d, %d) to return %d, got %d",
					tc.name, tc.a, tc.b, tc.expected, result)
			}
		})
	}
}
