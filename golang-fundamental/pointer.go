// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var numberA int = 4
// 	var numberB *int = &numberA

// 	fmt.Println("numberA (value) :", numberA)
// 	fmt.Println("number(value) :", &numberA)
// 	fmt.Println("numberB (value) :", *numberB)
// 	fmt.Println("numberB(value) :", numberB)
// }

package main

import (
	"fmt"
)

func pointer(mypointer *int) {
	*mypointer = 5
}

func main() {
	mypointer := 20
	pointer(&mypointer)
	fmt.Println(&mypointer)

}
