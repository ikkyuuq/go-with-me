package taskstore

import (
	"errors"
	"time"
)

type TaskStruct struct {
	Due     time.Time `json:"due"`
	Content string    `json:"string"`
	Tags    []string  `json:"tags"`
	Id      int       `json:"id"`
}

type TaskStore struct {
	Tasks []Task
}

type (
	Tasks *[]Task
	Task  *TaskStruct
)

func NewTaskStore() *TaskStore {
	return &TaskStore{
		Tasks: []Task{
			{
				Due:     time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
				Content: "Sample task 1",
				Tags:    []string{"work", "important"},
				Id:      1,
			},
			{
				Due:     time.Date(2022, time.February, 15, 0, 0, 0, 0, time.UTC),
				Content: "Sample task 2",
				Tags:    []string{"home", "urgent"},
				Id:      2,
			},
		},
	}
}

// CreateTask creates a new task in the store.
func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int { return 0 }

// GetTask retrieves a task from the store, by id. If no such id exists, an error is returned.
func (ts *TaskStore) GetTask(id int) (Task, error) {
	for _, task := range ts.Tasks {
		if id == task.Id {
			return task, nil
		}
	}
	return nil, errors.New("not found task")
}

// DeleteTask deletes the task with the given id. If no such id exists, an error
// is returned.
func (ts *TaskStore) DeleteTask(id int) error { return nil }

// DeleteAllTasks deletes all tasks in the store.
func (ts *TaskStore) DeleteAllTasks() error { return nil }

// GetAllTasks returns all the tasks in the store, in arbitrary order.
func (ts *TaskStore) GetAllTasks() ([]Task, error) {
	return ts.Tasks, nil
}

// GetTasksByTag returns all the tasks that have the given tag
func (ts *TaskStore) GetTasksByTag(tag string) Tasks { return nil }

// GetTasksByDueDate returns all the tasks that have the given due date
func (ts *TaskStore) GetTasksByDueDate(year int, month time.Month, day int) []Task { return nil }
