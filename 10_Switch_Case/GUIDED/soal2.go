package main

import (
	"fmt"
)

func main() {
	var namaTanaman string

	// Input nama tanaman
	fmt.Print("Masukkan nama tanaman: ")
	fmt.Scan(&namaTanaman)

	// Inisialisasi variabel untuk hasil
	var jenisTanaman, asalTanaman string

	// Menentukan jenis tanaman dan asalnya menggunakan switch-case
	switch namaTanaman {
	case "Kantong Semar":
		jenisTanaman = "Termasuk Tanaman Karnivora"
		asalTanaman = "Asli Indonesia"
	case "Venus Flytrap":
		jenisTanaman = "Termasuk Tanaman Karnivora"
		asalTanaman = "Bukan Asli Indonesia"
	case "Nepenthes":
		jenisTanaman = "Termasuk Tanaman Karnivora"
		asalTanaman = "Asli Indonesia"
	case "Sundew":
		jenisTanaman = "Termasuk Tanaman Karnivora"
		asalTanaman = "Bukan Asli Indonesia"
	default:
		jenisTanaman = "Tidak termasuk Tanaman Karnivora"
		asalTanaman = "Tidak Diketahui"
	}

	// Output hasil
	fmt.Println(jenisTanaman)
	if asalTanaman != "Tidak Diketahui" {
		fmt.Println(asalTanaman)
	}
}
