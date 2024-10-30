package main

import (
	"fmt"
)

func main() {
	var angka int

	// Meminta pengguna memasukkan sebuah angka
	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scan(&angka)

	// Mengecek apakah angka genap atau ganjil
	if angka%2 == 0 {
		fmt.Println("Angka adalah Genap.")
	} else {
		fmt.Println("Angka adalah Ganjil.")
	}
}
