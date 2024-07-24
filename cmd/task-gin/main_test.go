package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ikkyuuq/go-with-me/pkg/taskstore"
)

func assertJSON(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Failed to mashal task: %v", err)
		return
	}
}

func assertRequest(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
		return
	}
}

func assertResponeCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %#v want %#v", got, want)
		return
	}
}

func TestCreateTaskHandler(t *testing.T) {
	t.Run("with full json input", func(t *testing.T) {
		// Set Gin to "TestMode"
		gin.SetMode(gin.TestMode)

		router := gin.Default()
		server := NewTaskServer()
		router.POST("/create/", server.createTaskHandler)

		// Create instance of TaskStore
		task := taskstore.Task{
			Due:     time.Now(),
			Content: "This is just a test task",
			Tags:    []string{"test", "test2"},
		}

		// Parse task into json using json.Marshal
		body, err := json.Marshal(task)
		assertJSON(t, err)

		// Make NewRequest with method POST with url '/create/'
		// bytes.Buffer(body) to pass a JSON body in a POST method
		// when using json.Marshal it return a byte slice ([]byte)
		// http request in Go expect the body to be an io.Reader
		// so bytes.NewBuffer(body) converts the []byte into a *byte.Buffer,
		// which implements 'io.'
		req, err := http.NewRequest(http.MethodPost, "/create/", bytes.NewBuffer(body))
		assertRequest(t, err)

		req.Header.Set("Content-type", "application/json")

		// Use httptest.NewRecorder() to record the response
		record := httptest.NewRecorder()
		// Serve the HTTP request using the Gin router
		router.ServeHTTP(record, req)

		got := record.Code
		want := http.StatusCreated

		assertResponeCode(t, got, want)
	})
}
