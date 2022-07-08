package main

import "testing"

func TestSum(testing *testing.T) {
	// total := Sum(5, 5)

	// if total != 10 {
	// 	t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
	// }
	tests := []struct {
		a              int
		b              int
		expectedResult int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, test := range tests {
		total := Sum(test.a, test.b)
		if total != test.expectedResult {
			testing.Errorf("Sum was incorrect, got %d expected %d", total, test.expectedResult)
		}
	}
}

func TestMax(testing *testing.T) {
	tests := []struct {
		a              int
		b              int
		expectedResult int
	}{
		{3, 5, 5},
		{3, 2, 2},
	}

	for _, test := range tests {
		max := GetMax(test.a, test.b)
		if max != test.expectedResult {
			testing.Errorf("GetMax was incorrect, got %d, expected %d", max, test.expectedResult)
		}
	}
}

func TestFibonacci(testing *testing.T) {
	tests := []struct {
		value          int
		expectedResult int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, test := range tests {
		realResult := Fibonacci(test.value)
		if realResult != test.expectedResult {
			testing.Errorf("Fibonacci was incorrect, got %d, expected %d", realResult, test.value)
		}
	}
}

// go test -coverprofile=coverage.out
// go tool cover -func=coverage.out
// go tool cover -html=coverage.out
