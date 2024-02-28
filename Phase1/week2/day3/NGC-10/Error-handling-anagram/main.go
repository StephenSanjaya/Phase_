package main

import (
	"errors"
	"fmt"
	"regexp"
)

func checkAnagram(w1 string, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}

	b1, b2 := []byte(w1), []byte(w2)

	countB1, countB2 := 0, 0
	for i := 0; i < len(b1); i++ {
		countB1 += int(b1[i])
		countB2 += int(b2[i])
	}

	return countB1 == countB2
}

func checkSymbol(w1 string, w2 string)(err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("input tidak valid")
		}
	}()
	regex, _ := regexp.Compile(`[\(\!\@\#\$\%\^\&\*\(\)\-\+\[\]\{\}\;\:\'\"\<\>\?\/\.\,\) ]`)

	isMatchW1 := regex.MatchString(w1)
	isMatchW2 := regex.MatchString(w2)

	if isMatchW1 || isMatchW2 {
		panic(err.Error())
	}
	return nil
}

func main() {
	word1 := ""
	word2 := ""

	for {
		fmt.Print("Input kata pertama: ")
		_, err := fmt.Scanln(&word1)
		if err != nil || len(word1) > 10 {
			fmt.Println(errors.New("input tidak valid"))
			continue
		}
		break
	}

	for {
		fmt.Print("Input kata kedua: ")
		_, err := fmt.Scanln(&word2)
		if err != nil || len(word2) > 10 {
			fmt.Println(errors.New("input tidak valid"))
			continue
		}
		break
	}

	err := checkSymbol(word1, word2)
	if err != nil {
		fmt.Println(err)
	}else{
		if checkAnagram(word1, word2) {
			fmt.Println("anagram")
		}else{
			fmt.Println("not anagram")
		}
	}
}