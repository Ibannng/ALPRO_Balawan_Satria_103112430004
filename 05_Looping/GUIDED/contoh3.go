package main

import (
	"fmt"
)

func main() {
	var a, b int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan pertama (a): ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan kedua (b): ")
	fmt.Scan(&b)

	// Inisialisasi hasil perkalian
	hasil := 0

	// Perkalian dengan penjumlahan berulang
	for i := 0; i < b; i++ {
		hasil += a
	}

	// Menampilkan hasil perkalian
	fmt.Printf("Hasil perkalian %d dan %d adalah: %d\n", a, b, hasil)
}
