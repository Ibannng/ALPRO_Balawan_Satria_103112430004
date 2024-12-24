package main

import "fmt"

func main() {
    var a, b, c int

    // Meminta input tiga bilangan dari pengguna
    fmt.Print("Masukkan bilangan pertama: ")
    fmt.Scan(&a)
    fmt.Print("Masukkan bilangan kedua: ")
    fmt.Scan(&b)
    fmt.Print("Masukkan bilangan ketiga: ")
    fmt.Scan(&c)

    // Menentukan nilai terbesar dan terkecil
    max := a
    min := a

    if b > max {
        max = b
    }
    if c > max {
        max = c
    }

    if b < min {
        min = b
    }
    if c < min {
        min = c
    }

    // Menampilkan hasil
    fmt.Printf("Nilai terbesar: %d\n", max)
    fmt.Printf("Nilai terkecil: %d\n", min)
}
