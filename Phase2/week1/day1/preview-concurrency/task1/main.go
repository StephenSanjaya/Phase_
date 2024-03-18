package main

import (
	"fmt"
	"time"
)

func PrintNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func PrintLetters() {
	for i := 97; i <= 106; i++ {
		fmt.Println(string(byte(i)))
	}
}

func main() {

	go PrintLetters()
	go PrintNumbers()

	time.Sleep(2 * time.Second)
}
