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

func TestStatusTask_UpdateStatusTask(t *testing.T) {
	t.Parallel()

	t.Run("Test Update Status Task", func(t *testing.T) {
		t.Parallel()

		statusTask, _ := entities.NewStatusTask(
			"Status Task 1",
			nil,
		)

		err := statusTask.UpdateStatusTask("Status Task 2", nil)

		require.Nil(t, err)
		require.Equal(t, "Status Task 2", statusTask.Name)
	})

	t.Run("Test Update Status Task with description", func(t *testing.T) {
		t.Parallel()

		statusTask, _ := entities.NewStatusTask(
			"Status Task 1",
			nil,
		)

		description := "Description of status task 2"
		err := statusTask.UpdateStatusTask("Status Task 2", &description)

		require.Nil(t, err)
		require.Equal(t, "Status Task 2", statusTask.Name)
		require.Equal(t, description, *statusTask.GetDescription())
	})

	t.Run("Test Update Status Task with invalid title", func(t *testing.T) {
		t.Parallel()

		statusTask, _ := entities.NewStatusTask(
			"Status Task 1",
			nil,
		)

		err := statusTask.UpdateStatusTask("", nil)

		require.NotNil(t, err)
	})

	t.Run("Test Update Status Task with short title", func(t *testing.T) {
		t.Parallel()

		statusTask, _ := entities.NewStatusTask(
			"Status Task 1",
			nil,
		)

		err := statusTask.UpdateStatusTask("a", nil)

		require.NotNil(t, err)
	})
}

func TestStatusTask_ActivateStatusTask(t *testing.T) {
	t.Parallel()

	t.Run("Test Activate Status Task", func(t *testing.T) {
		t.Parallel()

		statusTask, _ := entities.NewStatusTask(
			"Status Task 1",
			nil,
		)

		err := statusTask.ActivateStatusTask()

		require.Nil(t, err)
		require.NotNil(t, statusTask.IsActived)
		require.True(t, *statusTask.IsActived)
	})
}

func TestStatusTask_DeactivateStatusTask(t *testing.T) {
	t.Parallel()

	t.Run("Test Deactivate Status Task", func(t *testing.T) {
		t.Parallel()

		statusTask, _ := entities.NewStatusTask(
			"Status Task 1",
			nil,
		)

		err := statusTask.DeactivateStatusTask()

		require.Nil(t, err)
		require.NotNil(t, statusTask.IsActived)
		require.False(t, *statusTask.IsActived)
	})
}

func TestStatusTask_GetID(t *testing.T) {
	t.Parallel()

	t.Run("Test Get ID", func(t *testing.T) {
		t.Parallel()

		statusTask, _ := entities.NewStatusTask(
			"Status Task 1",
			nil,
		)

		require.NotEmpty(t, statusTask.GetID())
	})
}

func TestStatusTask_GetName(t *testing.T) {
	t.Parallel()

	t.Run("Test Get Name", func(t *testing.T) {
		t.Parallel()

		statusTask, _ := entities.NewStatusTask(
			"Status Task 1",
			nil,
		)

		require.Equal(t, "Status Task 1", statusTask.GetName())
	})
}

func TestStatusTask_GetDescription(t *testing.T) {
	t.Parallel()

	t.Run("Test Get Description", func(t *testing.T) {
		t.Parallel()

		description := "Description of status task 1"
		statusTask, _ := entities.NewStatusTask(
			"Status Task 1",
			&description,
		)

		require.Equal(t, description, *statusTask.GetDescription())
	})
}

func TestStatusTask_GetIsActived(t *testing.T) {
	t.Parallel()

	t.Run("Test Get Is Actived", func(t *testing.T) {
		t.Parallel()

		statusTask, _ := entities.NewStatusTask(
			"Status Task 1",
			nil,
		)

		require.NotNil(t, statusTask.GetIsActived())
		require.True(t, *statusTask.GetIsActived())
	})
}

func TestStatusTask_GetCreatedAt(t *testing.T) {
	t.Parallel()

	t.Run("Test Get Created At", func(t *testing.T) {
		t.Parallel()

		statusTask, _ := entities.NewStatusTask(
			"Status Task 1",
			nil,
		)

		require.NotEmpty(t, statusTask.GetCreatedAt())
	})
}

func TestStatusTask_GetUpdatedAt(t *testing.T) {
	t.Parallel()

	t.Run("Test Get Updated At", func(t *testing.T) {
		t.Parallel()

		statusTask, _ := entities.NewStatusTask(
			"Status Task 1",
			nil,
		)

		require.NotEmpty(t, statusTask.GetUpdatedAt())
	})
}
