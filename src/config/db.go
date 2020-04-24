package config

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"

	"github.com/lnwu/todo-api-go/src/controller"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "password",
		Addr:     "localhost:5432",
		Database: "postgres",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controller.CreateTodoTable(db)
	controller.InitiateDB(db)
	return db
}
