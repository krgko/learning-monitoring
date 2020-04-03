package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	s := &http.Server{
		Addr:         ":8055",
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))
}

func hello(c echo.Context) error {
	response := map[string]string{}
	response["message"] = "hello world"
	return c.JSON(http.StatusOK, response)
}
