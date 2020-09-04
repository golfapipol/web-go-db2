package business

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

type testcase struct {
	expected int
	input    string
}

func TestGenderMapNumber(t *testing.T) {
	testcases := []testcase{
		{1, "M"},
		{2, "F"},
		{0, "G"},
		{0, "MF"},
	}
	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("Input %s Should be %d", testcase.input, testcase.expected), func(t *testing.T) {
			assert.Equal(t, testcase.expected, GenderMapNumber(testcase.input))
		})
	}
}
