package main

import (
	"fmt"
)

func main() {
	var number int

	// Input bilangan
	fmt.Print("Masukkan bilangan 4 digit (1000-9999): ")
	fmt.Scan(&number)

	// Validasi input harus 4 digit
	if number < 1000 || number > 9999 {
		fmt.Println("Bilangan harus 4 digit (1000-9999)")
		return
	}

	// Memisahkan setiap digit
	digit4 := number % 10          // digit terakhir
	digit3 := (number / 10) % 10   // digit ketiga
	digit2 := (number / 100) % 10  // digit kedua
	digit1 := (number / 1000) % 10 // digit pertama

	// Cek apakah digit terurut membesar
	if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
		fmt.Println("Digit terurut membesar")
	} else if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
		// Cek apakah digit terurut mengecil
		fmt.Println("Digit terurut mengecil")
	} else {
		fmt.Println("Digit tidak terurut")
	}
}