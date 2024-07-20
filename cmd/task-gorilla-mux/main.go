package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (ts *TaskServer) getTaskHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get task at %s\n", req.URL.Path)

	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	log.Printf("%d", id)

	task, err := ts.store.GetTask(id)
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
	router := mux.NewRouter()
	router.StrictSlash(true)
	server := NewTaskServer()

	router.HandleFunc("/tasks/", server.getAllTaskHandler).Methods("GET")
	router.HandleFunc("/task/{id: [0-9]+}", server.getTaskHandler).Methods("GET")

	log.Println("Server is running!!!")
	log.Fatal(http.ListenAndServe(":8080", router))
}
