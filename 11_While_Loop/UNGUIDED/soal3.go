package main

import "fmt"

func main() {
    var x, y int

    // Meminta input dari pengguna
    fmt.Print("Masukkan bilangan x (>= y): ")
    fmt.Scan(&x)
    fmt.Print("Masukkan bilangan y (divisor): ")
    fmt.Scan(&y)

    if y == 0 {
        fmt.Println("Pembagian dengan 0 tidak valid.")
        return
    }

    hasil := 0

    // Melakukan pengurangan berulang hingga x < y
    for x >= y {
        x -= y
        hasil++
    }

    // Menampilkan hasil
    fmt.Printf("Hasil dari x div y adalah: %d\n", hasil)
}
