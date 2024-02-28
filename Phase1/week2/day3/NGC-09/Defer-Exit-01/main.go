package main

import (
	"fmt"
	"math/rand"
)

func Factorial(ch <-chan int) chan int {
	res := make(chan int, 1)
	defer close(res)
	receive := <-ch
	fact := 1

	for i := 1; i <= receive; i++ {
		fact *= i
	}

	res <- fact
	return res
}

func main() {
	// NG Challenge 9 : Defer & Exit 1
	ch := make(chan int, 1)
	defer close(ch)

	n := int(rand.Intn(5) + 1)

	ch <- n

	fmt.Println(<-Factorial(ch))
}