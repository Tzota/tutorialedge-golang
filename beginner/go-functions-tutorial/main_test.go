package main

import "testing"

func TestChallenge01(t *testing.T) {
	var tests = []struct {
		value1   int
		value2   int
		expected int
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{0, 2, 2},
		{-5, -3, -8},
		{99999, 1, 100000},
	}

	for _, test := range tests {
		result := Add(test.value1, test.value2)
		if result != test.expected {
			t.Error("Expected output did not match the result")
		}
	}
}
