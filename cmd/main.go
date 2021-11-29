package main

import (
	_ "github.com/lib/pq"
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

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db due to error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
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
