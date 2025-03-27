package lab2

import (
	"math"
	"strconv"
	"strings"
)

func CalculatePostfix(expression string) (int, error) {
	stack := []int{}

	tokens := strings.Fields(expression)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/", "^":
			if len(stack) < 2 {
				return 0, strconv.ErrSyntax
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, strconv.ErrSyntax
				}
				stack = append(stack, a/b)
			case "^":
				stack = append(stack, int(math.Pow(float64(a), float64(b))))
			}
		default:
			number, err := strconv.Atoi(token)
			if err != nil {
				return 0, err
			}
			stack = append(stack, number)
		}
	}

	if len(stack) != 1 {
		return 0, strconv.ErrSyntax
	}

	return stack[0], nil
}
