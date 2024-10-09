package main

import (
	"fmt"
)

func main() {
    var jamKerja, upahPerJam float64

    fmt.Print("Masukkan jumlah jam kerja dalam seminggu: ")
    fmt.Scanln(&jamKerja)

    fmt.Print("Masukkan upah per jam: ")
    fmt.Scanln(&upahPerJam)

    var totalGaji float64
    jamNormal := 40.0
    jamLembur := 0.0

    if jamKerja > jamNormal {
        jamLembur = jamKerja - jamNormal
        totalGaji = (jamNormal * upahPerJam) + (jamLembur * 1.5 * upahPerJam)
    } else {
        totalGaji = jamKerja * upahPerJam
    }

    // Mengalikan dengan 4 untuk mendapatkan gaji bulanan (asumsi 4 minggu per bulan)
    gajiBulanan := totalGaji * 4

    fmt.Printf("Total gaji bulanan: Rp %.2f\n", gajiBulanan)
}