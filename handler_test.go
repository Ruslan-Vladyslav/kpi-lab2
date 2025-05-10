package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHandler_ValidInput(t *testing.T) {
	input := strings.NewReader("+ 2 3\n")
	var output bytes.Buffer

	handler := ComputeHandler{
		Input:  input,
		Output: &output,
	}

	err := handler.Compute()

	assert.Nil(t, err)
	assert.Equal(t, "(+ 2 3)\n", output.String())
}

func TestComputeHandler_InvalidInputSyntax(t *testing.T) {
	input := strings.NewReader("+ 2\n")
	var output bytes.Buffer

	handler := ComputeHandler{
		Input:  input,
		Output: &output,
	}

	err := handler.Compute()

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "not enough operands")
	assert.Empty(t, output.String())
}

func TestComputeHandler_EmptyInput(t *testing.T) {
	input := strings.NewReader("\n")
	var output bytes.Buffer

	handler := ComputeHandler{
		Input:  input,
		Output: &output,
	}

	err := handler.Compute()

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "input expression is empty")
	assert.Empty(t, output.String())
}

func TestComputeHandler_InvalidToken(t *testing.T) {
	input := strings.NewReader("+ a 2\n")
	var output bytes.Buffer

	handler := ComputeHandler{
		Input:  input,
		Output: &output,
	}

	err := handler.Compute()

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid token")
	assert.Empty(t, output.String())
}
