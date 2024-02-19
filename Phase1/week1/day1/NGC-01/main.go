package main

import (
	"fmt"
)

func main() {
	//Variable 1
	fmt.Println("NG Challenge 1 : Variabel 1")
	// A
	var myNum int32 = 50
	fmt.Println("myNum = ", myNum)

	// B
	var myNum2 float32 = 51
	fmt.Println("myNum2 = ", myNum2)

	// C
	var myNumStr string = "50"
	fmt.Println("myNumStr = ", myNumStr)
	fmt.Println()

	//Variable 2
	fmt.Println("NG Challenge 1 : Variabel 2")
	var x int32 = 5
	var y int32 = 10
	var z int32 = x + y
	fmt.Println("nilai z = ",z)
	fmt.Println()

	//CLI 
	fmt.Println("NG Challenge 1 : CLI")
	name := ""
	fmt.Print("Masukkan nama: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello %s\n", name)
	fmt.Println()

	//Array & Slice 1
	fmt.Println("NG Challenge 1 : Array & Slice 1")
	people := []string{"Walt", "Jesse", "Skyler", "Sau"}
	fmt.Println("length: ", len(people));
	people = append(people, "Hank", "Marie")
	fmt.Println("length after adding hank and marie: ", len(people));
	fmt.Println("slice after adding hank and marie: ", people);
	fmt.Println()

	//Array & Slice 2
	fmt.Println("NG Challenge 1 : Array & Slice 2")
	type M map[string]interface{}
	var el []M

	m1 := M{"name": "Hank", "gender": "M"}
    m2 := M{"name": "Heisenberg", "gender": "M"}
	m3 := M{"name": "Skyler", "gender": "F"}

	el = append(el, m1, m2, m3)

	for _, val := range el {
		if str, ok := val["name"].(string); ok {
			if val["gender"] == "M"{
				val["name"] = "Mr. " + str //val["name"].(string)
			}else{
				val["name"] = "Mrs. " + str
			}
		}
	 }	

	fmt.Println(el)
}