package main

import (
	"log"
	todo_app "todo-app"
	"todo-app/pkg/handler"
)

func main(){
	srv := new(todo_app.Server)
	handlers := new(handler.Handler)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil{
		log.Fatalf("error occured while running http server %d", err.Error())
	}
}