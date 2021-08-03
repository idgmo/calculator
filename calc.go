package main

import (
	"errors"
	"fmt"
)

func add(num1, num2 int) int {
	return num1 + num2
}

func subtract(num1, num2 int) int {
	return num1 - num2
}
func multiply(num1, num2 int) int {
	return num1 * num2
}
func divide(num1, num2 int) (int, error) {
	// error handling for divide by zero
	if num2 == 0 {
		return 0, errors.New("divide by zero")
	}
	return num1 / num2, nil
}
func modulus(num1, num2 int) int {
	return num1 % num2
}

func main() {

	// variables
	var op string
	var num1, num2 int
	var cont string

	for cont != "q" {
		fmt.Println("Please enter an equation using one of the following operators: +, -, /, %, *:")
		if cont == "" {
			fmt.Println("Example: 255 * 22")
		}
		fmt.Scanf("%d", &num1)
		fmt.Scanf("%s", &op)
		fmt.Scanf("%d", &num2)

		output := 0
		switch op {
		case "+":
			output = add(num1, num2)
		case "-":
			output = subtract(num1, num2)

		case "*":
			output = multiply(num1, num2)
		case "/":
			num, err := divide(num1, num2)

			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				output = num
			}
		case "%":
			output = modulus(num1, num2)
		default:
			fmt.Println("Invalid operation")
		}
		fmt.Printf("%d %s %d = %d\n", num1, op, num2, output)
		fmt.Println("Press any letter to clear.")
		fmt.Println("Press \"q\" to quit.")
		fmt.Scan(&cont)
	}
}
