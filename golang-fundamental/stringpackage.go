package main

import (
	"fmt"
	s "strings"
)

func main() {
	var text = "medancity bung !"
	var kata_diubah = "u"
	var diganti = "a"

	var nama1 = s.Contains("kardo parna", "parna")
	var nama2 = s.Count("kardo parna", "a")
	var nama3 = s.Index("kardo parna", "par")
	var nama4 = s.Replace(text, kata_diubah, diganti, 1)
	var nama5 = s.Replace(text, kata_diubah, diganti, -1)
	var pesan = s.Repeat("medan", 3)
	var message = s.Split("ini medan bung", " ")
	var data = []string{"satu", "dua", "tiga"}
	datajoin := s.Join(data, " - ")

	var hurufbesar = s.ToUpper("golang string")
	var hurufkecil = s.ToLower("BELAJAR STRING")

	fmt.Println("contains:", nama1)
	fmt.Println("count:", nama2)
	fmt.Println("index :", nama3)
	fmt.Println("replace :", nama4)
	fmt.Println("replace all :", nama5)
	fmt.Println("repeat :", pesan)
	fmt.Println("split :", message)
	fmt.Println("join :", datajoin)
	fmt.Println("toupper :", hurufbesar)
	fmt.Println("tolower :", hurufkecil)

}
