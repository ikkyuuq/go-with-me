package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ikkyuuq/go-with-me/pkg/taskstore"
)

func NewTaskServer() *TaskServer {
	store := taskstore.NewTaskStore()
	return &TaskServer{store: store}
}

type TaskServer struct {
	store *taskstore.TaskStore
}

func (ts *TaskServer) getAllTasksHandler(c *gin.Context) {
	allTasks, _ := ts.store.GetAllTasks()
	c.JSON(http.StatusOK, allTasks)
}

func (ts *TaskServer) getTaskHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	task, err := ts.store.GetTask(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (ts *TaskServer) createTaskHandler(c *gin.Context) {
	type RequestTask struct {
		Due  time.Time `json:"due"`
		Text string    `json:"text"`
		Tags []string  `json:"tags"`
	}

	var rt RequestTask
	if err := c.ShouldBindJSON(&rt); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	id := ts.store.CreateTask(rt.Text, rt.Tags, rt.Due)
	c.JSON(http.StatusCreated, gin.H{"Id": id})
}

func main() {
	// Creates a gin router with default middleware
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	server := NewTaskServer()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Parameter in path
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s", name)
	})

	// However, this one will match /user/:name/*action
	// If no other reouter match /user/:name, it will redirect to /user/:name/
	router.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// Custom handler
	router.GET("/task/", server.getAllTasksHandler)
	router.GET("/task/:id", server.getTaskHandler)
	router.POST("/create/", server.createTaskHandler)

	router.Run() // listen and serve on localhost:8080
}
