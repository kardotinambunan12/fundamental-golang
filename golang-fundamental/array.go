// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var fruits [5]string
// 	fruits[0] = "apple"
// 	fruits[1] = "maggo"
// 	fruits[2] = "banana"
// 	fruits[3] = "pineapple"
// 	fruits[4] = "watermelon"

// 	fmt.Println(fruits[0], fruits[1], fruits[2], fruits[3], fruits[4])

// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var fruits = [5]string{"apple", "maggo", "banana", "pineapple", "melon"}

// 	fmt.Println(fruits)
// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	//  var nilaike0 = [2][3]int{[3]int{1, 2, 3}, [3]int{3, 2, 1}}

// 	// atau bisa juga dengan deklarasi begini

// 	var nilaike1 = [2][3]int{{1, 2, 3}, {3, 2, 1}}
// 	var nilaike2 = [2][3]int{{4, 5, 6}, {6, 5, 4}}

// 	fmt.Println("nilai ke 1 :", nilaike1)
// 	fmt.Println("nilai ke 2 :", nilaike2)
// }

package main

import (
	"fmt"
)

func main() {
	var fruits = [5]string{"apple", "maggo", "banana", "pineapple", "melon"}
	//cara 1

	// for i := 0; i < len(fruits); i++ {
	// 	// fmt.Printf("elemen %d : %s\n", i, fruits[i])
	// 	fmt.Printf("element %d : %s\n", i, fruits[i])
	// }

	//cara 2 denggan menggunakan range

	// for i, fruits := range fruits {
	// 	fmt.Printf("elemnt %d :%s\n", i, fruits)  //nilai variabelnya di tampung dalam fruit dan indexnya di tampung dalam i
	// }

	//cara 3  dengan menggunakan range dan ( _ )

	for _, fruits := range fruits {
		fmt.Printf("list buah : %s\n", fruits)
	}

}
