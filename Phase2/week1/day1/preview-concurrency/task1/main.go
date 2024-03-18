package main

import (
	"fmt"
	"sync"
)

func PrintNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func PrintLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 97; i <= 106; i++ {
		fmt.Println(string(byte(i)))
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go PrintLetters(&wg)
	go PrintNumbers(&wg)

	wg.Wait()
}
