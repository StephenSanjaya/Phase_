package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
	// "strings"
)

func analyzeCsvFile(inputCSV string, outputCSV string) (error) {
	in, err := os.Open(inputCSV)
	if err != nil {
		fmt.Println(errors.New("FILE NOT FOUND"))
	}
	defer in.Close()

	out, err := os.Create(outputCSV)
    if err != nil {
        return err
    }
	defer out.Close()

	read := csv.NewReader(in)
	write := csv.NewWriter(out)
	defer write.Flush()

	records, _ := read.ReadAll()
	var newRow []string
	newRow2 := [][]string{}
	for i := 0; i < len(records); i++ {
		for j := 0; j < 3; j++ {
			row := records[i][j]
			newRow = append(newRow, row)
		}
		newRow2 = append(newRow2, newRow)
		newRow = nil
	}
	for i, v := range newRow2 {
		if i > 0 {
			v[0] = strings.ToUpper(v[0])
			v[2] = "mr." + v[2]
		}
		if err := write.Write(v); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var argsRaw = os.Args
	
	inputFilename := argsRaw[1:][0]
	outputFilename := argsRaw[1:][1]

	err := analyzeCsvFile(inputFilename, outputFilename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}