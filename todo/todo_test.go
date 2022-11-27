package todo_test

import (
	"testing"
	"os"
	"io/ioutil"
	"github.com/pqppq/todo"
)

// tests the Add method of the List type
func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New task"
	l.Add(taskName)

	if l[0].Task != taskName {
		// %q expand value with quates
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
}

// tests the Complete method of the list type
func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be completd.")
	}

	l.Complete(1)
	if !l[0].Done {
		t.Errorf("New task should be completed.")
	}
}

// tests the Delete method of the List type
func TestDelete(t *testing.T) {
	l := todo.List{}
	tasks := []string{
		"New task 1",
		"New task 2",
		"New task 3",
	}
	for _, v := range tasks {
		l.Add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead.", tasks[0], l[0].Task)
	}
	l.Delete(2)
	if len(l) != len(tasks)-1 {
		t.Errorf("Expected list length %d, got %d instead.", len(tasks)-1, len(l))
	}
	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], l[1].Task)
	}
}

// tests teh Save and Get method of the List type
func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l1[0].Task)
	}

	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s.", err)
	}
	defer os.Remove(tf.Name())

	if err  :=  l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file:%s.", err)
	}
	if err  :=  l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file:%s", err)
	}
	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task.", l1[0].Task, l2[0].Task)
	}
}
