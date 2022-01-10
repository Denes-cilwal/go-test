package tutorial1

import (
	"fmt"
	"testing"
)

func TestCheckIfPrime(test *testing.T) {
	input := 89
	expected := true
	output := CheckIfPrime(input)
	if output != expected {
		test.Logf("Expected %v, got %v", expected, output)
	}
}

// multiple test case on same function

func TestCheckIfPrimeTable(test *testing.T) {

	testCases := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{30, true},
		{41, true},
		{59, true},
	}
	for _, testCase := range testCases {
		test.Run(fmt.Sprintf("%v", testCase.input), func(t *testing.T) {
			got := CheckIfPrime(testCase.input)
			if got != testCase.expected {
				test.Logf("Expected %v, got %v", testCase.expected, testCase.input)

			}
		})
	}

}
