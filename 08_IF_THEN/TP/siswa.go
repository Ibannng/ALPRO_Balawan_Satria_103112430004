package main

import (
	"fmt"
)

func main() {
	var nilaiUjian int

	// Meminta pengguna memasukkan nilai ujian siswa
	fmt.Print("Masukkan nilai ujian siswa: ")
	fmt.Scan(&nilaiUjian)

	// Menentukan kelulusan
	if nilaiUjian >= 70 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}
}
