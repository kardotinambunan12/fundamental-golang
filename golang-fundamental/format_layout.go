package main

import (
	"fmt"
)

type Contoh struct {
	Nama   string
	Umur   uint8
	Tinggi float64
	Status bool
	Hobby  []string
}

var data = Contoh{
	Nama:   "Cyclop",
	Umur:   14,
	Tinggi: 120.1,
	Status: false,
	Hobby:  []string{"makan", "tidur", "makan lagi"},
}

//atau bisa juga menggunakan ini
//var Nama string = "nama"
//var Umur Uint8 = 14
//var Hobby  = [3]string{"data1", "data2", "data3"}

func main() {
	fmt.Printf("nama saya :%s \n", data.Nama)
	fmt.Printf("umur kamu %d \n", data.Umur)

	fmt.Printf("tinggi kamu %.f \n", data.Tinggi)
	fmt.Printf("tinggi kamu %.2f \n", data.Tinggi)
	fmt.Printf("tinggi kamu %f \n", data.Tinggi)

	fmt.Printf("status : %t\n", data.Status)

	fmt.Printf("angka biner sari umur :%b \n", data.Umur)

}
