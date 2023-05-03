package main

import (
	"fmt"
	"strconv"
)

func calculator(input string) (float64, error) {
	tokens := tokenize(input)
	postfix := infixToPostfix(tokens)
	result, err := evaluatePostfix(postfix)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func tokenize(input string) []string {
	tokens := []string{}
	current := ""
	for _, c := range input {
		if isOperator(string(c)) || string(c) == "(" || string(c) == ")" {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			tokens = append(tokens, string(c))
		} else {
			current += string(c)
		}
	}
	if current != "" {
		tokens = append(tokens, current)
	}
	return tokens
}

func infixToPostfix(tokens []string) []string {
	output := []string{}
	operatorStack := []string{}
	for _, token := range tokens {
		if isNumber(token) {
			output = append(output, token)
		} else if isOperator(token) {
			for len(operatorStack) > 0 && isOperator(operatorStack[len(operatorStack)-1]) &&
				hasHigherPrecedence(token, operatorStack[len(operatorStack)-1]) {
				output = append(output, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			operatorStack = append(operatorStack, token)
		} else if token == "(" {
			operatorStack = append(operatorStack, token)
		} else if token == ")" {
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != "(" {
				output = append(output, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			if len(operatorStack) == 0 || operatorStack[len(operatorStack)-1] != "(" {
				return []string{}
			}
			operatorStack = operatorStack[:len(operatorStack)-1]
		} else {
			return []string{}
		}
	}
	for len(operatorStack) > 0 {
		if operatorStack[len(operatorStack)-1] == "(" {
			return []string{}
		}
		output = append(output, operatorStack[len(operatorStack)-1])
		operatorStack = operatorStack[:len(operatorStack)-1]
	}
	return output
}

func evaluatePostfix(postfix []string) (float64, error) {
	valueStack := []float64{}
	for _, token := range postfix {
		if isNumber(token) {
			value, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			valueStack = append(valueStack, value)
		} else if isOperator(token) {
			if len(valueStack) < 2 {
				return 0, fmt.Errorf("Invalid expression")
			}
			b := valueStack[len(valueStack)-1]
			a := valueStack[len(valueStack)-2]
			valueStack = valueStack[:len(valueStack)-2]
			result, err := evaluateBinaryOperator(a, b, token)
			if err != nil {
				return 0, err
			}
			valueStack = append(valueStack, result)
		} else {
			return 0, fmt.Errorf("Invalid token: %s", token)
		}
	}
	if len(valueStack) != 1 {
		return 0, fmt.Errorf("Invalid expression")
	}
	return valueStack[0], nil
}

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}

func hasHigherPrecedence(op1 string, op2 string) bool {
	precedence := map[string]int{
		"+": 3,
		"-": 3,
		"*": 2,
		"/": 2,
		"^": 1,
	}
	return precedence[op1] > precedence[op2]
}

func evaluateBinaryOperator(a float64, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		return a / b, nil
	case "^":
		return pow(a, b), nil
	default:
		return 0, fmt.Errorf("Invalid operator")
	}
}

func pow(a float64, b float64) float64 {
	result := 1.0
	for i := 0; i < int(b); i++ {
		result *= a
	}
	return result
}
