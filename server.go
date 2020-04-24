package main

import (
	"github.com/gin-gonic/gin"
)

// Todo comment
type Todo struct {
	Title string `json:"title"`
}

func main() {
	router := gin.Default()

	router.GET("/todos", func(c *gin.Context) {
		todos := [2]Todo{{Title: "title1"}, {Title: "title2"}}

		c.JSON(200, todos)
	})

	router.POST("/todos", func(c *gin.Context) {
		newTodo := Todo{Title: "title"}
		c.JSON(201, newTodo)
	})

	// listen and serve on 0.0.0.0:8080
	router.Run()
}
