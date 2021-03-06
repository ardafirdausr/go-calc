package main

import (
	"fmt"

	"github.com/ardafirdausr/go-calc/operation"
)

func main() {
	var operator string
	var a, b int

	fmt.Println("Enter <operator> <operand 1> <operand 1>")
	fmt.Println("Available operator: add, sub, mul, div")
	fmt.Print("Enter command: ")
	_, err := fmt.Scan(&operator, &a, &b)
	if err != nil {
		panic("Failed read input")
	}

	switch operator {
	case "add":
		fmt.Println(operation.Add(a, b))
	case "sub":
		fmt.Println(operation.Sub(a, b))
	case "mul":
		fmt.Println(operation.Mul(a, b))
	case "div":
		result, err := operation.Div(a, b)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}
}
