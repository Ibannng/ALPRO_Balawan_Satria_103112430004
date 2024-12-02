package main

import (
	"fmt"
)

func main() {
	var input int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&input)

	fmt.Println("Digit pada bilangan (dari belakang ke depan):")

	// Memproses setiap digit dari belakang ke depan
	for input > 0 {
		digit := input % 10 // Mendapatkan digit terakhir
		fmt.Println(digit)  // Menampilkan digit
		input /= 10         // Menghapus digit terakhir
	}
}
