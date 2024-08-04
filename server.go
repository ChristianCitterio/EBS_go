package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.POST("/products", postProducts)
	e.PUT("/products/:id", putProducts)
	e.DELETE("/products/:id", deleteProduct)
	e.Logger.Fatal(e.Start(":1323"))
}
