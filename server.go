package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lnwu/todo-api-go/src/config"
	"github.com/lnwu/todo-api-go/src/routes"
)

func main() {
	config.Connect()

	router := gin.Default()

	router.Use(cors.Default())

	routes.Routes(router)

	log.Fatal(router.Run(":8080"))
}
