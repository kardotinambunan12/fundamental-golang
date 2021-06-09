// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// x := make(map[string]int)
// 	// x["key"] = 10
// 	// fmt.Println(x["key"])
// 	var umur map[string]int
// 	umur = map[string]int{}

// 	umur["budi"] = 50
// 	umur["ani"] = 40

// 	fmt.Println("budi", umur["budi"])
// 	fmt.Println("ani", umur["ani"])
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	//cara deklarasi secara horizontal

// 	var umur =map[string]int{"andi ":50, "ani":40}

// 	//cara deklarasi secara vertical
// 	var umur = map[string]int{
// 		"andi" :50,
// 		"ani"  :40,
// 	}
// 	//eksekusi
// }
package main

import (
	"fmt"
)

func main() {
	var nama = map[string]int{"andi": 50, "ani": 40, "ino": 30, "koko": 10}

	for key, value := range nama {
		fmt.Println(key, " \t:", value)

	}
}
