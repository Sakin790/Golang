package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	// Take input
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter operator (+ - * /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	// Perform calculation
	switch operator {
	case "+":
		fmt.Printf("Result: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Result: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Result: %.2f\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
		} else {
			fmt.Printf("Result: %.2f\n", num1/num2)
		}
	default:
		fmt.Println("Invalid operator! Please use +, -, * or /.")
	}
}
