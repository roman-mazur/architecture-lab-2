package lab2

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPrefixToPostfix(c *C) {
	testCases := []struct {
		prefixExpression string
		expectedResult   string
	}{
		{"+ 5 5","5 5 +"},
		{"+ + 1 * 2 3 9", "1 2 3 * + 9 +"},
		{"+ 2 4", "2 4 +"},
		{"* - 3 4 + 3 4", "3 4 - 3 4 + *"},
		{"/ - - / 8 3 1 + - 3 2 4 4", "8 3 / 1 - 3 2 - 4 + - 4 /"},
		{"+ + + - - + - * 32 42 / 4 12 3 12 44 5 90 12","32 42 * 4 12 / - 3 + 12 - 44 - 5 + 90 + 12 +"},
	
	//	{"",""}, ERROR CASE
	//	{"% 3 3","3 3 %"}, ERROR CASE
	}

	for _, tc := range testCases {
		result, err := PrefixToPostfix(tc.prefixExpression)
		c.Assert(err, IsNil)
		c.Assert(result, Equals, tc.expectedResult)
	}
}

