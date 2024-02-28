package main

import (
	"fmt"
)

type Employee interface {
	Language() string
	Age() int
}

type Engineer struct {
	name string
}

func (e Engineer) Language() string {
	return e.name + " programs in Go"
}

func (e Engineer) Age() int {
	return 24
}

// NG Challenge 6 : airport robot 2
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type German struct{}
type Italian struct{}
type Portuguese struct{}

func SayHello(name string, g Greeter) string {
	return "I can speak " + g.LanguageName() + " : " + g.Greet(name)
}

func (g German) Greet(name string) string {
	return "Hallo " + name + " !"
}

func (g German) LanguageName() string {
	return "German"
}

func (p Portuguese) Greet(name string) string {
	return "Olá " + name + " !"
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (i Italian) Greet(name string) string {
	return "Ciao  " + name + " !"
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func main() {
	// NG Challenge 6 : Interface 1
	var programmer []Employee

	elliot := Engineer{name: "Elliot"}

	programmer = append(programmer, elliot)

	fmt.Println(programmer)
	fmt.Println(elliot.Language())
	fmt.Println(elliot.Age())

	// NG Challenge 6 : airport robot 2
	germanGreeter := German{}
	italianGreeter := Italian{}
	portugueseGreeter := Portuguese{}

	fmt.Println(SayHello("Dietrich", germanGreeter))
	fmt.Println(SayHello("Mario", italianGreeter))
	fmt.Println(SayHello("Ronaldo", portugueseGreeter))
}
