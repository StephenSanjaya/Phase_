package main

import (
	"fmt"
	"math"
)

func hitungBonus(performa string, gaji float64) float64 {
	bonus := 0.0
	switch performa {
		case "A":
			bonus = gaji * 0.2
		case "B":
			bonus = gaji * 0.1
		case "C":
			bonus = gaji * 0.05
		case "D":
			bonus = 0
	}
	return bonus
}

func evaluasiKinerjaSiswa(nilaiSiswa []int)  {
	for i, v := range nilaiSiswa {
		if v >= 85 {
			fmt.Printf("Siswa %d mendapatkan predikat A\n", i+1)
		}else if v >= 70 {
			fmt.Printf("Siswa %d mendapatkan predikat B\n", i+1)
		}else if v >= 50 {
			fmt.Printf("Siswa %d mendapatkan predikat C\n", i+1)
		}else{
			fmt.Printf("Siswa %d mendapatkan predikat D\n", i+1)
		}
	}
}

func main() {

	//soal 1
	n := 5.0
	flag := 0

	r := n/2.0
	space := int(math.Floor(r))
	star := 1

	//5 -> space: 2
	//7 -> space: 3
	//9 -> space: 4

	for i := 1; i <= int(n); i++ {
		for j := 1; j <= space; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= star; j++ {
			fmt.Print("*")
		}
		if(int(n) > star && flag == 0){
			star += 2
			space--
		}else{
			flag = 1
			star -= 2
			space++
		}
		fmt.Print("\n")
	}

	//soal 1.5
	name := "IVAN SETIAWAN"
	row := 5
	k := 0

	for i := 1; i <= row; i++ {
		for j := 1; j <= i; j++ {
			if k < len(name) {
				fmt.Print(string(name[k]))
			}else{
				fmt.Print("*")
			}
			k++;
		}
		fmt.Println()
	}

	// soal 2
	fmt.Printf("Bonus: %.2f\n",  hitungBonus("A", 1000000))

	//soal 3
	nilaiSiswa := []int{85, 60, 78, 92, 45, 73}
	evaluasiKinerjaSiswa(nilaiSiswa)
	
}