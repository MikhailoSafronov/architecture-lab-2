package lab2

import (
	"errors"
	"strconv"
	"strings"
)

// EvaluatePostfix обчислює постфіксний вираз
func EvaluatePostfix(expression string) (string, error) {
	tokens := strings.Fields(expression)
	stack := []float64{}

	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return "", errors.New("недостатньо операндів для операції")
			}

			a, b := stack[len(stack)-2], stack[len(stack)-1]
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
					return "", errors.New("ділення на нуль")
				}
				stack = append(stack, a/b)
			case "^":
				res := 1.0
				for i := 0; i < int(b); i++ {
					res *= a
				}
				stack = append(stack, res)
			default:
				return "", errors.New("невідомий оператор: " + token)
			}
		}
	}

	if len(stack) != 1 {
		return "", errors.New("неправильний вираз")
	}

	return strconv.FormatFloat(stack[0], 'f', -1, 64), nil
}
