package main

import (
	"Phase3/week2/day2/NGC-5/ms-auth/cmd"
	"Phase3/week2/day2/NGC-5/ms-auth/config"
	"Phase3/week2/day2/NGC-5/ms-auth/controller"
	"Phase3/week2/day2/NGC-5/ms-auth/repository"
)

func main() {
	db := config.GetConnection()

	authRepo := repository.NewAuthRepository(db.User)
	authCtrler := controller.NewAuthController(authRepo)

	cmd.InitGrpc(authCtrler)
}
