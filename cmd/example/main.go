package main

import (
	"fmt"
	lab1 "github.com/roman-mazur/design-lab-1"
)

func main() {
	// TODO: Get input from the command line, handle errors.
	res, _ := lab1.PrefixToPostfix("+ 2 2")
	fmt.Println(res)
}
