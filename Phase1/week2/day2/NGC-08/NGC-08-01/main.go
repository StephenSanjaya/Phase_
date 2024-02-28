package main

import (
	"fmt"
)

func main() {
	ch := make(chan [99]int, 1)
	defer close(ch)

	var x [99]int

	for i := 1; i <= 99; i++ {
		x[i-1] = i
	}
		
	ch <- x

	total := 0
	totalGanjil := 0
	totalGenap := 0
	
	for _, v := range <-ch {
		if v % 3 == 0 && v % 5 == 0 {
			fmt.Println("15FizzBuzz")
			total += v
		}else if v % 3 == 0 {
			fmt.Println("3Fizz")
			total += v
		}else if v % 5 == 0 {
			fmt.Println("5Buzz")
			total += v
		}	

		if v % 2 == 0 {
			totalGenap += v
		}else{
			totalGanjil += v
		}
	}

	fmt.Println()
	fmt.Println("total: ", total)
	fmt.Println("total ganjil: ", totalGanjil)
	fmt.Println("total genap: ", totalGenap)

}