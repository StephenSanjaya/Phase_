package main

import (
	"fmt"
	"math"
)

const (
	RECTANGLE = "rectangle"
	CIRCLE    = "circle"
	TRIANGLE  = "triangle"
)

type Shape struct {
	ShapeType string
	Length    int
	Area      float32
}

func Area(shape Shape, ch chan Shape) {
	var area float32

	switch shape.ShapeType {
	case RECTANGLE:
		area = float32(math.Pow(float64(shape.Length), 2))
	case CIRCLE:
		area = math.Pi * float32(math.Pow(float64(shape.Length), 2))
	case TRIANGLE:
		area = 0.5 * float32(math.Pow(float64(shape.Length), 2))
	}

	ch <- Shape{shape.ShapeType, shape.Length, area}
}

func main() {

	ch := make(chan Shape)
	inputs := []Shape{
		{ShapeType: RECTANGLE, Length: 5},
		{ShapeType: CIRCLE, Length: 3},
		{ShapeType: TRIANGLE, Length: 5},
		{ShapeType: RECTANGLE, Length: 15},
		{ShapeType: CIRCLE, Length: 5},
	}

	outputs := []Shape{}
	for _, input := range inputs {
		go Area(input, ch)
		outputs = append(outputs, <-ch)
	}

	for _, output := range outputs {
		fmt.Printf("\n%v\n", output.ShapeType)
		fmt.Printf("Length\t: %v\n", output.Length)
		fmt.Printf("Area\t: %v\n", output.Area)

	}

}