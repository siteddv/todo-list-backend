package main

import (
	"log"
	todoListBackend "todolistBackend"
	"todolistBackend/pkg/handler"
	"todolistBackend/pkg/repository"
	"todolistBackend/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoListBackend.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
