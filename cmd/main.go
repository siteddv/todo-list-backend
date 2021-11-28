package main

import (
	"log"
	todoListBackend "todolistBackend"
	"todolistBackend/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todoListBackend.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
