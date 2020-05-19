package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hexermelvin117/go-restgin/platform/taskmanager"
)

type createTaskRequest struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func GetAllTasks() gin.HandlerFunc {
	myTasks := taskmanager.CreateTasks()
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, myTasks)
	}
}

func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := createTaskRequest{}
		c.Bind(&requestBody)
		taskmanager.UserProvidedTask(requestBody.ID, requestBody.Type, requestBody.Description)

		c.JSON(200, gin.H{
			"result": requestBody,
		})
	}
}
