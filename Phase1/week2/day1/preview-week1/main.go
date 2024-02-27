package main

import (
	"fmt"
	// "os"
	// "strconv"
)

func checkPrime(num int) bool {
	for i := 2; i < num/2; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func printNumbers(str string, num int)  {
	fmt.Printf("%s numbers upto your input are: \n", str)
	a := 0
	if str == "Even" {
		a = 0
	}else{
		a = 1
	}
	for i := 1; i <= num; i++ {
		if i % 2 == a{
			fmt.Println(i)
		}
	}
	if checkPrime(num) {
		fmt.Println("The number is a prime number.")
	}else{
		fmt.Println("The number is not a prime number.")
	}
}

func main() {
	// Instructions #1
	// var argsRaw = os.Args
	// var args = argsRaw[1:]
	// idx, _ := strconv.Atoi(args[0])

	// data := [][]string{
	// 	{"Budi", "Jln. Budi123", "Pekerjaan Budi"},
	// 	{"Rudi", "Jln. Rudi123", "Pekerjaan Rudi"},
	// 	{"Siti", "Jln. Siti123", "Pekerjaan Siti"},
	// }
	// fmt.Printf("Data teman dengan absen %d\n", idx)
	// fmt.Printf("Nama: %s\n", data[idx-1][0])
	// fmt.Printf("Alamat: %s\n", data[idx-1][1])
	// fmt.Printf("Pekerjaan: %s\n", data[idx-1][2])

	// Instructions #2
	// var argsRaw = os.Args
	// var args = argsRaw[1:]
	// result := 0

	// firstNum, _ := strconv.Atoi(args[1])
	// secondNum, _ := strconv.Atoi(args[2])

	// if args[0] == "add" {
	// 	result = firstNum + secondNum
	// }else if args[0] == "sub" {
	// 	result = firstNum - secondNum
	// }else if args[0] == "mul" {
	// 	result = firstNum * secondNum
	// }else if args[0] == "div" {
	// 	result = firstNum / secondNum
	// }

	// fmt.Printf("Result = %.2f", float64(result))

	// Instructions #3
	num:=0
	str:=""

	fmt.Print("Please enter a number: ")
	fmt.Scanln(&num)

	if num % 2 == 0 {
		fmt.Println("The number is even")
		str = "Even"
		printNumbers(str, num)
	}else{
		fmt.Println("The number is odd")
		str = "Odd"
		printNumbers(str, num)
	}

}