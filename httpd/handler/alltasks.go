package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hexermelvin117/go-restgin/platform/taskmanager"
)

func GetAllTasks() gin.HandlerFunc {
	myTasks := taskmanager.CreateTasks()
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, myTasks)
	}
}
