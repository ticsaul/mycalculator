package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(input string, operator string) int {
	values := strings.Split(input, operator)
	// fmt.Println(values)

	num1 := parse(values[0])
	num2 := parse(values[1])
	switch operator {
	case "+":
		fmt.Println("Operation: ", input)
		return num1 + num2
	case "-":
		fmt.Println("Operation: ", input)
		return num1 - num2
	case "*":
		fmt.Println("Operation: ", input)
		return num1 * num2
	case "/":
		fmt.Println("Operation: ", input)
		return num1 / num2
	default:
		fmt.Println(operator, " is not supported!")
		return 0
	}
}

func parse(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func StartCalc() {
	for true {
		fmt.Println("Input operator [+, -, *, /] or bye to exit calc: ")
		operator := readInput()
		if operator == "bye" {
			os.Exit(3)
		}

		fmt.Println("Input operation: ")
		operation := readInput()
		// fmt.Println(operation)
		c := calc{}
		fmt.Println("Result = ", c.operate(operation, operator))
	}
}
