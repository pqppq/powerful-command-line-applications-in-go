package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// verify and parse arguments
	op := flag.String("op", "sum", "Operaiton to be executed")
	column := flag.Int("col", 1, "CSV column on which to execute operation")

	flag.Parse()

	if err := run(flag.Args(), *op, *column, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func run(filenames []string, op string, column int, out io.Writer) error {

	if len(filenames) == 0 {
		return ErrNoFiles
	}

	if column < 1 {
		return fmt.Errorf("%w: %d", ErrInvalidaColumn, column)
	}

	var opFunc statsFunc

	switch op {
	case "sum":
		opFunc = sum
	case "avg":
		opFunc = avg
	default:
		return fmt.Errorf("%w: %s", ErrInvalidOperetaion, op)
	}

	consolidate := make([]float64, 0)

	// loop through all files and their data to consolidate
	for _, fname := range filenames {
		// open the file for reading
		f, err := os.Open(fname)
		if err != nil {
			return fmt.Errorf("Cannot open file: %w", err)
		}

		// parse the csv into a slice of float64 numbers
		data, err := csv2float(f, column)
		if err != nil {
			return err
		}

		if err := f.Close(); err != nil {
			return err
		}

		// append the data to consolidate
		consolidate = append(consolidate, data...)
	}

	// execute operation
	_, err := fmt.Fprintln(out, opFunc(consolidate))

	return err
}
