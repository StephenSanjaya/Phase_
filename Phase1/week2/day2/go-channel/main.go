package main

import (
	"fmt"
	"time"
)

// type Result struct {
// 	Operation string
// 	Value     float64
// }

// func penjumlahan(chInput chan [2]float64, chOutput chan Result) {
// 	data := <-chInput
// 	// chOutput <- data[0] + data[1]
// 	chOutput <- Result{"penjumlahan", data[0] + data[1]}
// }

// func pengurangan(chInput chan [2]float64, chOutput chan Result) {
// 	data := <-chInput
// 	chOutput <- Result{"pengurangan", data[0] - data[1]}
// }
// func perkalian(chInput chan [2]float64, chOutput chan Result) {
// 	data := <-chInput
// 	chOutput <- Result{"perkalian", data[0] * data[1]}
// }
// func pembagian(chInput chan [2]float64, chOutput chan Result) {
// 	data := <-chInput
// 	chOutput <- Result{"pembagian", data[0] / data[1]}
// }
func main() {

	ch := make(chan string, 1)

	for i := 0; i < 4; i++ {
		go func() {
			fmt.Println("Mulai mengirim pesan...")
			ch <- "Hello dari buffered channel"
			fmt.Println("Pesan terkirim")
		}()
	}

	time.Sleep(2 * time.Second)

	// go func() {
	// 	message := <-ch
	// 	fmt.Println("Pesan diterima", message)
	// }()

	time.Sleep(3 * time.Second)
	// chPenjumlahan := make(chan [2]float64)
	// chPengurangan := make(chan [2]float64)
	// chPerkalian := make(chan [2]float64)
	// chPembagian := make(chan [2]float64)
	// chOutput := make(chan Result)

	// go penjumlahan(chPenjumlahan, chOutput)
	// go pengurangan(chPengurangan, chOutput)
	// go perkalian(chPerkalian, chOutput)
	// go pembagian(chPembagian, chOutput)

	// //set nilai kedalam channel
	// chPenjumlahan <- [2]float64{10, 5}
	// chPengurangan <- [2]float64{10, 5}
	// chPerkalian <- [2]float64{10, 5}
	// chPembagian <- [2]float64{10, 5}

	// for i := 0; i < 4; i++ {
	// 	result := <-chOutput
	// 	fmt.Println("Hasil", result.Operation, ":", result.Value)
	// }

	// fmt.Println("hasil penjumlahan: ", <-chOutput)
	// fmt.Println("hasil pengurangan: ", <-chOutput)
	// fmt.Println("hasil perkalian: ", <-chOutput)
	// fmt.Println("hasil pembagian: ", <-chOutput)

}

// 1. Membuat channel : make function
//     ch := make(chan int)
// 2. Mengirim = operator (variable channel <- value)
//     ch <- 10
// 3. Menerima = operator (value := <-ch)
// 4. Buffered Channel : punya kapasitas,
//    pengiriman data hanya akan terblokir jika kapasitas penuh
//    penerimaan data hanya akan terblokir jika channel kosong
//    ch := make(chan int, 5)
// 5. Unbuffered Channel : tidak punya kapasitas
//    pengiriman data akan terblokir sampai ada go routine lain yang menerima datanya
// 6. Synchronize
// 7. Closing Channel : fungsi close()
// 8. Range pada channel : for dengan range
