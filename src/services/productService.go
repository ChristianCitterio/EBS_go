package services

import (
	"EBS_go/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetProducts(c echo.Context) error {
	var products []models.Products
	if products == nil || len(products) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Products not found")
	}

	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	var product models.Products
	err := c.Bind(&product)

	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, product)
}

func PostProducts(c echo.Context) error {
	var product = new(models.Products)
	err := c.Bind(&product)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusCreated, product)
}

func PutProducts(c echo.Context) error {
	var product models.Products
	err := c.Bind(&product)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	var product models.Products
	err := c.Bind(&product)
	if err != nil {
		return c.String(http.StatusNotFound, "Not found")
	}

	return c.JSON(http.StatusOK, "Deleted")
}
