package main

import (
	"github.com/labstack/echo/v4"
	"todo-app/internal/app/api"
	"todo-app/internal/app/config"
)

func main(){
	e := echo.New()
	api.MainGroup(e)
	e.Logger.Fatal(e.Start(config.NewConfig().BindAddr))
}