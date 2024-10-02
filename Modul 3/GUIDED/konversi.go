package main

import "fmt"

func main() {
    var idr float64

    // Meminta input jumlah uang dalam IDR
    fmt.Print("Masukkan jumlah uang dalam IDR: ")
    fmt.Scanln(&idr)

    // Menghitung konversi ke USD dengan kurs langsung 1 USD = 15,000 IDR
    usd := idr / 15000

    // Menampilkan hasil konversi
    fmt.Printf("Jumlah %.2f IDR sama dengan %.2f USD\n", idr, usd)
}
