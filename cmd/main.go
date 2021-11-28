package cmd

import (
	"log"
	todoListBackend "todolistBackend"
)

func main() {
	srv := new(todoListBackend.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
