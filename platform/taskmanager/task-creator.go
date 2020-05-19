package taskmanager

type task struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func CreateTasks() []task {
	var mySavedTasks []task
	mySavedTasks = append(mySavedTasks,
		task{ID: 1, Type: "Homework", Description: "This is a test"},
		task{ID: 2, Type: "Maintenance", Description: "Computer Reparation"},
	)

	return mySavedTasks
}
