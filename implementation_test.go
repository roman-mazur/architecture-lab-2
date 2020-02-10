package main

import (
	"fmt"
	"testing"
)

func TestPrefixToPostfix(t *testing.T) {
	res, _ := PrefixToPostfix("+ 5 * - 4 2 3")
	expected := "4 2 - 3 * 5 +"
	if res != expected {
		t.Errorf("Unexpected conversion: %s vs %s", res, expected)
	}
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
