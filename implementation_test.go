package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluatePostfix(t *testing.T) {
	testCases := []struct {
		expression string
		expected   int
		isError    bool
	}{
		{"5 4 2 - 3 2 ^ * +", 23, false},
		{"2 3 +", 5, false},
		{"7 3 - 2 *", 8, false},
		{"8 2 /", 4, false},
		{"3 4 ^", 81, false},
		{"1 +", 0, true},
		{"5 0 /", 0, true},
		{"a b +", 0, true},
	}

	for _, tc := range testCases {
		result, err := EvaluatePostfix(tc.expression)
		if tc.isError {
			assert.Error(t, err, "expected error for input '%s'", tc.expression)
		} else {
			assert.NoError(t, err, "unexpected error for input '%s'", tc.expression)
			assert.Equal(t, tc.expected, result, "incorrect result for input '%s'", tc.expression)
		}
	}
}

func ExampleEvaluatePostfix() {
	res, _ := EvaluatePostfix("5 1 2 + 4 * + 3 -")
	fmt.Println(res)
	// Output: 14
}
