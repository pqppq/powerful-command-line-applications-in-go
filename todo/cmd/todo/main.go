package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pqppq/todo"
)

// hardcoding the file name
const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	// use  the Get method to read ToDo items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	switch {
	case len(os.Args) == 1:
		// list current ToDo items
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		// concatenate all provided arguments with a space
		// and add to the list as an item
		item := strings.Join(os.Args[1:], " ")
		// add th task
		l.Add(item)

		// save the new list
		if err := l.Save(todoFileName) ; err != nil{
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
