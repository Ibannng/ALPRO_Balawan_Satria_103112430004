package main

import (
	"fmt"
)

func main() {
	var angka int

	// Meminta pengguna memasukkan sebuah angka
	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scan(&angka)

	if angka < 0 && angka%2 == 0 {
		fmt.Println("true.")
	} else {
		fmt.Println("false.")
	}
}
