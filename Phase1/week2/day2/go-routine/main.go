package main

import (
	"fmt"
	"sync"
	// "runtime"
	// "time"
)

func firstProcess(wg *sync.WaitGroup,num int)  {
	fmt.Println("first process")
	for i := 1; i < num; i++ {
		fmt.Println("i = ",i)
	}	
	fmt.Println("first end")
}

// func secondProcess(num int)  {
// 	fmt.Println("second starting")
// 	for i := 1; i < num; i++ {
// 		fmt.Println("j = ",i)
// 	}	
// 	fmt.Println("second end")
// }

func Task(wg *sync.WaitGroup, msg string)  {
	defer wg.Done()
	fmt.Println(msg)
}

func main() {

	// go firstProcess(8)
	// secondProcess(8)

	// fmt.Println("num of goroutine", runtime.NumGoroutine())

	// time.Sleep(2 * time.Second)

	//--------------------------------------------------------------------------
	wg := sync.WaitGroup{}

	// wg.Add(1)
	go firstProcess(&wg,8)
	// wg.Wait()

	for i := 0; i < 20; i++ {
		wg.Add(1)
		var text = fmt.Sprintf("task %d", i)
		go Task(&wg, text)
	}

	wg.Wait()

}