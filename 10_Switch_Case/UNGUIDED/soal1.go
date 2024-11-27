package main

import (
	"fmt"
)

func main() {
	var ph float64

	// Meminta input kadar pH dari pengguna
	fmt.Print("Masukkan kadar pH air: ")
	fmt.Scan(&ph)

	// Menentukan kelayakan air dengan switch-case
	switch {
	case ph < 0 || ph > 14:
		fmt.Println("Input tidak valid, rentang pH 0 - 14")
	case ph >= 6.5 && ph <= 8.6:
		fmt.Println("Air Layak Minum")
	default:
		fmt.Println("Air Tidak Layak Minum")
	}
}
