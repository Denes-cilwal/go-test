package tutorial_testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	// define inputs
	inputA := 1
	inputB := 2

	// defined expected outputs

	expected := 3

	// perform test

	actual := Add(inputA, inputB)

	// assert that the actual result is equal to expected
	assert.Equal(t, expected, actual)
}
