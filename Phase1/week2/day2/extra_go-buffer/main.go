package main

import (
	"fmt"
	"time"
)

func main() {
	channel()
	// bufferChannel()
}

// contoh implementasi channel -> unbuffered channel
func channel() {
	ch := make(chan int) // Membuat unbuffered channel
	// defer close(ch)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Sent:", i)
			//ch <- i // Mengirim nilai ke channel (bloking sampai diterima)
		}
		close(ch)
	}()
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

	//pakai ini untuk loop channel
	for value := range ch {
		fmt.Println("Received:", value)
		time.Sleep(5 * time.Second) // Menyimulasikan proses penerimaan yang memakan waktu
	}
}

func bufferChannel() {
	ch := make(chan int, 10) // Membuat buffered channel dengan kapasitas 10

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // Mengirim nilai ke channel (tidak bloking sampai buffer penuh)
			fmt.Println("Sent:", i)
		}
		close(ch)
	}()

	for value := range ch {
		fmt.Println("Received:", value)
	}
}
