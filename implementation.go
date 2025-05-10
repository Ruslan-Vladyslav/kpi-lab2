package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: document this function.
// PrefixToPostfix converts
func PrefixToPostfix(input string) (string, error) {
	if len(input) == 0 {
		return "", fmt.Errorf("input is empty")
	}

	tokens := strings.Fields(input)
	stack := []string{}

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		_, err := strconv.ParseFloat(token, 64)

		if err == nil {
			stack = append(stack, token)
		} else {
			if token == "+" || token == "-" || token == "*" || token == "/" || token == "^" {
				if len(stack) < 2 {
					return "", fmt.Errorf("invalid prefix expression")
				}

				operand_1 := stack[len(stack)-1]
				operand_2 := stack[len(stack)-2]

				var result string

				switch token {
				case "+":
					result = fmt.Sprintf("(%s %s %s)", token, operand_1, operand_2)
				case "-":
					result = fmt.Sprintf("(%s %s %s)", token, operand_1, operand_2)
				case "*":
					result = fmt.Sprintf("(%s %s %s)", token, operand_1, operand_2)
				case "/":
					result = fmt.Sprintf("(%s %s %s)", token, operand_1, operand_2)
				case "^":
					result = fmt.Sprintf("(pow %s %s)", operand_1, operand_2)
				}

				stack = append(stack[:len(stack)-2], result)
			} else {
				return "", fmt.Errorf("invalid token: %s", token)
			}
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid prefix expression")
	}

	return stack[0], nil
}
