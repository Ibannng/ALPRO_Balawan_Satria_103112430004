package main

import (
	"fmt"
)

func main() {
	var angka int

	// Meminta pengguna memasukkan sebuah angka
	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scan(&angka)

	if angka > 0 {
		fmt.Println("positif.")
	} else {
		fmt.Println("negatif.")
	}
}
