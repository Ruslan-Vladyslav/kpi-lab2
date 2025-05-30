package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

// PrefixToLisp converts a prefix expression into a Lisp-style expression
func PrefixToLisp(input string) (string, error) {
	if len(input) == 0 {
		return "", fmt.Errorf("input is empty")
	}

	tokens := strings.Fields(input)
	stack := []string{}

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		if isValidNumber(token) {
			stack = append(stack, token)
		} else if IsValidToken(token) {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid expression: not enough operands for operator %s", token)
			}

			operand_1 := stack[len(stack)-1]
			operand_2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

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

			stack = append(stack, result)
		} else {
			return "", fmt.Errorf("invalid token: %s", token)
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid expression: many operands in final stack: %v", stack)
	}

	return stack[0], nil
}

// Validation functions:
func IsValidToken(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}

func isValidNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}
