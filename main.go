package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Task representa uma tarefa
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Simulando um "banco de dados" em memória
var tasks = []Task{
	{ID: 1, Title: "Aprender Go", Done: false},
	{ID: 2, Title: "Fazer café", Done: true},
}

func main() {
	router := gin.Default()

	// Rotas
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskByID)
	router.POST("/tasks", createTask)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)

	// Starta o servidor
	router.Run(":8080")
}

// Handlers

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func getTaskByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func createTask(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

func updateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, task := range tasks {
		if task.ID == id {
			var updated Task
			if err := c.ShouldBindJSON(&updated); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			updated.ID = id
			tasks[i] = updated
			c.JSON(http.StatusOK, updated)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func deleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
