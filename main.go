package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/task", GetTasks)
	app.POST("/task", AddTask)
	app.DELETE("/task/:id", DeleteTask)

	if err := app.Run(":3000"); err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}