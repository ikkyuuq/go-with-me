package taskstore

import (
	"reflect"
	"testing"
	"time"
)

func TestGetAllTask(t *testing.T) {
	t.Run("Get all tasks", func(t *testing.T) {
		taskstore := NewTaskStore()
		got, _ := taskstore.GetAllTasks()
		want := []Task{
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
		}
		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want []Task) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but want %v", got, want)
	}
}
