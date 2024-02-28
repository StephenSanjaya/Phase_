package main

/*
References :
- https://gosamples.dev/exit/#:~:text=To%20exit%20an%20application%20in,the%20non%2Dzero%20an%20error.
- https://tldp.org/LDP/abs/html/exitcodes.html

*/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {

	// channel()
	// deferGolang()

	// fmt.Println("hello world")
	// os.Exit(1) // memberhitikan program, sehingga code di bawah baris ini tidak akan dieksekusi
	// fmt.Println("Hi semuanya")

	// implementationDeferExit()

	/*
		filename := "data.txt"
		lines, err := processFile(filename)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Number of lines in %s: %d\n", filename, lines)
	*/

	// errorHandling()

	result, err := divide(10, 5)
	fmt.Println(result, err)

	result, err = divide(10, 0)
	fmt.Println(result, err)

}

// contoh implementasi goroutine dan channel
func channel() {
	// membuat channel untuk mengirim data bertipe int
	messages := make(chan int)

	// membuat goroutine
	go func() {
		messages <- 42
	}()

	// menerima data dari channel, lalu di print
	msg := <-messages
	fmt.Println(msg)

	go func(message string) {
		fmt.Println(message)
	}("testong")

	time.Sleep(time.Second) // memberikan delay 1 sec
	fmt.Println("DONE")
}

// contoh implementasi defer
func deferGolang() {
	// ini akan dieksekusi setelah semua code pada func main() telah selesai
	defer fmt.Println("Hola hola 1")
	fmt.Println("Hello world 1")

	defer fmt.Println("Hola hola 2")
	fmt.Println("Hello world 2")

	defer fmt.Println("Hola hola 3")
	fmt.Println("Hello world 3")

	// Implementasi defer dalam proses pembuatan file
	fmt.Println("Mulai ...")
	file, err := os.Create("testong.txt")
	if err != nil {
		fmt.Println("Gagal membuat file :", err)
		return
	}

	defer file.Close()
	fmt.Println("File berhasil dibuat")
}

func implementationDeferExit() {
	// melakukan defer terhadap function cleanup
	defer cleanup()

	fmt.Println("Doing some critical work")

	if !true {
		fmt.Println("Ada error nih!")
		os.Exit(1)
	}

	fmt.Println("Critical work done successfully")
}

func cleanup() {
	fmt.Println("Cleanup : Program already clean")
}

// implementasi defer tambahan
func processFile(filename string) (int, error) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close() // implementasi defer

	// membaca isi file
	scanner := bufio.NewScanner(file)

	// menghitung ada berapa baris (line) pada file txt
	lines := 0
	for scanner.Scan() {
		lines++
	}

	// jika ada error, maka return error
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// jika tidak ada error, maka return lines tanpa ada error
	return lines, nil
}

func errorHandling() {
	// implementasi error biasa
	file, err := os.Open("abc.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() // implementasi defer

	// implementasi custom error -> errors.New()
	file, err = os.Open("abc.txt")
	if err != nil {
		fmt.Println(errors.New("FILE TIDAK ADA"))
	}
	defer file.Close() // implementasi defer
}

// Implementasi recover() -> recover berfungsi untuk menangkap panic dan mengubah hasil error nya
func divide(a, b int) (result int, err error) {
	// A deferred function is declared. This will only run when the surrounding function (divide) finishes.
	defer func() {
		// recover() is used to catch a panic, preventing the panic from crashing the program.
		if r := recover(); r != nil {
			// If a panic is caught, it will be converted into an error.
			// This allows the function to return an error instead of causing a crash.
			err = fmt.Errorf("ada panic nih. ga boleh dibagi 0")
		}
	}()

	if b == 0 {
		panic("Division by zero")
	}

	return a / b, nil
}
