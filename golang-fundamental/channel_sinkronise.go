// package main

// import (
// 	"fmt"
// 	"time"
// )

// func tugas(selesai chan bool) {
// 	fmt.Println("sedang proses")
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("selesai")
// 	selesai <- true
// }

// func main() {
// 	selesai := make(chan bool, 1)
// 	go tugas(selesai)
// 	<-selesai

// }
package main

import (
	"fmt"
	"time"
)

func tugas1(done chan bool) {
	fmt.Print("tugas 1 (goroutine) running...")
	time.Sleep(2 * time.Second)
	fmt.Println("done")

	done <- true
}

func tugas2() {
	fmt.Println("tugas 2 (goroutine)")
}

func main() {
	done := make(chan bool, 1)
	go tugas1(done)

	if <-done {
		go tugas2()
		fmt.Scanln()
	}
}
