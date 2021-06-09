// func  (n nama_struct_model) nama_fungsi() {
// 	// n di atas merupakandeklarasi dari struct model
// 	// sama halnya  n := nama_struct_model
// }

// func  (n *nama_struct_model) nama_fungsi() {

// }
//contoh 1

// package main

// import (
// 	"fmt"
// )

// type Profile struct {
// 	nama string
// 	umur int
// }

// func (p Profile) detailProfile() {
// 	fmt.Println("Nama \t:", p.nama)
// 	fmt.Println("Umur \t:", p.umur)

// }

// func main() {
// 	P := Profile{
// 		nama: "getch",
// 		umur: 20,
// 	}
// 	P.detailProfile()
// }

//contoh 2 menggunakan pointer

package main

import (
	"fmt"
)

type Profile struct {
	Nama string
}

func (p Profile) detailProfile() {
	p.Nama = "Getch"

}
func (p *Profile) newdetailProfile() {
	p.Nama = "Miya"

}

func main() {
	P := Profile{
		Nama: "Jhonson",
	}
	P.detailProfile()
	fmt.Println(P)

	(&P).newdetailProfile()
	fmt.Println(P)

}
