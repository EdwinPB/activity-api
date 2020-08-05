package models

import (
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/uuid"
)

func (ms ModelSuite) Test_Create_Task() {
	task := Task{
		ID:             uuid.Must(uuid.NewV4()),
		Status:         "DONE",
		CompletionDate: nulls.NewTime(time.Now()),
		Description:    "this activity",
		ExecutorName:   "Edwin",
		RequesterName:  "Larry",
	}
	ms.NoError(ms.DB.Create(&task))
}
