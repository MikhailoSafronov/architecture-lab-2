package lab2

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	data, err := io.ReadAll(ch.Input)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	expression := strings.TrimSpace(string(data))
	if expression == "" {
		return errors.New("empty expression")
	}

	result, err := EvaluatePostfix(expression)
	if err != nil {
		return fmt.Errorf("computation error: %w", err)
	}

	_, err = ch.Output.Write([]byte(result))
	return err
}
