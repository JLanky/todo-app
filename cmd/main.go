package main

import (
	"log"

	"github.com/JLanky/todo-app"
	"github.com/JLanky/todo-app/pkg/handler"
	"github.com/JLanky/todo-app/pkg/repository"
	"github.com/JLanky/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)
	srv := todo.New()

	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
