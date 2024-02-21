package main

import "fmt"

type Employee struct {
	name    string
	age     int
	address string
}

func (e Employee) Introduce(msg string) string {
	return fmt.Sprintf("%s My name is %s and I'm %d years old", msg, e.name, e.age)
}

func (e Employee) ChangeName1() {
	e.name = "Reinard"
}

func (e *Employee) ChangeName2() {
	e.name = "Reinard"
}

func main() {

	employee := Employee{
		name:    "Agus",
		age:     25,
		address: "Tambun",
	}

	employee.ChangeName1()
	fmt.Println("Ubah dengan method 1", employee.name)
	employee.ChangeName2()
	fmt.Println("Ubah dengan method 2", employee.name)

	// fmt.Println(employee.Introduce("Hello All, "))

	// employee4 := []Employee{
	// 	{
	// 		name:    "Irwan",
	// 		age:     15,
	// 		address: "Bandung",
	// 	},
	// 	{
	// 		name:    "Jawir",
	// 		age:     18,
	// 		address: "Madura",
	// 	},
	// }

	// for _, v := range employee4 {
	// 	fmt.Printf("%+v\n", v)
	// }

	// employee3 := struct {
	// 	name   string
	// 	salary Salary
	// }{}

	// employee3.name = "Raihan"
	// employee3.salary.nominal = 10000000

	// fmt.Println(employee3)

	// employee1 := Employee{
	// 	name:    "Bambang",
	// 	age:     23,
	// 	address: "Bintaro",
	// 	salary: Salary{
	// 		nominal: 5000000,
	// 	},
	// }

	// var employee2 *Employee = &employee1

	// fmt.Printf("value employee1 %+v\n", employee1)
	// fmt.Printf("value employee2 (pointer) %+v\n", employee2)

	// employee2.name = "Jono"
	// fmt.Println("===================")

	// fmt.Printf("value employee1 %+v\n", employee1)
	// fmt.Printf("value employee2 (pointer) %+v\n", employee2)

	// var andrew Employee

	// andrew.name = "Andrew"
	// andrew.age = 20
	// andrew.address = "Jakarta"

	// var budi Employee

	// budi.name = "Budi"
	// budi.age = 20
	// budi.address = "Wakanda"

	// charly := Employee{}
	// charly.name = "Charly"
	// charly.age = 18
	// charly.address = "Kuningan"

	// doni := Employee{
	// 	name:    "Doni",
	// 	age:     40,
	// 	address: "Bogor",
	// }

	// fmt.Println(doni)
	// fmt.Printf("%+v\n", doni)

	// var num int = 4
	// var num2 *int = &num // <-- akan menampilkan alamat memori dari var num

	// fmt.Println("value dari num = ", num)
	// fmt.Println("--------")

	// // fmt.Println("alamat memori dari num2 = ", num2)
	// // fmt.Println("value dari num2 = ", *num2)

	// *num2 = 500
	// fmt.Println("value dari num = ", num)
	// fmt.Println("alamat memori dari num = ", &num)
	// fmt.Println("value dari num2 = ", *num2)
	// fmt.Println("alamat memori dari num2 = ", num2)

	// age := 17
	// fmt.Println(age, "<-- before")
	// addNumber(&age)
	// fmt.Println(age, "<-- after")

}

// func addNumber(num *int) {
// 	*num += 10
// }
