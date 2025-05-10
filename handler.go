package lab2

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

// Compute reads an expression from Input, processes it, and writes the result to Output.
// Returns an error if the expression is invalid.
func (ch *ComputeHandler) Compute() error {
	reader := bufio.NewReader(ch.Input)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to read input: %w", err)
	}

	line = strings.TrimSpace(line)
	if line == "" {
		return fmt.Errorf("input expression is empty")
	}

	result, err := PrefixToLisp(line)
	if err != nil {
		return err // will be handled in main.go and reported to stderr with exit code 1
	}

	_, writeErr := fmt.Fprintln(ch.Output, result)
	if writeErr != nil {
		return fmt.Errorf("failed to write output: %w", writeErr)
	}

	return nil
}
