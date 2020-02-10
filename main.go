package main

import "fmt"

func main() {
	// TODO: Get input from the command line, handle errors.
	res, _ := PrefixToPostfix("+ 2 2")
	fmt.Println(res)
}
