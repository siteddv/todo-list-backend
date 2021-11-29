package main

import (
	"github.com/spf13/viper"
	"log"
	todoListBackend "todolistBackend"
	"todolistBackend/pkg/handler"
	"todolistBackend/pkg/repository"
	"todolistBackend/pkg/service"
)

const (
	configPath = "configs"
	configName = "config"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoListBackend.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}
