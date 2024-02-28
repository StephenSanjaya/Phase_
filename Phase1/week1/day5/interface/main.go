package main

import "fmt"

type Number interface {
	int | float32 | float64 | int32
}

func Example[K any, T comparable](name K, age T) bool {
	return true
}

func Sum[T Number](numbers []T) T {
	var result T = 0
	for _, num := range numbers {
		result += num
	}
	return result
}

type Student[T Number] struct {
	Name  string
	Score []T
}

func takeOrders(orders []order) {
	orderName := ""
	qty := 0
	for {
		fmt.Println("What would you like to order? (Type 'done' to finish)")
		_, err := fmt.Scanf("%s\n", &orderName)
		if err != nil {
			fmt.Println(err)
		}

		if orderName == "done" {
			break
		}

		fmt.Printf("How many %s would you like?\n", orderName)
		_, err = fmt.Scanf("%d\n", &qty)
		if err != nil {
			fmt.Println(err)
		}
		o.name = append(o.name, orderName)
	}
}

func (s *Student[T]) AppendScore(newScore T) {
	s.Score = append(s.Score, newScore)
}

func main() {

	captainAmerica := Student[int]{
		Name: "Agus",
	}

	fmt.Println(captainAmerica)
	captainAmerica.AppendScore(90)
	fmt.Println(captainAmerica)
	//fmt.Println(Sum([]int{1, 2, 3, 4, 5, 6}))
	//fmt.Println(Sum([]float32{1.2, 2.2, 3.1, 4.2, 5.1}))
}
