package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getem(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
