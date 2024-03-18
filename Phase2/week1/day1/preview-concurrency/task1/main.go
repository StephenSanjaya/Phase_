package main

import (
	"fmt"
	"strconv"
	"sync"
)

func PrintNumbers(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- strconv.Itoa(i)
	}
}

func PrintLetters(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	for i := 97; i <= 106; i++ {
		ch <- string(byte(i))
	}
}

func main() {

	var ch = make(chan string)
	defer close(ch)

	var wg sync.WaitGroup

	wg.Add(2)
	go PrintLetters(&wg, ch)
	go PrintNumbers(&wg, ch)

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg.Wait()
}
