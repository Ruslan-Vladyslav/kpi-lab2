package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/Ruslan-Vladyslav/kpi-lab2"
)

func main() {
	// Оголошення прапорців
	exprFlag := flag.String("e", "", "Expression string")
	fileFlag := flag.String("f", "", "Input file with expression")
	outputFlag := flag.String("o", "", "Output file for result (optional)")

	flag.Parse()

	// Перевірка: не можна вказувати одночасно -e та -f
	if *exprFlag != "" && *fileFlag != "" {
		fmt.Fprintln(os.Stderr, "Error: cannot use both -e and -f at the same time")
		os.Exit(1)
	}

	var input io.Reader
	var output io.Writer

	// Обробка input
	switch {
	case *exprFlag != "":
		input = io.NopCloser(strings.NewReader(*exprFlag))
	case *fileFlag != "":
		file, err := os.Open(*fileFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	default:
		fmt.Fprintln(os.Stderr, "Error: either -e or -f must be specified")
		os.Exit(1)
	}

	// Обробка output
	if *outputFlag != "" {
		file, err := os.Create(*outputFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintf(os.Stderr, "Syntax error: %v\n", err)
		os.Exit(1)
	}
}
