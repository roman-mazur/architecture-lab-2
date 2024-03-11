package lab2

import (
	"errors"
	"strings"
	"regexp"
)

func PrefixToPostfix(input string) (string, error) {
	stack := []string{}
	tokens := strings.Fields(input)

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
	

		if(!isNumber(token) && !isOperator(token[0])){
			return "", errors.New("invalid character")
		}

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

	return stack[0], nil
}


func isNumber(input string) bool {
	pattern := `^[+-]?\d*\.?\d+$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(input)
}

func isOperator(ch byte) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '^'
}

