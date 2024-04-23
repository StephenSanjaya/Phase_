package main

import "Phase3/week1/day2/go-design-pattern/helper"

func main() {
	db := helper.GetInstance().Conn
	defer db.Close()
}
