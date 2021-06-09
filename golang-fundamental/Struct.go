// package main

// import "fmt"

// type Profile struct {
// 	nama  string
// 	email string
// 	umur  int
// }

// type Person struct {
// 	Profile

// 	umur  int
// 	grade int
// }

// func main() {
// 	var p = Person{}
// 	p.nama = "getch"
// 	p.umur = 19
// 	p.grade = 3
// 	p.Profile.umur = 3
// 	p.Profile.email = "getch@mail.com"

// 	fmt.Println(p.nama)
// 	fmt.Println(p.umur)
// 	fmt.Println(p.grade)
// 	fmt.Println(p.Profile.umur)
// 	fmt.Println(p.Profile.email)

// }

package main

import "fmt"

type Profile struct {
	nama string
	umur int
}

func newProfile(nama string) *Profile {

	p := Profile{nama: nama}
	p.umur = 42
	return &p
}

func main() {
	fmt.Println(Profile{"Zilong", 20})

	fmt.Println(Profile{nama: "Alice", umur: 30})

	fmt.Println(Profile{nama: "miya"})

	fmt.Println(&Profile{nama: "Jhonson", umur: 40})

	fmt.Println(newProfile("Cyclop"))

	s := Profile{nama: "Harley", umur: 50}
	fmt.Println(s.nama)

	sp := &s
	fmt.Println(sp.umur)

	sp.umur = 51
	fmt.Println(sp.umur)
}
