// package main

// import "fmt"

// type Nama_variabel interface{
// 	object() type_data
// }
// import (
// 	"fmt"
// )

// func Sampel(s interface{}) {
// 	val := s.(string)
// 	fmt.Println("message:", val)

// }
// func main() {
// 	var val interface{} = 1
// 	Sampel(val)

// }
package main

import (
	"fmt"
)

type Lemari interface {
	rak()
	laci()
}
type Jenis struct {
	ukuran int
	bentuk string
}

func (J Jenis) rak() {
	fmt.Println(J.ukuran)

}
func (J Jenis) laci() {
	fmt.Println(J.bentuk)

}
func Tempat(k Lemari) {
	k.rak()
	k.laci()

}

func main() {
	ukr := Jenis{20, "bulat"}
	Tempat(ukr)

}
