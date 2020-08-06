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

func (ms ModelSuite) Test_Storage_Task_Method() {

	cuuid, err := uuid.FromString("b7f11fc9-478a-4d65-8099-05bbb5799537")
	ms.NoError(err)

	task := Task{
		ID:             cuuid,
		Status:         "DONE",
		CompletionDate: nulls.NewTime(time.Date(2020, time.August, 5, 0, 0, 0, 0, time.UTC)),
		Description:    "this activity",
		ExecutorName:   "Edwin",
		RequesterName:  "Larry",
	}

	err = task.Storage(ms.DB)
	ms.NoError(err)

	stask := Task{}
	err = ms.DB.Find(&stask, cuuid)
	ms.NoError(err)

	ms.Equal(task.Description, stask.Description)
	ms.Equal(task.Status, stask.Status)
	ms.Equal(task.ExecutorName, stask.ExecutorName)
	ms.Equal(task.RequesterName, stask.RequesterName)

	completionDate := task.CompletionDate
	ms.NoError(err)
	sCompletionDate := stask.CompletionDate
	ms.NoError(err)
	ms.Equal(completionDate.Time.Format("01/02/2006"), sCompletionDate.Time.Format("01/02/2006"))

}

func (ms ModelSuite) Test_Storage_Multiple_Tasks_Method() {
	tasks := Tasks{
		{
			Status:         "DONE",
			CompletionDate: nulls.NewTime(time.Date(2020, time.August, 5, 0, 0, 0, 0, time.UTC)),
			Description:    "This activity",
			ExecutorName:   "Edwin",
			RequesterName:  "Larry",
		},
		{
			Status:         "DONE",
			CompletionDate: nulls.NewTime(time.Date(2020, time.August, 5, 0, 0, 0, 0, time.UTC)),
			Description:    "New activity",
			ExecutorName:   "Rodo",
			RequesterName:  "Larry",
		},
	}

	err := tasks.Storage(ms.DB)
	ms.NoError(err)

	stasks := Tasks{}
	ms.NoError(ms.DB.All(&stasks))
	ms.Equal(len(tasks), len(stasks))
}

func (ms ModelSuite) Test_Show_Storage_Tasks() {
	tasks := Tasks{
		{
			Status:         "DONE",
			CompletionDate: nulls.NewTime(time.Date(2020, time.August, 5, 0, 0, 0, 0, time.UTC)),
			Description:    "This activity",
			ExecutorName:   "Edwin",
			RequesterName:  "Larry",
		},
		{
			Status:         "DONE",
			CompletionDate: nulls.NewTime(time.Date(2020, time.August, 5, 0, 0, 0, 0, time.UTC)),
			Description:    "New activity",
			ExecutorName:   "Rodo",
			RequesterName:  "Larry",
		},
	}

	err := tasks.Storage(ms.DB)
	ms.NoError(err)

	task := Task{}
	stasks := task.Tasks(ms.DB)
	ms.Equal(len(tasks), len(stasks))
}
