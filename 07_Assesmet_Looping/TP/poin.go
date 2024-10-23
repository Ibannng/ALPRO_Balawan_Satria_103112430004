package main

import (
	"fmt"
)

func main() {
	var jumlahBarang, totalPoin int

	// Input jumlah barang yang dibeli oleh pelanggan
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlahBarang)

	// Menghitung poin berdasarkan jumlah barang yang dibeli
	if jumlahBarang <= 5 {
		totalPoin = jumlahBarang * 10
	} else {
		// Untuk lebih dari 5 barang, 5 barang pertama mendapat 10 poin, sisanya mendapat 15 poin
		totalPoin = 5*10 + (jumlahBarang-5)*15
	}

	// Output total poin yang didapatkan pelanggan
	fmt.Printf("Total poin yang didapatkan: %d poin\n", totalPoin)
}
