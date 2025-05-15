package main

import (
	"fmt"
	"math"
)

func main() {
	var first, second float64
	var opString string

	fmt.Println("Enter first number")
	fmt.Scanln(&first)
	fmt.Println("Enter second number")
	fmt.Scanln(&second)

	fmt.Println("Enter arithmetic operator [ + | - | * | / ]")
	fmt.Scanln(&opString)

	operator := rune(opString[0])

	result := calculate(first, second, operator)
	if math.IsNaN(result) {
		fmt.Println("Result is", result)
	}
}

// calculate performs arithmetic operations on given 2 numbers and operator and return result
func calculate(num1, num2 float64, operator rune) float64 {
	switch operator {
	case '+':
		return num1 + num2
	case '-':
		return num1 - num2
	case '*':
		return num1 * num2
	case '/':
		if num2 == 0.0 {
			fmt.Println("Division by zero")
			return math.NaN()
		}
		return num1 / num2
	default:
		fmt.Println("Invalid operator")
		return math.NaN()
	}
}
