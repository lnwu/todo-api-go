package main

import "github.com/gin-gonic/gin"

// Todo comment
type Todo struct {
	Title string `json:"title"`
}

func main() {
	r := gin.Default()

	r.GET("/todos", func(c *gin.Context) {
		todos := [1]Todo{{Title: "title1"}}

		c.JSON(200, todos)
	})

	r.POST("/todos", func(c *gin.Context) {
		newTodo := Todo{Title: "title"}
		c.JSON(201, newTodo)
	})

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
