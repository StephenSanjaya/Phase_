package main

import (
	"Phase3/week1/day1/preview-phase2/config"
	"Phase3/week1/day1/preview-phase2/handler"
	"Phase3/week1/day1/preview-phase2/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.GetConnection()

	authService := handler.NewAuthService(db)
	loanService := handler.NewLoanService(db)

	e := echo.New()

	ms := e.Group("/v1/ms-paylater")
	ms.Use(middleware.AuthMiddleware)
	{
		ms.POST("/register", authService.RegisterHandler)
		ms.POST("/register", authService.LoginHandler)
		ms.POST("/loan", loanService.Loan)
		ms.GET("/limit", loanService.GetLimit)
		ms.POST("/tarik-saldo", loanService.WithdrawBalance)
		ms.POST("/pay", loanService.PayLoanBalance)
	}
}
