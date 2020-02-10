package lab1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrefixToPostfix(t *testing.T) {
	res, err := PrefixToPostfix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 2 - 3 * 5 +", res)
	}
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
