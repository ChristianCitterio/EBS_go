package main

import (
	"EBS_go/src/routes"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	routes.InitProductRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
