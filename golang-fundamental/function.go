// package main

// import (
// 	"fmt"
// )

// func main() {
// 	hello()

// }
// func hello() {
// 	fmt.Print("Hello, world!")
// }
// package main

// import "fmt"

// func jumlah(a, b int) int {

// 	return a + b
// }

// func multijumlah(a, b, c int) int {
// 	return a + b + c
// }

// func main() {

// 	res := jumlah(1, 2)
// 	fmt.Println("1+2 =", res)

// 	res = multijumlah(3, 4, 5)
// 	fmt.Println("3+4+5 =", res)
// }
package main

import (
	"log"
)

func sum(angka ...int) int {
	total := 0
	for _, val := range angka {
		total += val
	}
	return total
}

func main() {
	log.Println(sum(77, 52, 67, 83, 12))
}
