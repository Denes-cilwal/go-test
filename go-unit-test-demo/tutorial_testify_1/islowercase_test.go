package tutorial_testify_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAllLowerCase(t *testing.T) {
	// Define a test case struct
	type testCase struct {
		name           string // test case name
		input          string // function input
		expectedResult bool   // expected outcome
		hasError       bool   // a flag to check that an error is expected
	}

	// Define a slice of testCase as test table
	testTable := []testCase{
		{
			name:           "all lowercase string input",
			input:          "aabbcc",
			expectedResult: true,
			hasError:       false,
		},
		{
			name:           "some uppercase string input",
			input:          "aAbBcC",
			expectedResult: false,
			hasError:       false,
		},
		{
			name:           "input string contains a number",
			input:          "aabbc1",
			expectedResult: false,
			hasError:       true,
		},
	}

	// Begin test
	for _, test := range testTable {
		actual, err := IsAllLowerCase(test.input)
		assert.Equal(t, test.expectedResult, actual, test.name)

		// assert.Nil and assert.NotNil to assert that the current test case does not or does expect an error respectively.
		if test.hasError {
			assert.NotNil(t, err, test.name)
		} else {
			assert.Nil(t, err, test.name)
		}
	}
}
