package main

import (
	"Phase1/week3/day5/config"
	"Phase1/week3/day5/entity"
	"Phase1/week3/day5/handler"
	"fmt"
)

func main() {

	config.ConnectDB()
	for {
		fmt.Println("Pilih Opsi :")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var email, username, password, dob string
			fmt.Println("Register")
			fmt.Println("Masukan email : ")
			fmt.Scanln(&email)
			fmt.Println("Masukan username : ")
			fmt.Scanln(&username)
			fmt.Println("Masukan password : ")
			fmt.Scanln(&password)
			fmt.Println("Masukan dob : ")
			fmt.Scanln(&dob)

			// define entity
			user := entity.User{
				Email: email,
				Fullname: username,
				Password: password,
				Dob: dob,
			}

			//define handler
			err := handler.Register(user)
			if err != nil {
				fmt.Println("Kesalahan ketika registrasi:", err)
			} else {
				fmt.Println("Registrasi Berhasil!")
			}

		case 2:
			var username, password string
			fmt.Println("Login")
			fmt.Println("Masukan username : ")
			fmt.Scanln(&username)
			fmt.Println("Masukan password : ")
			fmt.Scanln(&password)

			//define handler
			user, isAuthenticated := handler.Login(username, password)
			if isAuthenticated {
				fmt.Println("Login berhasil! selamat datang, ", user.Fullname)
			} else {
				fmt.Println("Gagal Login")
			}
		case 3:
			fmt.Println("keluar dari program")
			return
		default:
			fmt.Println("Pilihan tidak ada")
		}
	}

}
