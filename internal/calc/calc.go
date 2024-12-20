package calc

import (
	"errors"
	"strconv"
)

type operator struct {
	Symbol   string
	Priority int
}

var Errors = struct {
	DivisionByZero        error
	MismatchedParentheses error
	InvalidExpression     error
}{
	DivisionByZero:        errors.New("деление на ноль"),
	MismatchedParentheses: errors.New("ошибка в расставлении скобок"),
	InvalidExpression:     errors.New("неправильное выражение"),
}

var operators = map[string]operator{
	"*": {"*", 2},
	"/": {"/", 2},
	"+": {"+", 1},
	"-": {"-", 1},
}

func transform(expression string) []string {
	var numbers []string
	var currentNumber string

	for _, char := range expression {
		if (char >= '0' && char <= '9') || char == '.' {
			currentNumber += string(char)
		} else {
			if currentNumber != "" {
				numbers = append(numbers, currentNumber)
				currentNumber = ""
			}
			if char != ' ' {
				numbers = append(numbers, string(char))
			}
		}
	}

	if currentNumber != "" {
		numbers = append(numbers, currentNumber)
	}

	return numbers
}

func toRPN(expression string) ([]string, error) {
	var stack []string
	var output []string

	for _, char := range transform(expression) {
		if isDigit(char) {
			output = append(output, char)
		} else if op, exists := operators[char]; exists { //выглдяит страшно, мб упростить?
			if len(stack) > 0 && operators[stack[len(stack)-1]].Priority >= op.Priority {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, op.Symbol)
		} else if char == "(" {
			stack = append(stack, "(")
		} else if char == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, Errors.MismatchedParentheses
			}
			stack = stack[:len(stack)-1]
		} else {
			return nil, Errors.InvalidExpression
		}
	}
	for len(stack) > 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return output, nil
}

func caclulateRPN(output []string) (float64, error) {
	var stack []float64

	for _, char := range output {
		if isDigit(char) {
			number, _ := strconv.ParseFloat(char, 64)
			stack = append(stack, number)
		}
		if op, exists := operators[char]; exists {
			if len(stack) < 2 {
				return 0, Errors.InvalidExpression
			}
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch op.Symbol {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b != 0 {
					stack = append(stack, a/b)
				} else {
					return 0, Errors.DivisionByZero
				}
			default:
				return 0, Errors.InvalidExpression
			}
		}
	}
	if len(stack) != 1 {
		return 0, Errors.InvalidExpression
	}
	return float64(stack[0]), nil

}

func isDigit(char string) bool {
	_, err := strconv.ParseFloat(char, 64)
	return err == nil //мб тут число возращать надо, чтоб сразу сделать с ним выражение
}

func Calculate(expression string) (float64, error) {
	postfix, error1 := toRPN(expression)
	if error1 != nil {
		return 0, error1
	}
	number, error2 := caclulateRPN(postfix)
	if error2 != nil {
		return 0, error2
	}
	return number, nil

}
