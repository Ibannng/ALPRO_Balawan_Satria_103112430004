package main

import (
	"fmt"
	"math"
)

func main() {
    var angka float64

    fmt.Print("Masukkan bilangan desimal: ")
    fmt.Scan(&angka)

    target := math.Ceil(angka) // Pembulatan ke atas
    for selesai := false; !selesai; {
        fmt.Printf("%.1f\n", angka)
        angka += 0.1
        selesai = angka >= target
    }
}
