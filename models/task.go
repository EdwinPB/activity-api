package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

type Task struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	Description    string     `json:"description" db:"description"`
	Status         string     `json:"status" db:"status"`
	CompletionDate nulls.Time `json:"completion_date" db:"completion_date"`
	RequesterName  string     `json:"requester_name" db:"requester_name"`
	ExecutorName   string     `json:"executor_name" db:"executor_name"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t Task) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Tasks is not required by pop and may be deleted
type Tasks []Task

// String is not required by pop and may be deleted
func (t Tasks) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Task) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Description, Name: "Description"},
		&validators.StringIsPresent{Field: t.Status, Name: "Status"},
		&validators.StringIsPresent{Field: t.RequesterName, Name: "RequesterName"},
		&validators.StringIsPresent{Field: t.ExecutorName, Name: "ExecutorName"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Task) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Task) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (t Task) Storage(tx *pop.Connection) error {
	if err := tx.Create(&t); err != nil {
		return err
	}
	return nil
}

func (ts Tasks) Storage(tx *pop.Connection) error {
	if err := tx.Create(&ts); err != nil {
		return err
	}
	return nil
}
