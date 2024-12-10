package main

import "fmt"

func main() {
    var target, total, donasi, jumlahDonatur int

    fmt.Print("Masukkan target donasi: ")
    fmt.Scan(&target)

    for selesai := false; !selesai; {
        jumlahDonatur++
        fmt.Printf("Donatur %d, masukkan jumlah donasi: ", jumlahDonatur)
        fmt.Scan(&donasi)

        total += donasi
        fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, total)
        selesai = total >= target
    }

    fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, jumlahDonatur)
}
