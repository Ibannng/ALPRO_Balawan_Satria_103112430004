package main

import (
	"fmt"
)

func main() {
	var jam24 int

	// Input jam dalam format 24 jam
	fmt.Print("Masukkan jam dalam format 24 jam (0-23): ")
	fmt.Scan(&jam24)

	// Konversi ke format 12 jam menggunakan switch-case
	var jam12 int
	var periode string

	switch {
	case jam24 == 0:
		jam12 = 12
		periode = "AM"
	case jam24 > 0 && jam24 < 12:
		jam12 = jam24
		periode = "AM"
	case jam24 == 12:
		jam12 = 12
		periode = "PM"
	case jam24 > 12 && jam24 <= 23:
		jam12 = jam24 - 12
		periode = "PM"
	default:
		fmt.Println("Input tidak valid. Masukkan angka antara 0 dan 23.")
		return
	}

	// Output dalam format 12 jam
	fmt.Printf("Jam dalam format 12 jam: %02d %s\n", jam12, periode)
}
