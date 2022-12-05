package main

import "errors"

var (
	ErrNotNumber         = errors.New("data is not numeric")
	ErrInvalidaColumn    = errors.New("Invalid column number")
	ErrNoFiles           = errors.New("No input files")
	ErrInvalidOperetaion = errors.New("Invalid operation")
)
