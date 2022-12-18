package cmd

import (
	"github.com/pqppq/pomo/pomodoro"
	"github.com/pqppq/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
