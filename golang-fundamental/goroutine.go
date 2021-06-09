package main

import (
	"fmt"
	"time"
)

func message(dari string) {
	for i := 0; i < 3; i++ {
		fmt.Println(dari, ":", i)
	}
}

func main() {

	message("nilai")

	go message("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("lanjut")

	time.Sleep(time.Second)
	fmt.Println("selesai")
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func say() {
// 	fmt.Println("Halo ini berasal dari goroutine baru")
// }
// func main() {
// 	go say()
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Ini Fungsi Utama")
// }
