package taskmanager

type task struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

var Tasks []task

func CreateTasks(mySlice *[]task) []task {
	*mySlice = append(Tasks,
		task{ID: 1, Type: "Homework", Description: "This is a test"},
		task{ID: 2, Type: "Maintenance", Description: "Computer Reparation"},
	)

	return *mySlice
}

func UserProvidedTask(id int, taskType string, description string, mySlice *[]task) task {
	createdTask := task{ID: id, Type: taskType, Description: description}
	*mySlice = append(Tasks, createdTask)

	return createdTask
}
