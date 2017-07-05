package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.GET("/:id", helloId)
	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func helloId(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
