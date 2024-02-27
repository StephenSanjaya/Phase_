package main

import (
	"fmt"
	"time"
)

func processImage(url string) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Processing image: https: %s\n", url)
	time.Sleep(1 * time.Second)
	fmt.Printf("Image Processing completed: https: %s\n", url)
}

func main() {
	urls := [][]string{
		{"//example.com/image1.jpg"},
		{"//example.com/image2.jpg"},
		{"//example.com/image3.jpg"},
		{"//example.com/image4.jpg"},
	}

	for _, u := range urls {
		go processImage(u[0])
	}

	fmt.Println("Image processing start main application continues...")

	time.Sleep(3 * time.Second)

	fmt.Println("All image processing completed.")
}