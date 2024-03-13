package main

import (
	"Phase1/week4/day2/NGC-17/product"
	"fmt"
)

func main() {
	var prod = product.Product{
		Name:  "Kopi",
		Price: 100,
	}
	p := prod.AddProduct()

	fmt.Println(p)
}