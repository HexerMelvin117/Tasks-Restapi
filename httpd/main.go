package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hexermelvin117/go-restgin/httpd/handler"
)

func main() {
	r := gin.Default()

	r.GET("/alltasks", handler.GetAllTasks())

	r.Run()
}
