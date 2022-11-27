package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Byte count")

	// parsing the flags provided by the user
	flag.Parse()

	// calling the count function to count the number of words
	// received from the standard input and printing it out
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	// a scanner is used to read text from a reader (such as files)
	scanner := bufio.NewScanner(r)

	if countLines && countBytes {
		fmt.Println("both -l flang and -b flag are set.")
	}

	// if the count lines flag is not set, we want to count words, we define
	// the scanner split type to words (default is split by lines)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	// if the count bytes flag is set, we want to count bytes , we define
	// the scanner split type to bytes
	if  countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	// defining a counter
	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
