package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Area          float64
	Circumference float64
}

func luas_keliling(l chan<- Circle, diameter []float64) {
	for _, d := range diameter {
		r := d / 2
		luas := math.Pi * r * r
		keliling := 2 * math.Pi * r

		l <- Circle{luas, keliling}	
	}
}

func main() {
	// NG Challenge 9 : Defer & Exit 2
	chCircle := make(chan Circle)
	defer close(chCircle)
	
	diameter := []float64{10.0, 8.0, 5.0}

	go luas_keliling(chCircle, diameter)	

	fmt.Println(<-chCircle)
	fmt.Println(<-chCircle)
	fmt.Println(<-chCircle)
}