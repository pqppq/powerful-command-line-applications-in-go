/*
Copyright © 2022 pqppq

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

const timeFormat = "Jan/02 @15:04"

var (
	ErrConnection      = errors.New("Connection error")
	ErrNotFound        = errors.New("Not found")
	ErrInvalidResponse = errors.New("Invalid server response")
	ErrInvalid         = errors.New("Invalid data")
	ErrNotNumber       = errors.New("Not a number")
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type response struct {
	Results      []item `json:"results"`
	Date         int    `json:"date"`
	TotalResults int    `json:"total_results"`
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client called")
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}

func newClient() *http.Client {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	return c
}
func getItems(url string) ([]item, error) {
	r, err := newClient().Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrConnection, err)
	}
	// ensure to close body after reading
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		msg, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, fmt.Errorf("Cannot read body: %w", err)
		}
		err = ErrInvalidResponse
		if r.StatusCode == http.StatusNotFound {
			err = ErrNotFound
		}
		return nil, fmt.Errorf("%w: %s", err, msg)
	}

	var resp response
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}
	if resp.TotalResults == 0 {
		return nil, fmt.Errorf("%w: No results found", ErrNotFound)
	}
	return resp.Results, nil
}

func getAll(apiRoot string) ([]item, error) {
	u := fmt.Sprintf("%s/todo", apiRoot)
	return getItems(u)
}

func getOne(url string, id int) (item, error) {
	u := fmt.Sprintf("%s/todo/%d", url, id)

	items, err := getItems(u)
	if err != nil {
		return item{}, err
	}
	if len(items) != 1 {
		return item{}, err
	}
	return items[0], nil
}

func sendRequest(apiRoot, method, contentType string, expStatus int, body io.Reader) error {
	req, err := http.NewRequest(method, apiRoot, body)
	if err != nil {
		return err
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	r, err := newClient().Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != expStatus {
		msg, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("Cannot read body: %w", err)
		}
		err = ErrInvalidResponse
		if r.StatusCode != http.StatusNotFound {
			return fmt.Errorf("%w: %s", err, msg)
		}
	}
	return nil
}

func addItem(apiRoot, task string) error {
	u := fmt.Sprintf("%s/todo", apiRoot)
	item := struct {
		Task string `json:"task"`
	}{
		Task: task,
	}

	var body bytes.Buffer

	if err := json.NewEncoder(&body).Encode(item); err != nil {
		return err
	}

	return sendRequest(u, http.MethodPost, "application/json", http.StatusCreated, &body)
}

func completeItem(apiRoot string, id int) error {
	u := fmt.Sprintf("%s/todo/%d?complete", apiRoot, id)
	return sendRequest(u, http.MethodPatch, "", http.StatusNoContent, nil)
}

func deleteItem(apiRoot string, id int) error {
	u := fmt.Sprintf("%s/todo/%d", apiRoot, id)
	return sendRequest(u, http.MethodDelete, "", http.StatusNoContent, nil)
}
