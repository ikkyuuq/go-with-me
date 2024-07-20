package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func assert(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}

func TestHandler(t *testing.T) {
	server := NewTaskServer()
	t.Run("getAllTaskHandler", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/tasks/", nil)

		record := httptest.NewRecorder()
		handler := http.HandlerFunc(server.getAllTaskHandler)

		handler.ServeHTTP(record, req)

		got := record.Code
		want := http.StatusOK

		assert(t, got, want)
	})
	t.Run("getTaskHandler", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/task/123", nil)

		record := httptest.NewRecorder()
		handler := http.HandlerFunc(server.getTaskHandler)

		handler.ServeHTTP(record, req)

		got := record.Code
		want := http.StatusOK

		assert(t, got, want)
	})
}
