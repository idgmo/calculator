package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(numerator, denominator int) (int, error) {

	// error handling for divide by zero

	// if denominator == 0 {
	// 	return 0, errors.New("divide by zero")
	// }

	return numerator / denominator, nil
}

func modulus(a, b int) int {
	return a % b
}

func main() {

	// variables

	var op string
	var num1, num2 int
	var cont string

	// for loop continues command line interactions until "q" is entered at the prompt

	for cont != "q" {

		fmt.Println("Please enter an expression using integers, spaces, and one of the following operators: +, -, /, %, *:")

		if cont == "" {
			fmt.Println("Example: 255 * 22")
		}

		// gather user input

		reader := bufio.NewReader(os.Stdin)

		// wait until newline (enter) to continue

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input syntax.")
			return
		}

		input = strings.TrimSuffix(input, "\n")

		re, err := regexp.Compile(`^(?P<num1>[0-9]{1,})\s*(?P<op>[+\/\-%*])\s*(?P<num2>[0-9]{1,})$`)
		if err != nil {
			fmt.Println("Invalid input format")
		}

		// match user input against regex

		if re.MatchString(input) {
			matches := re.FindStringSubmatch(input)

			num1, _ = strconv.Atoi(matches[re.SubexpIndex("num1")])
			op = matches[re.SubexpIndex("op")]
			num2, _ = strconv.Atoi(matches[re.SubexpIndex("num2")])

			// maths logic

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

				// error handling for divide by zero: see divide()

				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
					fmt.Println("Zero returned:")
				} else {
					output = num
				}
			case "%":
				output = modulus(num1, num2)
			default:
				fmt.Println("Invalid operation")
			}

			// readout for successful calculation

			fmt.Printf("%d %s %d = %d\n", num1, op, num2, output)
			fmt.Println("Press any letter to clear. Press \"q\" to quit.")

			fmt.Scanf("%s", &cont)
		} else {
			fmt.Printf("\nInvalid input format\n\n")
		}
	}
}
