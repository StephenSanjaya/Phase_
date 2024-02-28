package main

import (
	"fmt"
	"math"
)

func SumSquare(num int, result chan int) {
	total := 0
	for i := 1; i <= num; i++ {
		total += i
	}
	total = int(math.Pow(float64(total),2))
	result <- total
}

func SquareSum(num int, result chan int) {
	total := 0
	for i := 1; i <= num; i++ {
		total += int(math.Pow(float64(i),2))
	}
	result <- total
}

func main() {
	
	resultSumSquare := make(chan int)
	resultSquareSum := make(chan int)

	go SumSquare(100, resultSumSquare)
	go SquareSum(100, resultSquareSum)

	fmt.Printf("Semua angka dijumlahkan lalu dikuadratkan: %v\n", <-resultSumSquare)
	fmt.Printf("Semua angka dikuadratkan lalu dijumlahkan: %v\n", <-resultSquareSum)

}