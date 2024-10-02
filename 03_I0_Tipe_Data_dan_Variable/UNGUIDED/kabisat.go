package main

import "fmt"

// Fungsi untuk memeriksa apakah tahun kabisat
func isKabisat(tahun int) bool {
    return tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0)
}

func main() {
    var tahun int
    fmt.Print("Tahun: ")
    fmt.Scanln(&tahun)

    // Menggunakan fmt.Println untuk output
    if isKabisat(tahun) {
        fmt.Println("Tahun:", tahun)
        fmt.Println("Kabisat: true")
    } else {
        fmt.Println("Tahun:", tahun)
        fmt.Println("Kabisat: false")
    }
}
