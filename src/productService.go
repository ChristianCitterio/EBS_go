package src

import (
	"EBS_go/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func getProducts(c echo.Context) {
	var products []models.Products
	err := c.Bind(&products)
	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) {
	var product models.Products
	err := c.Bind(&product)
	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, product)
}

func postProducts(c echo.Context) {
	var product = new(models.Products)
	err := c.Bind(&product)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusCreated, product)
}

func putProducts(c echo.Context) {
	var product models.Products
	err := c.Bind(&product)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) {
	var product models.Products
	err := c.Bind(&product)
	if err != nil {
		return c.String(http.StatusNotFound, "Not found")
	}

	return c.JSON(http.StatusOK, "Deleted")
}
