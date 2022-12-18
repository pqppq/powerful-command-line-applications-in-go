package pomodoro_test

import (
	"testing"

	"github.com/pqppq/pomo/pomodoro"
	"github.com/pqppq/pomo/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()

	return repository.NewInMemoryRepo(), func() {}
}
