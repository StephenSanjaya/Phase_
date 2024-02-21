package main

import "fmt"

type Employee struct {
	name string
}

func main() {
	var emp Employee

	emp.name = "asd"

	fmt.Println(emp)
}