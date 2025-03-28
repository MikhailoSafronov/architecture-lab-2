package lab2

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestComputeHandler_Valid(t *testing.T) {
	input := strings.NewReader("5 5 +")
	output := &bytes.Buffer{}

	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := "10"
	if output.String() != expected {
		t.Errorf("Expected %q, got %q", expected, output.String())
	}
}

func TestComputeHandler_Invalid(t *testing.T) {
	input := strings.NewReader("invalid")
	output := &bytes.Buffer{}

	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestComputeHandler_Empty(t *testing.T) {
	input := strings.NewReader("")
	output := &bytes.Buffer{}

	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if err == nil {
		t.Error("Expected error for empty input, got nil")
	}
}

func ExampleComputeHandler() {
	input := strings.NewReader("5 3 +")
	output := &bytes.Buffer{}

	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	if err := handler.Compute(); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(output.String())
	
}
