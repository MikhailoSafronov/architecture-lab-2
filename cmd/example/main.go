package main

import (
	"flag"
	"fmt"
	"io"
	"lab2"
	"os"
	"strings"
)

func main() {
	var (
		exprFlag   = flag.String("e", "", "Postfix expression to evaluate")
		fileFlag   = flag.String("f", "", "Path to file with postfix expression")
		outputFlag = flag.String("o", "", "Path to output file")
	)
	flag.Parse()

	if *exprFlag != "" && *fileFlag != "" {
		fmt.Fprintln(os.Stderr, "Error: cannot use both -e and -f")
		os.Exit(1)
	}

	var input io.Reader
	switch {
	case *exprFlag != "":
		input = strings.NewReader(*exprFlag)
	case *fileFlag != "":
		file, err := os.Open(*fileFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	default:
		fmt.Fprintln(os.Stderr, "Error: must provide either -e or -f")
		os.Exit(1)
	}

	var output io.Writer
	if *outputFlag != "" {
		file, err := os.Create(*outputFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
