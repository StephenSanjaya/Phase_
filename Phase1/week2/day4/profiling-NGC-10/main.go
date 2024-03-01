package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {

	//CPU
	cpuFile, err := os.Create("cpu.pprof")
	if err != nil {
		fmt.Println("error")
		return 
	}
	defer cpuFile.Close()

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		fmt.Println("error")
		return
	}
	defer pprof.StopCPUProfile()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	operand := ""
	angka1 := 0
	angka2 := 0

	fmt.Print("Pilih operasi aritmatika: ")
	fmt.Println("penjumlahan (+) a")
	fmt.Println("pengurangan (-) b")
	fmt.Println("perkalian (*) c")
	fmt.Println("pembagian (/) d")
	_, err = fmt.Scanln(&operand)
	if err != nil {
		panic("input invalid")
	}

	fmt.Print("masukkan angka: ")
	_, err = fmt.Scanln(&angka1)
	if err != nil {
		panic("input invalid")
	}

	fmt.Print("masukkan angka: ")
	_, err = fmt.Scanln(&angka2)
	if err != nil {
		panic("input invalid")
	}

	result := 0
	switch operand {
	case "a":
		result = angka1 + angka2
		fmt.Printf("Hasil penjumlahan %d dan %d adalah %d\n", angka1, angka2, result)
	case "b":
		result = angka1 - angka2
		fmt.Printf("Hasil pengurangan %d dan %d adalah %d\n", angka1, angka2, result)
	case "c":
		result = angka1 * angka2
		fmt.Printf("Hasil perkalian %d dan %d adalah %d\n", angka1, angka2, result)
	case "d":
		result = angka1 / angka2
		fmt.Printf("Hasil pembagian %d dan %d adalah %d\n", angka1, angka2, result)
	}

	//MEMORY
	memFile, err := os.Create("mem.pprof")
	if err != nil {
		fmt.Println("error")
		return 
	}
	defer memFile.Close()

	runtime.GC()
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		fmt.Println("error")
		return
	}
	
}