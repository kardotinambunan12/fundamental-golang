// package main

// import "fmt"

// func main() {
// 	for i := 0; i <= 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i, "genap")
// 		} else {
// 			fmt.Println(i, "ganjil")
// 		}

// 	}

// }

//penggunaan switch expression

// package main

// import "fmt"
// func main() {

//     var score = 'A'
//     switch score {
//         case 'A':
//             fmt.Println("Cool!")
//         case 'B':
//             fmt.Println("Great!")
//         case 'C':
//             fmt.Println("Not bad")
//         default:
//             fmt.Println("Study hard son...")
//     }
// }

package main

import (
	"fmt"
)

func main() {
	var nilai = 100
	switch {
	case nilai > 100:
		fmt.Println("sangat super")
	case nilai > 70 && nilai <= 99:
		fmt.Println("nilai bagus")
	case nilai > 90 && nilai <= 100:
		fmt.Println("nilai sempurna")
	default:
		fmt.Println("nilai tidak memenuhi")

	}

	// var price = 6500
	// switch {
	// case price > 20000:
	// 	fmt.Println("Mahal!")
	// case price > 10000 && price <= 20000:
	// 	fmt.Println("Lumayan")
	// case (price > 0) && (price <= 10000):
	// 	fmt.Println("Murah banget iki...")
	// default:
	// 	fmt.Println("Mangan gratisss")
	// }
}
