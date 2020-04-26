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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000/", "http://todo-web-development.s3-ap-northeast-1.amazonaws.com/", "http://todo-web-staging.s3-ap-northeast-1.amazonaws.com/"}
	router.Use(cors.New(config))

	routes.Routes(router)

	log.Fatal(router.Run(":8080"))
}
