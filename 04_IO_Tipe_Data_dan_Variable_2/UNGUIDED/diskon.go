package main

import (
	"fmt"
)

func main() {
	var totalBelanja, diskon int

	// Meminta input total belanja awal dari pengguna
	fmt.Print("Masukkan total belanja awal: ")
	fmt.Scan(&totalBelanja)

	// Meminta input besaran diskon dalam persen dari pengguna
	fmt.Print("Masukkan besarnya diskon (%): ")
	fmt.Scan(&diskon)

	// Menghitung jumlah diskon
	jumlahDiskon := float64(totalBelanja) * float64(diskon) / 100

	// Menghitung total belanja akhir setelah diskon
	totalAkhir := float64(totalBelanja) - jumlahDiskon

	// Menampilkan hasil total belanja akhir
	fmt.Printf("Total belanja akhir setelah diskon adalah: %.2f\n", totalAkhir)
}
