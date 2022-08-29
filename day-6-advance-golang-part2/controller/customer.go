package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type Customer struct {
	ID int
	Name string
	Email string
}

var CustomerList = make(map[int]Customer)

func GetCustomer(c echo.Context) (err error) {
	var result []Customer

	dataDummyCustomer := Customer {
		ID: 1,
		Name: "Raka",
		Email: "raka@gmail.com",
	}

	result = append(result, dataDummyCustomer)

	for key, _ := range CustomerList {
		tempCustomer := CustomerList[key]

		result = append(result, tempCustomer)
	}

	return c.JSON(http.StatusOK, result)
}