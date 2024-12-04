package main

import "fmt"

func main() {
    var harga, totalBelanja int

    fmt.Println("Masukkan harga barang satu per satu (ketik 0 untuk selesai):")

    for {
        fmt.Print("Masukkan harga barang: ")
        fmt.Scan(&harga)

        if harga == 0 {
            break
        }

        totalBelanja += harga
    }

    fmt.Printf("Total belanja Anda: Rp %d\n", totalBelanja)
}
