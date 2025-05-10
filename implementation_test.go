package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToLisp(t *testing.T) {
	// Simple expressions with 2-3 operands
	test1, err := PrefixToLisp("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "(+ 5 (* (- 4 2) 3))", test1)
	}

	test2, err := PrefixToLisp("+ 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "(+ 2 3)", test2)
	}

	// Hard expressions with more operands
	test3, err := PrefixToLisp("+ 5 * - 4 2 ^ 3 2")
	if assert.Nil(t, err) {
		assert.Equal(t, "(+ 5 (* (- 4 2) (pow 3 2)))", test3)
	}

	test4, err := PrefixToLisp("* + 2 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "(* (+ 2 2) 3)", test4)
	}

	test5, err := PrefixToLisp("+ * 2 + 3 4 5")
	if assert.Nil(t, err) {
		assert.Equal(t, "(+ (* 2 (+ 3 4)) 5)", test5)
	}

	// Hard expressions with 7-10 operands
	test6, err := PrefixToLisp("+ 1 + 2 + 3 + 4 + 5 + 6 + 7 8")
	if assert.Nil(t, err) {
		assert.Equal(t, "(+ 1 (+ 2 (+ 3 (+ 4 (+ 5 (+ 6 (+ 7 8)))))))", test6)
	}

	test7, err := PrefixToLisp("+ * + 2 2 3 + + 4 5 + 6 7")
	if assert.Nil(t, err) {
		assert.Equal(t, "(+ (* (+ 2 2) 3) (+ (+ 4 5) (+ 6 7)))", test7)
	}

}

func TestPrefixToLisp_InvalidInputs(t *testing.T) {
	// Empty line
	test8, err := PrefixToLisp("")
	assert.NotNil(t, err)
	assert.Equal(t, "", test8)

	_, err = PrefixToLisp("+ a 2")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid token")

	_, err = PrefixToLisp("+ 1")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "not enough operands")

	_, err = PrefixToLisp("+ 1 2 3")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "many operands")
}

func ExamplePrefixToLisp() {
	res, _ := PrefixToLisp("+ 5 * - 4 2 ^ 3 2")
	fmt.Println(res)

	// Output:
	// (+ 5 (* (- 4 2) (pow 3 2)))
}
