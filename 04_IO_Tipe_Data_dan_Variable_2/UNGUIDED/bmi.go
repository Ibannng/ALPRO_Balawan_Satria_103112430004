package main

import (
	"fmt"
)

func main() {
	var bmi, tinggi float64

	// Meminta input nilai BMI dari pengguna
	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&bmi)

	// Meminta input tinggi badan dalam meter dari pengguna
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scan(&tinggi)

	// Menghitung berat badan menggunakan rumus: berat = BMI * tinggi^2
	berat := bmi * (tinggi * tinggi)

	// Menampilkan hasil berat badan
	fmt.Printf("Berat badan yang sesuai dengan BMI %.2f dan tinggi %.2f m adalah: %.2f kg\n", bmi, tinggi, berat)
}
