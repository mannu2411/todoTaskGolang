package main

import (
	"fmt"

	"github.com/todoTask/database"
	"github.com/todoTask/server"
)

func main() {
	err := database.ConnectAndMigrate("localhost", "5432", "todo", "local", "local", database.SSLModeDisable)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected")
	srv := server.SetupRoutes()
	err = srv.Run(":8080")
	if err != nil {
		panic(err)
	}
}
