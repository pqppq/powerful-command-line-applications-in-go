package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// calling the count function to count the number of words
	// received from the standard input and printing it out
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	// a scanner is used to read text from a reader (such as files)
	scanner := bufio.NewScanner(r)

	// define the scanner split type to words (default is split by lines)
	scanner.Split(bufio.ScanWords)

	// defining a counter
	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
