package main

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":1234"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello echo")
}
