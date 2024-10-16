package main

import (
	"fmt"
)

func main() {
	var n int

	// Meminta input jumlah segitiga
	fmt.Print("Masukkan jumlah segitiga (n): ")
	fmt.Scan(&n)

	// Loop untuk menerima input dan menghitung luas dari setiap segitiga
	for i := 1; i <= n; i++ {
		var alas, tinggi float64
		fmt.Printf("Masukkan alas dan tinggi segitiga ke-%d: ", i)
		fmt.Scan(&alas, &tinggi)

		// Menghitung luas segitiga
		luas := 0.5 * alas * tinggi

		// Menampilkan hasil luas segitiga
		fmt.Printf("Luas segitiga ke-%d adalah: %.2f\n", i, luas)
	}
}
