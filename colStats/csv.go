package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

type statsFunc func(data []float64) float64

func sum(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

func csv2float(r io.Reader, column int) ([]float64, error) {
	cr := csv.NewReader(r)
	cr.ReuseRecord = true
	// adjusting for 0-indexed
	column--

	var data []float64

	// looping through all records
	for i := 0; ; i++ {
		row, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Cannot read data from file: %w", err)
		}
		if i == 0 {
			// skip header
			continue
		}
		// checking number of columns in csv file
		if len(row) <= column {
			// file does not have that manu columns
			return nil, fmt.Errorf("%w: File has only %d columns", ErrInvalidaColumn, len(row))
		}

		// try to convert data read into a float number
		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}
		data = append(data, v)
	}

	return data, nil
}
