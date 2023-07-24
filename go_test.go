package main

import (
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func TestArithmeticOperations(t *testing.T) {
	// Test addition
	if result := Add(5, 2); result != 7 {
		t.Errorf("Add(5, 2) returned %d, expected %d", result, 7)
	}

	// Test subtraction
	if result := Subtract(10, 3); result != 7 {
		t.Errorf("Subtract(10, 3) returned %d, expected %d", result, 7)
	}

	// Test multiplication
	if result := Multiply(4, 5); result != 20 {
		t.Errorf("Multiply(4, 5) returned %d, expected %d", result, 20)
	}

	// Test division
	if result := Divide(15, 3); result != 5 {
		t.Errorf("Divide(15, 3) returned %d, expected %d", result, 5)
	}

	// Test division by zero
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Divide(10, 0) did not panic")
		}
	}()
	Divide(10, 0)
}

func TestMain(m *testing.M) {
	// Run tests
	m.Run()
}
