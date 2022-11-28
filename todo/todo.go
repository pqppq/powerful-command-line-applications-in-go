package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// item struct represents a ToDO item
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

// creates a new todo item and appends it to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// delete the item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist.", i)
	}

	*l = append(ls[ :i-1 ], ls[i:]...)
	return nil
}

// marks a ToDo item as completed by
// setting Done = true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist.", i)
	}
	// adjusting index for 0-indexed
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// encodes the List as JSON and saves it
// using the provided file name
func (l *List) Save(filename string) error {
	// parse List to JSON
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, js, 0644)
}

// opens the provided file name, decodes
// the JSON data and parses it into a List
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	// parse JSON to List
	return json.Unmarshal(file, l)
}

func (l *List) String() string {
	formatted := ""
	for k,t := range *l{
		prefix := "[ ] "
		if t.Done {
			prefix = "[x] "
		}

		// adjusting the item number k to preint numbers starting from 1 instead of 0
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}
