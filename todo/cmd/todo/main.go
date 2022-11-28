package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pqppq/todo"
)

// default file name
var todoFileName = ".todo.json"

func main() {
	// check if the use defined the env variable for a custom file name
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"%s tool. Developed for The Pragmatic Bookshelf\n",
			os.Args[0],
		)
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2022\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	task := flag.String("task", "", "Task to be included int the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	// define an items list
	l := &todo.List{}

	// use  the Get method to read ToDo items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// decide what to do based on provided flags
	switch {
	case *list:
		// list current ToDo items
		fmt.Print(l)
	case *complete > 0:
		// complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)
		// save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
