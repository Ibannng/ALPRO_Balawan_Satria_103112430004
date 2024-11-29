package main

import "fmt"

func main() {
    var totalBelanja float64
    var harga float64
    var lanjut string

    fmt.Println("Aplikasi Kasir Sederhana")

    for {
        fmt.Print("Masukkan harga barang: ")
        fmt.Scan(&harga)
        totalBelanja += harga

        fmt.Print("Tambah barang lagi? (y/n): ")
        fmt.Scan(&lanjut)

        if lanjut == "n" || lanjut == "N" {
            break
        }
    }

    fmt.Printf("Total belanja: Rp %.2f\n", totalBelanja)
    fmt.Println("Terima kasih telah berbelanja!")
}
