package main

import "testing"

func TestPrefixToPostfix(t *testing.T) {
	res, _ := PrefixToPostfix("+ 5 * - 4 2 3")
	expected := "4 2 - 3 * 5 +"
	if res != expected {
		t.Errorf("Unexpected conversion: %s vs %s", res, expected)
	}
}
