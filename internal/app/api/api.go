package api

import (
	"github.com/labstack/echo/v4"
	"todo-app/internal/app/handler"
)

func MainGroup(e *echo.Echo){
	e.GET("/api/tasks", handler.ShowActions)
}
