package main

import (
	"fmt"
)

func main() {
	var num int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan bulat 3 digit (100-999): ")
	fmt.Scan(&num)

	// Memastikan bilangan dalam rentang 100 hingga 999
	if num < 100 || num > 999 {
		fmt.Println("Masukkan bilangan 3 digit yang valid.")
		return
	}

	// Memisahkan digit-digit bilangan
	// ratusan, puluhan, satuan
	ratusan := num / 100       // digit pertama
	puluhan := (num / 10) % 10 // digit kedua
	satuan := num % 10         // digit ketiga

	// Mengecek apakah setiap digit terurut membesar
	if ratusan < puluhan && puluhan < satuan {
		fmt.Println(true) // Terurut membesar
	} else {
		fmt.Println(false) // Tidak terurut membesar
	}
}
