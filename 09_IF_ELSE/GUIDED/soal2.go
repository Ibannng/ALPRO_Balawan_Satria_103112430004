package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	// Meminta pengguna memasukkan satu huruf
	fmt.Print("Masukkan sebuah huruf: ")
	fmt.Scan(&input)

	// Mengubah input menjadi huruf besar untuk memudahkan pengecekan
	input = strings.ToUpper(input)

	// Memeriksa apakah input adalah huruf vokal
	if input == "A" || input == "I" || input == "U" || input == "E" || input == "O" {
		fmt.Println("Huruf Vokal")
	} else if input >= "A" && input <= "Z" {
		fmt.Println("Huruf Konsonan")
	} else {
		fmt.Println("Bukan huruf")
	}
}
