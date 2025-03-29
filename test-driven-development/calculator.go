package main

type Calculator struct{}

func (c Calculator) Add(a, b int) int {
	return a + b
}

func (c Calculator) Subtract(a, b int) int {
	return a - b
}

func (c Calculator) Multiply(a, b int) int {
	return a * b
}

func (c Calculator) Divide(a, b int) int {
	if b == 0 {
		panic("Cannot divide by zero")
	}
	return a / b
}
