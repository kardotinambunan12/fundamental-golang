// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var Slice = []string{"apple", "maggo", "banana", "pineapple", "melon"}
// 	fmt.Println(Slice[2])

// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var Slice = []string{"apple", "maggo", "banana", "pineapple", "melon"}

// 	fmt.Println(Slice[0:2])
// }

//penggunaan Len untuk menghitung panjang slice

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var Slice = []string{"apple", "maggo", "banana", "pineapple", "melon"}
// 	fmt.Println(len(Slice))

// }

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	var Slice = []string{"apple", "maggo", "banana", "pineapple", "melon"}
// 	var newSlice = append(Slice, "Durian")

// 	fmt.Println(Slice)
// 	fmt.Println(newSlice)
// }

package main

import "fmt"

func main() {
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}
