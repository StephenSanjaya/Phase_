package main

import (
	"fmt"
	"math"
)

/*
Reference : https://dasarpemrogramangolang.novalagung.com/A-slice.html

*/

/*
Conditional Statement : IF ELSE Condition / Percabangan
IF CONDITION
- IF CONDITION
- IF ELSE CONDITION
- IF ELSE NESTED CONDITION

if condition {
	action
}

Notes :
- condition isinya pasti operator perbandingan (> < >= <= != ==) atau operator logical (&& || !)
- IF hanya ada 1 dalam 1 case dan letaknya selalu paling atas
- ELSE hanya ada 1 dalam 1 case dan letaknya selalu paling bawah
- SEMUA YANG MENGANDUNG KATA IF, WAJIB DITULISKAN CONDITION NYA

SWITCH

switch apa_yang_mau_di_check {
case nilai_pertama:
	action
case nilai_kedua:
	action
default:
	action
}
*/

func main() {
	// menentukan apakah sebuah angka ini negatif, positif, atau netral

	number := 0
	// if number > 0 {
	// 	fmt.Println("Bilangan Positif")

	// 	if number > 5 {
	// 		fmt.Println("Lebih besar dari 5")
	// 	} else {
	// 		fmt.Println("Kurang dari 5")
	// 	}
	// } else if number == 0 {
	// 	fmt.Println("Bilangan Netral")
	// } else {
	// 	fmt.Println("Bilangan Negatif")
	// }

	if number > 0 {
		fmt.Println("Bilangan Positif")
	} else if number == 0 {
		fmt.Println("Bilangan Netral")
	} else {
		fmt.Println("Bilangan Negatif")
	}

	jenjangPendidikan := "Kuliah"

	switch jenjangPendidikan {
	case "TK":
		fmt.Println("ANAK TK")
	case "SD":
		fmt.Println("ANAK SD")
	case "SMP":
		fmt.Println("ANAK SMP")
	case "SMA":
		fmt.Println("ANAK SMA")
	default:
		fmt.Println("NILAI SALAH!")
	}

	/*
		Convert Score to Alphabet

		- Jika value dari score lebih dari 100, maka print Invalid Score
		- Jika value dari score kurang dari 0, maka print Invalid Score
		- Jika value dari score berada diantara 90 hingga 100, print A
		- Jika value dari score berada diantara 80 hingga 89, print B
		- Jika value dari score berada diantara 70 hingga 79, print C
		- Jika value dari score berada diantara 60 hingga 69, print D
		- Jika value dari score berada dibawah 60, print E

	*/
	score := 89

	// Logic pertama
	if score <= 100 && score >= 90 {
		fmt.Println("A")
	} else if score < 90 && score >= 80 {
		fmt.Println("B")
	} else if score < 80 && score >= 70 {
		fmt.Println("C")
	} else if score < 70 && score >= 60 {
		fmt.Println("D")
	} else if score < 60 && score >= 0 {
		fmt.Println("E")
	} else {
		fmt.Println("Invalid Score")
	}

	// Logic kedua
	if score <= 100 && score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else if score < 60 && score >= 0 {
		fmt.Println("E")
	} else {
		fmt.Println("Invalid Score")
	}

	/*
		LOOPING : perulangan
		- FOR LOOPING : kita tau kapan harus berhenti
		for condition 1; condition 2; condition 3 {
			action
		}

		Notes :
		- condition 1 -> inisiasi variable -> i := 0, i := 1, i := 100
		- condition 2 -> condition -> i > 10, i < 100
		- condition 3 -> increment / decrement -> i++, i+=2, i--


		- WHILE LOOPING : kita ga tau kapan harus berhenti
		initiate variable

		for condition {
			action
			increment/decrement
		}

	*/

	// menampilkan angka 1 sampai 5
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// menampilkan angka 1 sampai 5 dengan lompatan 2 angka
	for i := 1; i <= 5; i += 2 {
		fmt.Println(i)
	}

	// bentuk kedua for
	// menampilkan bilangan genap dari 1 sampai 10
	angka := 2        // initiate
	for angka <= 10 { // condition
		fmt.Println(angka) // action
		angka += 2         // increment
	}

	// bentuk ketiga for
	// menampilkan bilangan ganjil dari 1 sampai 10
	angka = 1
	for {
		fmt.Println(angka)

		if angka == 10 {
			break
		}

		angka++
	}

	// break & continue
	// break -> MEMBUAT PERULANGAN BERHENTI
	// continue -> NGE SKIP PART TERTENTU

	// implementasi break
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// implementasi continue
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	/*
		Tampilkan angka 1 sampai 15, dengan ketentuan sbb :
		- Jika angka tersebut merupakan kelipatan 3, print Hello
		- Jika angka tersebut merupakan kelipatan 5, print World
		- Jika angka tersebut merupakan kelipatan 3 DAN kelipatan 5, print Hello World
		- Jika angka tersebut bukan kelipatan 3 atau kelipatan 5, print angka tersebut

		Expected output :
		1 2 Hello 4 World Hello 7 8 Hello World ..... 13 14 Hello World
	*/

	for i := 0; i <= 15; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Hello World")
		} else if i%3 == 0 {
			fmt.Println("Hello")
		} else if i%5 == 0 {
			fmt.Println("World")
		} else {
			fmt.Println(i)
		}
	}

	// taro kondisi yang paling banyak itu diatas

	// SLICE - LIST - ARRAY -> bisa menampung banyak data dalam 1 tempat
	/*
		- Array  =  Nama Hotel
		- Index  =  Nomor Kamar -> SELALU DIMULAI DARI 0
		- Value  =  Penghuni Kamar
	*/

	// names adalah Array (Nama Hotel)
	var names [4]string // artinya array "names" hanya bisa max menampung 4 data
	names[0] = "panji"  // panji adalah value yang nginap di kamar (index) ke 0
	names[1] = "budi"   // budi adalah value yang nginap di kamar (index) ke 1
	names[2] = "anto"
	names[3] = "tini"

	// fruits adalah slice yang bisa menampung data secara unlimited
	var fruits = []string{"apple", "grape", "banana", "melon"}

	// index 				0        1         2         3
	// value			   apple    grape    banana    melon

	fmt.Println(fruits[2]) // akan nge print banana

	/*
		FUNCTION

		func function_name() {
			action
		}

		1. Function with Parameter -> alat bantu / nilai yang bisa berubah ubah
		2. Function with Return -> mengembalikan nilai
			- Jika menggunakan function with return, maka ketika dipanggil, WAJIB DIMASUKKAN KE VARIABLE
		3. Function Variadic -> kita ga tau parameternya itu ada berapa banyak ~
	*/

	sayHello()
	menyapa("panji", 10)
	menyapa("budi", 15)
	menyapa("tomi", 8)

	saldo := 10000
	total := deposit(saldo, 20000)

	fmt.Println(total)

	result, name := addition(10, 20)
	fmt.Println(result, name)

	result = perkalian(1, 2, 3)
	fmt.Println(result)

	result = perkalian(1, 3)
	fmt.Println(result)

	fmt.Println(pangkat(10, 2))

	show(2)

	// factorial(3)
	// fmt.Println(factorial(3))
	// fmt.Println(fact(3))
	fmt.Println(faktorialAngka(3))
}

func sayHello() {
	fmt.Println("Hello dunia")
}

// function with parameter
func menyapa(name string, age int) {
	fmt.Printf("Halo nama saya %s dan saya berumur %d tahun \n", name, age)
}

// function with return
func deposit(saldo int, amount int) (total int) {
	total = saldo + amount
	return total
}

func addition(firstNumber, secondNumber int) (result int, name string) {
	result = firstNumber + secondNumber

	return result, "panji"
}

// function variadic
func perkalian(numbers ...int) (result int) {
	result = 1

	for _, angka := range numbers {
		result = result * angka
	}

	return result
}

/*
Andrew - Azillah :
- Buat function untuk menampilkan 5 angka dengan kelipatan X
Expected :
show(2) -> 1 3 5 7 9

- Stephen, Steven, Yohanes
- Buat function untuk pangkat dengan 2 parameter (angka, pangkat)
Expected :
pangkat(5,2) -> 5 pangkat 2 = 25
*/

func pangkat(a float64, b float64) float64 {
	return math.Pow(a, b)
}

func show(x int) {
	var hasil = 1

	for i := 1; i <= 5; i++ {
		fmt.Println(hasil)
		hasil += x
	}
}

/*
Buatlah function factorial(angka) yang akan menampilkan hasil akhir total dari faktorial tersebut

Expected output :

factorial(5) -> 1 * 2 * 3 * 4 * 5 = ...
*/

// contoh implementasi dari "recursive"
// func factorial(angka int) int {
// 	if angka == 0 {
// 		return 1
// 	}
// 	return angka * factorial(angka-1)
// }

// func factorial(a int) int {
// 	res := 1
// 	for i := 1; i <= a; i++ {
// 		res *= i
// 	}
// 	return res
// }

// func factorial(x int) {
// 	datas := 1
// 	for i := 1; i <= x; i++ {
// 		datas *= i
// 	}

// 	fmt.Println(datas)
// }

// func fact(n int) uint64 {
// 	factorial := uint64(1)
// 	if n < 0 {
// 		fmt.Print("erorr")
// 	} else {
// 		for i := 1; i <= n; i++ {
// 			factorial *= uint64(i)
// 		}
// 	}
// 	return factorial
// }

func faktorialAngka(angka int) int {
	hasil := 1

	for i := 1; i <= angka; i++ {
		hasil *= i
	}
	return hasil
}
