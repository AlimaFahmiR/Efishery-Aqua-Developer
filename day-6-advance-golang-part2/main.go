package main

import (
	"day-6-advance-golang-part2/controller"
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	//inisialisasi echo framework
	e := echo.New()

	//root
	e.GET("/", func(c echo.Context) error {
		result := map[string]string{
			"response_code": "200",
			"message": 			 "Success to connect service",
		}

		return c.JSON(http.StatusOK, result)
	})

	//cust := e.Group("customer")
	//cust.POST("/", controller.GetCustomer)
	e.GET("/customer", controller.GetCustomer)
	// start service echo
	e.Logger.Fatal(e.Start(":5002"))
}