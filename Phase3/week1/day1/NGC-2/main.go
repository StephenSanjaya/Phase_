package main

import (
	"Phase3/week1/day1/NGC-2/config"
	"Phase3/week1/day1/NGC-2/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	coll := config.GetConnection().Conn
	employeeService := handler.NewEmployeeService(coll)

	e := echo.New()

	e.POST("/employee", employeeService.CreateEmployee)
	e.GET("/all-employee", employeeService.GetAllEmployee)
	e.GET("/employee/:id", employeeService.GetEmployeeById)
	e.PUT("/employee/:id", employeeService.UpdateDataEmployee)
	e.DELETE("/employee/:id", employeeService.DeleteDataEmployee)

	e.Logger.Fatal(e.Start(":8081"))
}
