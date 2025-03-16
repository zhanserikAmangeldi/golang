package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Knetic/govaluate"
)

func main() {
	fmt.Println("Write your math expression: ")
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	fmt.Println("Result: ", calculate(input))
}

func calculate(input string) float64 {
	expression, err := govaluate.NewEvaluableExpression(input)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return result.(float64)
}
