package task_actions

type Status int

const (
	Todo Status = iota
	Progress
	Done
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	// TODO! add createdAt and updatedAt
}

func CreateTask(id int, description string, status Status) Task {
	newTask := Task{
		Id:          id,
		Description: description,
		Status:      status,
	}

	return newTask
}

func (t *Task) UpdateTask(description string) {
	t.Description = description
}

func (t *Task) MarkProgress() {
	t.Status = Progress
}

func (t *Task) MarkDone() {
	t.Status = Done
}
