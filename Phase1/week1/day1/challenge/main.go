package main

import "fmt"

func main() {
	i := 21

	//soal1
	fmt.Print(i, "\n");

	//soal2
	fmt.Printf("%T\n", i);

	//soal3
	fmt.Println("%")

	//soal4
	var j bool =  true
	fmt.Println(j)

	//soal5
	fmt.Println(j)
	
	//soal6
	unicode := "Я"
    fmt.Printf("%+q\n", unicode) // show unicode
	
	//soal7
	num := 21
	num2 := 25
	fmt.Printf("%d\n", num)
	fmt.Printf("%o\n", num2)
	
	//soal8
	huruf := `f`
	fmt.Printf("%x\n", huruf)
	
	//soal9
	huruf2 := `F 13`
	fmt.Printf("%X\n", huruf2)
	
	//soal10
	unicodeR := "Я"
    fmt.Printf("%+q\n", unicodeR) // show unicode
	
	//soal11
	var k float64 = 123.456
	fmt.Printf("%.6f\n", k)

	//soal12
	fmt.Printf("%e\n", k)

}