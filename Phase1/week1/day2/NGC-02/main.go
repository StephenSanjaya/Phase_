package main

import (
	"fmt"
	"math/rand"
)

func main() {
	
	// NG Challenge 2 : Conditional 1

	min := 1
	max := 100
	randNum := rand.Intn(max - min) + min

	name := ""
	fmt.Print("Masukkan nama: ")
	fmt.Scanf("%s\n", &name)

	switch {
		case randNum > 80:
			fmt.Printf("Selamat %s, anda sangat beruntung", name)
		case (randNum <= 80) && (randNum > 60):
			fmt.Printf("Selamat %s, anda beruntung", name)
		case (randNum <= 60) && (randNum > 40):
			fmt.Printf("Mohon maaf %s, anda kurang beruntung", name)
		default: 
			fmt.Printf("Mohon maaf %s, anda sial", name)
	}
	fmt.Println()


	// NG Challenge 2 : Conditional 2

	umur := 0

	fmt.Printf("Masukkan nama: ")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("Masukkan umur: ")
	_, err := fmt.Scanf("%d", &umur)

	if err != nil {
		fmt.Println(err)
	}else{
		if(umur < 0 || umur > 100){
			fmt.Println("umur invalid")
		}else if(umur > 18){
			fmt.Println("silahkan masuk")
		}else{
			fmt.Println("dilarang masuk, maksimal umur 19")
		}
	}
}