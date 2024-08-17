package entities_test

import (
	"testing"

	"github.com/jraphaelo/taskfllow-pro/internal/task/domain/entities"
	"github.com/stretchr/testify/require"
)

func TestStatusTask_NewStatusTaks(t *testing.T) {
	t.Parallel()

	t.Run("Test New Status Task", func(t *testing.T) {
		t.Parallel()

		statusTask, err := entities.NewStatusTask(
			"Status Task 1",
			nil,
		)

		require.Nil(t, err)
		require.NotEmpty(t, statusTask)
	})

	t.Run("Test New Status Task with description", func(t *testing.T) {
		t.Parallel()

		description := "Description of status task 1"
		statusTask, err := entities.NewStatusTask(
			"Status Task 1",
			&description,
		)

		require.Nil(t, err)
		require.NotEmpty(t, statusTask)
		require.Equal(t, description, *statusTask.GetDescription())
	})

	t.Run("Test New Status Task with invalid title", func(t *testing.T) {
		t.Parallel()

		_, err := entities.NewStatusTask(
			"",
			nil,
		)

		require.NotNil(t, err)
	})

	t.Run("Test New Status Task with short title", func(t *testing.T) {
		t.Parallel()

		_, err := entities.NewStatusTask(
			"a",
			nil,
		)

		require.NotNil(t, err)
	})
}
