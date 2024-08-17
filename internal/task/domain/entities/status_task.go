package entities

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type StatusTaskInterface interface {
	// StatusTask management methods
	UpdateStatusTask(name string, description *string) error
	ActivateStatusTask() error
	DeactivateStatusTask() error

	// Get methods for StatusTask
	GetID() string
	GetName() string
	GetDescription() *string
	GetIsActived() *bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type StatusTask struct {
	ID          string
	Name        string
	Description *string
	IsActived   *bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewStatusTask(name string, description *string) (*StatusTask, error) {
	statusTask := &StatusTask{
		ID:          uuid.NewV4().String(),
		Name:        name,
		Description: nil,
		IsActived:   nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if description != nil {
		statusTask.Description = description
	}

	statusActive := true
	statusTask.IsActived = &statusActive

	if err := statusTask.validate(); err != nil {
		return nil, err
	}

	return statusTask, nil
}

func (st *StatusTask) UpdateStatusTask(name string, description *string) error {
	st.Name = name

	if description != nil {
		st.Description = description
	}

	st.UpdatedAt = time.Now()

	if err := st.validate(); err != nil {
		return err
	}

	return nil
}

func (st *StatusTask) GetDescription() *string {
	return st.Description
}

func (st *StatusTask) validate() error {
	if st.Name == "" {
		return errors.New("The name of the status task is required")
	}

	if len(st.Name) < 3 {
		return errors.New("The name of the status task needs to have at least 3 characters")
	}

	return nil
}
