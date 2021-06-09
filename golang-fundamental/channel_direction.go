// package main

// import (
// 	"fmt"
// )

// func tampilkan(value chan<- string) {
// 	value <- "belajar channel direction"

// }

// func main() {
// 	value := make(chan string, 3)
// 	tampilkan(value)
// 	fmt.Println(<-value)

// }
package main

import "fmt"

func data1(datas2 chan<- string, msg string) {
	datas2 <- msg
}

func data2(datas2 <-chan string, datas1 chan<- string) {
	msg := <-datas2
	datas1 <- msg
}

func main() {
	datas2 := make(chan string, 1)
	datas1 := make(chan string, 1)
	data1(datas2, "channel direction")
	data2(datas2, datas1)
	fmt.Println(<-datas1)
}
