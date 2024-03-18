package main

import (
	"fmt"
	"sync"
)

func EvenOdd(chEven, chOdd, chErr chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 22; i++ {
		if i > 20 {
			chErr <- i
		} else if i%2 == 0 {
			chEven <- i
		} else if i%2 == 1 {
			chOdd <- i
		}

	}
}

func main() {
	chEven := make(chan int)
	chOdd := make(chan int)
	chErr := make(chan int)
	// defer close(chEven)
	// defer close(chOdd)

	var wg sync.WaitGroup

	wg.Add(1)
	go EvenOdd(chEven, chOdd, chErr, &wg)

	go func() {
		// time.Sleep(2 * time.Second)
		wg.Wait()
		close(chEven)
		close(chOdd)
		close(chErr)
	}()

	for i := 1; i <= 22; i++ {
		select {
		case even := <-chEven:
			fmt.Println("Receive an even number:", even)
		case odd := <-chOdd:
			fmt.Println("Receive an odd number:", odd)
		case err := <-chErr:
			fmt.Printf("Error: number %d is greater than 20\n", err)
		}
	}

	// wg.Wait()

}
