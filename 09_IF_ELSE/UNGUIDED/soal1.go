package main

import (
	"fmt"
)

func main() {
	var beratGr float64

	// Input berat parsel dalam gram
	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scan(&beratGr)

	// Hitung total berat dalam kg dan sisa dalam gram
	beratKg := int(beratGr) / 1000
	sisaGram := int(beratGr) % 1000

	// Biaya dasar (Rp. 10.000,- per kg)
	biayaDasar := beratKg * 10000
	tambahanBiaya := 0

	// Hitung tambahan biaya jika total berat <= 10kg
	if beratKg <= 10 {
		if sisaGram >= 500 {
			tambahanBiaya = sisaGram * 5
		} else {
			tambahanBiaya = sisaGram * 15
		}
	}

	// Hitung total biaya pengiriman
	totalBiaya := biayaDasar + tambahanBiaya

	// Tampilkan hasil perhitungan
	fmt.Printf("Total Berat: %d kg %d gram\n", beratKg, sisaGram)
	fmt.Printf("Biaya Dasar: Rp %d\n", biayaDasar)
	if beratKg <= 10 {
		fmt.Printf("Tambahan Biaya: Rp %d\n", tambahanBiaya)
	} else {
		fmt.Println("Tambahan Biaya: Gratis")
	}
	fmt.Printf("Total Biaya Pengiriman: Rp %d\n", totalBiaya)
}
