package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(input string) Expression {

	tokens := strings.Split(input, " ")
	stack := []Expression{}

	for _, token := range tokens {
		if token != "+" && token != "*" {
			number, _ := strconv.Atoi(token)
			stack = append(stack, &NumberExpression{number})
			continue
		}

		left := stack[len(stack)-2]
		right := stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		if token == "+" {
			stack = append(stack, &AddExpression{left, right})
		} else {
			stack = append(stack, &MultiplyExpression{left, right})
		}

	}

	return stack[0]
}

func main() {
	input := "3 4 + 2 *"
	result := Parse(input).Interpret()

	fmt.Printf("Result: %d\n", result)
	// Result: 14
}
