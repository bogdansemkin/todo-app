package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ShowActions(c echo.Context) error {
	return c.String(http.StatusOK, "Show actions")
}

