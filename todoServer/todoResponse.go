package main

import (
	"encoding/json"
	"time"

	"github.com/pqppq/todo"
)

type todoResponse struct {
	Results todo.List `json:"results"`
}

func (r *todoResponse) MarshalJSON() ([]byte, error) {
	resp := struct {
		Results       todo.List `json:"results"`
		Date          int64     `json:"date"`
		TotaolResults int       `json:"total_results"`
	}{
		Results:       r.Results,
		Date:          time.Now().Unix(),
		TotaolResults: len(r.Results),
	}

	return json.Marshal(resp)
}
