package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	result := calculate(num1, num2, operator)

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}

func calculate(num1, num2 float64, operator string) float64 {
	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Division by zero")
			return 0
		}
	default:
		fmt.Println("Error: Invalid operator")
		return 0
	}

	return result
}
