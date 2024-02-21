package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	name 	string
	age		int
	job		string	
}

func (p Person) GetInfo() string {
	return fmt.Sprintf("Name: %s\nAge: %d\nJob: %s\n", p.name, p.age, p.job)
}

func (p *Person) AddYear() {
	p.age += 1
	if p.age >= 50 {
		p.job = "Retired"
	}
}

type Hero struct {
	name			string
	baseAttack		int
	defence			int
	criticalDamage	int
	healthPoint		int
	weapon			Weapon
}

type Weapon struct {
	attack	int
}

func (h Hero) CountDamage() int {
	critDmg := 0
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	myRand := randomizer.Intn(2)

	if myRand == 1 {
		// fmt.Println(myRand)
		critDmg = 0
	}else{
		// fmt.Println(myRand)
		critDmg = h.criticalDamage
	}
	
	totalDmg := h.baseAttack + h.weapon.attack + critDmg

	return totalDmg
}

func (h *Hero)isAttackedBy(attack Hero) {
	totalDamage := attack.CountDamage() - h.defence

	if totalDamage > 0 {
		h.healthPoint -= totalDamage
	}
}

func Battle(attack Hero, attacked Hero)  {
	attacked.isAttackedBy(attack)

	// fmt.Printf("Attack Total Dmg: %d\n", attack.CountDamage())
	// fmt.Printf("attacked defend : %d\n", attacked.defence)
	fmt.Printf("%s HP: %d\n", attacked.name, attacked.healthPoint)
}

func main() {

	// NG Challenge 5 : Struct & Method 1
	person := Person{
		name: "Bambang", age: 45, job: "Gambler",
	}
	fmt.Println(person.GetInfo())

	person.AddYear()
	person.AddYear()
	person.AddYear()
	person.AddYear()
	person.AddYear()

	fmt.Println(person.GetInfo())

	// NG Challenge 5 : struct & method 2
	person2 := []Person{
		{"Budi", 24, "Tukang Cukur"},
		{"Siti", 54, "Tukang Pukul"},
		{"Tudi", 45, "Tukang Kebun"},
	}

	for _, v := range person2 {
		fmt.Printf("%+v",v.GetInfo())
	}

	// NG Challenge 5 : struct HERO 1
	hero := Hero {
		"Fandi", 70, 25, 35, 250, Weapon{50},
	}

	fmt.Printf("%d", hero.CountDamage())

	// NG Challenge 5 : struct HERO 2
	heroA := Hero {
		"Fendi", 70, 25, 35, 250, Weapon{50},
	}
	heroB := Hero {
		"Rudi", 40, 20, 15, 250, Weapon{30},
	}

	Battle(heroA, heroB)

}