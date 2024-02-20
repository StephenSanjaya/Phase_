package main

import (
	"fmt"
)

func calcAvgTotalMedian(slice []float64) (float64, float64 ,float64) {
	avg := 0.0
	total := 0.0
	median := 0.0

	for i := 0; i < len(slice); i++ {		
		total += slice[i]
		if(i == len(slice)/2){
			median = slice[i]
		}
	}
	avg = total / float64(len(slice))

	return avg, total, median
}

func main() {

	// NG Challenge 3 : Looping 2
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}

	avg, total, median := calcAvgTotalMedian(slice1)
	avg2, total2, median2 := calcAvgTotalMedian(slice2)

	fmt.Println("NG Challenge 3 : Looping 2")
	fmt.Printf("slice1: total = %.2f, avg = %.2f, median = %.2f\n", total, avg, median)
	fmt.Printf("slice2: total = %.2f, avg = %.2f, median = %.2f\n", total2, avg2, median2)
	fmt.Println()

	// NG Challenge 3 : Logic 1 - Palindrome
	myString := "katak"
	myString2 := ""

	for i := len(myString)-1; i >= 0; i-- {
		myString2 += string(myString[i])
	}

	fmt.Println("NG Challenge 3 : Logic 1 - Palindrome")
	fmt.Print(myString + ": ")
	if(myString == myString2){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
	fmt.Println()

	// NG Challenge 3 : Logic 2 - XOXO
	xoxo := "xoxo"
	countX := 0
	countO := 0

	for i := 0; i < len(xoxo); i++ {
		if(string(xoxo[i]) == "x"){
			countX++;
		}else{
			countO++;
		}
	}

	fmt.Println("NG Challenge 3 : Logic 2 - XOXO")
	fmt.Print(xoxo + ": ")
	if(countO == countX){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
	fmt.Println()

	// NG Challenge 3 : Logic 3 - XOXO
	sliceNum := []int {3, 1, 2, 5, 4}
	temp := 0

	for i := 0; i < len(sliceNum)-1; i++ {
		for j := 0; j < len(sliceNum)-i-1; j++ {
			if (sliceNum[j] < sliceNum[j + 1]) {
				temp = sliceNum[j];
				sliceNum[j] = sliceNum[j + 1];
				sliceNum[j + 1] = temp;
			}
		}
	}
	fmt.Println("NG Challenge 3 : Logic 3 - XOXO")
	fmt.Println(sliceNum)
	fmt.Println()

	// NG Challenge 3 : Logic 4 - Asterisk Level 1
	rows1 := 5

	fmt.Println("NG Challenge 3 : Logic 4 - Asterisk Level 1")
	for i := 0; i < rows1; i++ {
		fmt.Println("*")
	}
	fmt.Println()

	// NG Challenge 3 : Logic 5 - Asterisk Level 2
	rows2 := 5

	fmt.Println("NG Challenge 3 : Logic 5 - Asterisk Level 2")
	for i := 0; i < rows2; i++ {
		for j := 0; j < rows2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	// NG Challenge 3 : Logic 6 - Asterisk Level 3
	rows3 := 5

	fmt.Println("NG Challenge 3 : Logic 6 - Asterisk Level 3")
	for i := 0; i < rows3; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	// NG Challenge 3 : Logic 7 - Asterisk Level 4
	rows4 := 5

	fmt.Println("NG Challenge 3 : Logic 7 - Asterisk Level 4")
	for i := 0; i < rows4; i++ {
		for j := rows4-i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
}