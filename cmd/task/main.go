package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ikkyuuq/go-with-me/pkg/taskstore"
)

func NewTaskServer() *TaskServer {
	store := taskstore.NewTaskStore()
	return &TaskServer{store: store}
}

type TaskServer struct {
	store *taskstore.TaskStore
}

func (ts *TaskServer) getAllTaskHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get task at %s\n", req.URL.Path)

	task, err := ts.store.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	mux := http.NewServeMux()
	server := NewTaskServer()

	mux.HandleFunc("GET /task/", server.getAllTaskHandler)

	log.Println("Server is running!!!")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
