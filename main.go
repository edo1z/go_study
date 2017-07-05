package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.GET("/:id", helloId)
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func helloId(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	msg := "no member"
	if team != "" || member != "" {
		msg = "team: " + team + ", member: " + member
	}
	return c.String(http.StatusOK, msg)
}
