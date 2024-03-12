package lab2

import (
	"testing"
        . "gopkg.in/check.v1"	
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})


func (s *MySuite) TestPrefixToPostfix(c *C) {
	test1Result,e1 := PrefixToPostfix("+ + 1 * 2 3 9")
  	c.Assert(test1Result, Equals, "1 2 3 * + 9 +")

	test2Result,e2 := PrefixToPostfix("+ + 1 * 2 3 9")
  	c.Assert(test2Result, Equals, "1 2 3 * + 9 +")

	test3Result,e3 := PrefixToPostfix("+ + 1 * 2 3 9")
  	c.Assert(test3Result, Equals, "1 2 3 * + 9 +")

	test4Result,e4 := PrefixToPostfix("+ + 1 * 2 3 9")
  	c.Assert(test4Result, Equals, "1 2 3 * + 9 +")

}

