package taskmanager

type task struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

var tasks []task

func CreateTasks() []task {
	tasks = append(tasks,
		task{ID: 1, Type: "Homework", Description: "This is a test"},
		task{ID: 2, Type: "Maintenance", Description: "Computer Reparation"},
	)

	return tasks
}

func UserProvidedTask(id int, taskType string, description string) task {
	createdTask := task{ID: id, Type: taskType, Description: description}
	tasks = append(tasks, createdTask)

	return createdTask
}
