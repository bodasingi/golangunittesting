// main_test.go
package main

import "testing"

// Mock of FunctionB
func mockFunctionB(input int) int {
	return 3 * input // For testing, we'll simulate FunctionB returning input * 3
}

func TestFunctionA(t *testing.T) {
	// Replace FunctionB with our mock
	originalFunctionB := FunctionB
	FunctionB = mockFunctionB
	defer func() { FunctionB = originalFunctionB }() // Restore the original FunctionB after the test

	// Test cases
	testCases := []struct {
		input    int
		expected int
	}{
		{1, 4}, // FunctionA(FunctionB(1)) = FunctionA(3) = 4
		{2, 7}, // FunctionA(FunctionB(2)) = FunctionA(6) = 7
	}

	for _, tc := range testCases {
		actual := FunctionA(tc.input)

		if actual != tc.expected {
			t.Errorf("For input %d, got %d, expected %d", tc.input, actual, tc.expected)
		}
	}
}
