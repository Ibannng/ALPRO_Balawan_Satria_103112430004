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

	// Validasi durasi parkir
	if durasiParkir <= 0 {
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
		return
	}

	// Inisialisasi variabel untuk tarif
	var tarif int

	// Menentukan tarif berdasarkan jenis kendaraan dan durasi
	switch jenisKendaraan {
	case "motor":
		if durasiParkir <= 2 {
			tarif = 7000
		} else {
			tarif = 9000
		}
	case "mobil":
		if durasiParkir <= 2 {
			tarif = 15000
		} else {
			tarif = 20000
		}
	case "truk":
		if durasiParkir <= 2 {
			tarif = 25000
		} else {
			tarif = 35000
		}
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
		return
	}

	// Output hasil
	fmt.Printf("Tarif parkir untuk %s dengan durasi %d jam adalah: Rp %d\n", jenisKendaraan, durasiParkir, tarif)
}
