package main

import (
	"fmt"
	"math"
)

func pembagian(a float64, b float64) (float64) {
	return a / b
}

func pangkat(a float64, b float64) (float64) {
	return math.Pow(a, b)
}

func factorial(a int) (int) {
	res := 1
	for i := 1; i <= a; i++ {
		res *= i
	}
	return res
}

func main() {
	pembagianRes := pembagian(5, 2)
	fmt.Println(pembagianRes)

	pangkatRes := pangkat(10,2)
	fmt.Println(pangkatRes)

	factorialRes := factorial(5)
	fmt.Println(factorialRes)
}