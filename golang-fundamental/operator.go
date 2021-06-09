// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var nilai int
// 	nilai = 1 + 2

// 	// fmt.Println(nilai)

// 	// ata jika mau hasil yang spesifik bisa melakukan dengan cara seperti ini
// 	fmt.Printf("nilai :%d \n", nilai)
// }

// contoh 2

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var nilai = ((((2+3)%2)*3 - 1) / 2)

// 	// fmt.Println(nilai) atau
// 	fmt.Printf("hasil jumlah :%d \n", nilai)
// }

// menggunakan operator perbandingan

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var nilai = 10

// 	var perbandingan = (nilai < 10) //jika  (nilai <10) maka akan bernilai false
// 	fmt.Printf("nilai %d(%t) \n", nilai, perbandingan)
// }

// menggunakan operator logika

package main

import (
	"fmt"
)

func main() {
	var kiri = false
	var kanan = true

	var kiridankanan = kiri && kanan
	var kiriataukanan = kiri || kanan
	var kiriterbalik = !kiri

	fmt.Printf("hasil logika kiri dan kanan \t(%t) \n", kiridankanan)   // jika kedunaya bernilai true
	fmt.Printf("hasil logika kiri atau kanan \t(%t) \n", kiriataukanan) //jika kedua atau salah satu berniali true
	fmt.Printf("hasil logika kiri terballik \t(%t) \n", kiriterbalik)   // jika tidak sama dengan
}
