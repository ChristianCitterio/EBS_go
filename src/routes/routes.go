package routes

import (
	"EBS_go/src/services"
	"github.com/labstack/echo/v4"
)

func InitProductRoutes(e *echo.Echo) {
	e.GET("/products", services.GetProducts)
	e.GET("/products/:id", services.GetProduct)
	e.POST("/products", services.PostProducts)
	e.PUT("/products/:id", services.PutProducts)
	e.DELETE("/products/:id", services.DeleteProduct)
}
