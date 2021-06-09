package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 1)

	messages <- "belajar golang"
	messages <- "gelajar buffered golang"
	messages <- "gelajar buffered golang"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
