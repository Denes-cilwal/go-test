package tutorial1

import "testing"

func TestCheckIfPrime(test *testing.T) {
	input := 89
	expected := true
	output := CheckIfPrime(input)
	if output != expected {
		test.Logf("Expected %v, got %v", expected, output)
	}
}
