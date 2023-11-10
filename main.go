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

// package main

// import (
//     "github.com/gin-gonic/gin"
//     "net/http"
// )

// func main() {
//     // Create a new Gin router
//     router := gin.Default()

//     // Define a route
//     router.GET("/", func(c *gin.Context) {
//         c.JSON(http.StatusOK, gin.H{
//             "message": "Hello, Gin!",
//         })
//     })

//     // Run the server on port 8080
//     router.Run(":8080")
// }