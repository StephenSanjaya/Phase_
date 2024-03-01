package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
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

	ch := make(chan string, 1)

	for i := 0; i < 4; i++ {
		go func() {
			fmt.Println("Mulai mengirim pesan...")
			ch <- "Hello dari buffered channel"
			fmt.Println("Pesan terkirim")
		}()
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

	time.Sleep(3 * time.Second)

}