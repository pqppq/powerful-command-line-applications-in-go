package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"

	"github.com/pqppq/todo"
)

// default file name
var todoFileName = ".todo.json"

// get task function decides where to get the description for a new
// task from: arguments or stdin
func getTask(r io.Reader, args ...string) (*todo.List, error) {
	l := &todo.List {}
	if len(args) > 0 {
		t := strings.Join(args, " ")
		l.Add(t)
		return l, nil
	}
	s := bufio.NewScanner(r)
	for s.Scan() {
		if err := s.Err(); err != nil {
			return l, nil
		}
		t := s.Text()
		if len(t) == 0 {
			return l, fmt.Errorf("Task cannot be blank")
		}
		l.Add(t)
	}
	return l, nil
}

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

	add := flag.Bool("add", false, "Add task to the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	delete := flag.Int("delete", 0, "Delete task")
	remain:= flag.Bool("remain", false, "Show completed tasks")

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
	case *add:
		// when any arguments (excluding flags) are provided, they will be
		// used as the new task
		ls, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Merge(*ls)
		// save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *list:
		if *remain {
			// show only remained task
			fmt.Print(l.RemainedTask())
		} else {
			// show all tasks including completed
			fmt.Print(l)
		}
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
	case *delete > 0:
		// delete item from ToDo list
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
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
