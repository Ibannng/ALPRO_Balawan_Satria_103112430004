package main

import (
	"fmt"
)

func main() {
	var jenisKendaraan string
	var durasiParkir int

	// Input jenis kendaraan
	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&jenisKendaraan)

	// Input durasi parkir
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasiParkir)

	// Jika durasi parkir kurang dari 1 jam, anggap sebagai 1 jam
	if durasiParkir < 1 {
		durasiParkir = 1
	}

	// Menentukan tarif parkir berdasarkan jenis kendaraan
	var tarifPerJam int
	switch jenisKendaraan {
	case "motor":
		tarifPerJam = 2000
	case "mobil":
		tarifPerJam = 5000
	case "truk":
		tarifPerJam = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid")
		return
	}

	// Menghitung total biaya parkir
	totalBiaya := tarifPerJam * durasiParkir

	// Menampilkan total biaya parkir
	fmt.Printf("Total biaya parkir untuk kendaraan %s selama %d jam adalah: Rp %d\n", jenisKendaraan, durasiParkir, totalBiaya)
}
