package main

import (
	"fmt"
)

func main() {
	var angka int

	// Meminta pengguna memasukkan sebuah bilangan bulat
	fmt.Print("Masukkan sebuah bilangan bulat: ")
	fmt.Scan(&angka)

	// Menghitung nilai absolut
	if angka < 0 {
		angka = -angka
	}

	// Menampilkan nilai absolut
	fmt.Printf("Nilai absolut: %d\n", angka)
}
