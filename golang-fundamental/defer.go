package main

import (
	"fmt"
)

func main() {

	fmt.Print("Hello, golang!\n")
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("belajar")
}
