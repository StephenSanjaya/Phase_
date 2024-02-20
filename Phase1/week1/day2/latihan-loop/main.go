package main

import "fmt"

func main()  {
	for i := 1; i <= 15; i++ {
		if i % 3 == 0 && i % 5 == 0{
			fmt.Println("Hello World")
		}else if i % 3 == 0 {
			fmt.Println("Hello")
		}else if i % 5 == 0 {
			fmt.Println("World")
		}else{
			fmt.Println(i)
		}
	}
}