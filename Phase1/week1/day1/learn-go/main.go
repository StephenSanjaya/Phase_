package main

import (
	"errors"
	"fmt"
)

func main() {
	var name string = "test"
	var age int = 18

	name = "asdasd"
	age = 20

	fmt.Println("name = ", name + "\nage = ", age)
	// fmt.Printf("%T", age) // %T -> nampilin tipe data
	fmt.Printf("%v\n", age) // %v -> semua tipe data, %d -> int, %s -> string

	//short version
	name2 := "test2"
	fmt.Printf("%T\n", name2)

	var student1, student2, student3 = "satu", "dua", "tiga"
	// fmt.Println(student1, student2, student3)
	//underscore var
	_,_,_ = student1, student2, student3

	// string to char using runes
	str := "Hello GO!"
	// 1. convert string to runes
	runes := []rune(str)
	fmt.Println("convert string to runes: ", runes)
	// 2. convert runes to string
	strFromRunes := string(runes)
	fmt.Println("convert runes to string: ", strFromRunes)
	if str != strFromRunes {
		fmt.Println("strings are not matched")
	} else {
		fmt.Println("strings are matched")
	}

	decimal := 3.434334
	fmt.Printf("%.3f\n", decimal) //3 titik dibelakang koma

	condition := true
	fmt.Printf("%t\n", condition) // %t -> boolean to string

	message := `alllllooo alo alo "allo" 'allloo'`
	fmt.Print(message)

	//error handling
	res, err := process()
	if (err != nil) {
		fmt.Println(res)
	}

}

// Func process(parameter)(type data output){}
func process()(int, error)  {
	return 1, errors.New("ini error")
}