package main

import (
	"fmt"
	"time"
)

func value1(value chan int) {
	value <- 2

}
func value2(value chan int) {
	value <- 3

}

func main() {
	Nvalue1 := make(chan int)
	Nvalue2 := make(chan int)

	go value1(Nvalue1)
	go value2(Nvalue2)

	// go func() {
	// 	value1(Nvalue1)

	// }()
	// go func() {
	// 	value2(Nvalue2)
	// }()
	time.Sleep(time.Second)

	for i := 0; i < 2; i++ {
		select {
		case cv1 := <-Nvalue1:
			fmt.Println(cv1)
		case cv2 := <-Nvalue2:
			fmt.Println(cv2)
		}
	}

}
