// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	nilai_string := "12345"
// 	var nilai, err = strconv.Atoi(nilai_string)

// 	if err == nil {
// 		fmt.Println(nilai)
// 	}
// }

// cara deklarasi ITOA atau int ke string
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var nilai = "1233"
// 	var nilaikonversi, err = strconv.ParseInt(nilai, 10, 16)

// 	if err == nil {
// 		fmt.Println(nilaikonversi)
// 	}

// 	// nilai_string := 12345

// 	// var nilai = strconv.Itoa(nilai_string)

// 	// // if err == "" {
// 	// fmt.Println(nilai)
// 	// // }
// }
package main

import (
	"fmt"
	"strconv"
)

func main() {
	string_pecahan := "123.12"
	var nilai_string, err = strconv.ParseFloat(string_pecahan, 64)

	if err == nil {
		fmt.Println(nilai_string)
	}
}
