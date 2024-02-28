package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func analyzeLogFile(filename string, l map[string] int) (error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(errors.New("FILE NOT FOUND"))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "INFO") {
			l["INFO"]++
		}else if strings.Contains(scanner.Text(), "DEBUG") {
			l["DEBUG"]++
		}else if strings.Contains(scanner.Text(), "ERROR") {
			l["ERROR"]++
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	//Scenario : CLI Log Analyzer
	log_level_counter := map[string]int{
		"INFO":  0,
		"ERROR": 0,
		"DEBUG": 0,
	}

	filename := "log.txt"
	err := analyzeLogFile(filename, log_level_counter)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Log analysis result: ")
	fmt.Println()
	for key, v := range log_level_counter {
		fmt.Printf("[%s] level: %d occurences\n", key, v)
	}

}