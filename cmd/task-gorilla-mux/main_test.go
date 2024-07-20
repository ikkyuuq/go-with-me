package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
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

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(server.getAllTaskHandler)

		handler.ServeHTTP(res, req)

		got := res.Code
		want := http.StatusOK

		assert(t, got, want)
	})
	t.Run("getTaskHandler", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/task/123", nil)

		// try to fake gorilla/mux vars
		vars := map[string]string{
			"id": "123",
		}

		// set req include vars of gorilla/mux
		req = mux.SetURLVars(req, vars)

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(server.getTaskHandler)

		handler.ServeHTTP(res, req)

		got := res.Code
		want := http.StatusOK

		assert(t, got, want)
	})
	t.Run("getTaskHandler empty id", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/task/", nil)

		// fake gorilla/mux vars as empty string
		vars := map[string]string{
			"id": "",
		}

		// set req include vars of gorilla/mux
		req = mux.SetURLVars(req, vars)

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(server.getTaskHandler)

		handler.ServeHTTP(res, req)

		got := res.Code
		want := http.StatusBadRequest

		if got != want {
			t.Errorf("got %#v want %#v", got, want)
		}
	})
	t.Run("getTaskHandler with string", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/task/abc", nil)

		vars := map[string]string{
			"id": "abc",
		}

		req = mux.SetURLVars(req, vars)

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(server.getTaskHandler)

		handler.ServeHTTP(res, req)

		got := res.Code
		want := http.StatusBadRequest

		assert(t, got, want)
	})
}
