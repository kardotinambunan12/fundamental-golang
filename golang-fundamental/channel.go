// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func Nilai(nilai chan int) {
// 	hasil := rand.Intn(10)
// 	fmt.Println("nilai hasil random adalam :", hasil)
// 	nilai <- hasil
// }

// func main() {
// 	fmt.Println("sekarang belajar channel")

// 	nilai := make(chan int)
// 	defer close(nilai)

// 	go Nilai(nilai)

// 	hasil := <-nilai
// 	fmt.Println(hasil)
// }
package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "belajar go channel" }()

	msg := <-messages
	fmt.Println(msg)
}
