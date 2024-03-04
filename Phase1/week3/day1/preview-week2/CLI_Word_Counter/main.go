package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func analyzeTextFile(filename string, wordCounter *map[string]int) (error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(errors.New("FILE NOT FOUND"))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		go func ()  {
			arr := strings.Split(scanner.Text(), " ")
			for _, v := range arr{
				(*wordCounter)[v]++
			}
		}()
		time.Sleep(1 * time.Second)
	}

	return nil
}

func main() {
	//1. CLI Word Counter
	wordCounter := map[string]int{}
	var argsRaw = os.Args

	filename := argsRaw[1:][0]

	err := analyzeTextFile(filename, &wordCounter)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Word counter : ")
	for v := range wordCounter {
		fmt.Println(v, ": ", wordCounter[v])
	}
}
