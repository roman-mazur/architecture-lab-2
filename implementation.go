package lab2

import (
	"errors"
	"strings"
)

// PrefixToPostfix converts prefix notation to postfix notation.
func PrefixToPostfix(input string) (string, error) {
	stack := []string{}
	tokens := strings.Fields(input)

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		if isOperator(token[0]) {
			if len(stack) < 2 {
				return "", errors.New("not enough operands for operator")
			}
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			postfix := operand1 + " " + operand2 + " " + token
			stack = append(stack, postfix)
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "", errors.New("invalid expression format")
	}

	return stack[0], nil
}

func isOperator(ch byte) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '^'
}

