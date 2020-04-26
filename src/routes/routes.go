package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lnwu/todo-api-go/src/controller"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/todos", controller.GetAllTodos)
	router.POST("/todos", controller.CreateTodo)
	router.DELETE("/todo/:todoId", controller.DeleteTodo)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
