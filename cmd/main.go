package main

import (
	"github.com/spf13/viper"
	"log"
	todo_app "todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"
)

func main(){
	if err := initConfig(); err != nil{
		log.Fatalf("error initializating configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil{
		log.Fatalf("error occured while running http server %d", err.Error())
	}
}

func initConfig() error{
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}