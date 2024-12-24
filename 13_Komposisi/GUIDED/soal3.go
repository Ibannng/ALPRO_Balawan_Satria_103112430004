package main

import "fmt"

func main() {
    var x int

    // Meminta input dari pengguna
    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&x)

    if x <= 0 {
        fmt.Println("Input tidak valid. Masukkan bilangan bulat positif.")
        return
    }

    fmt.Printf("Faktor bilangan dari %d adalah:\n", x)

    // Mencari dan mencetak faktor bilangan
    for i := 1; i <= x; i++ {
        if x%i == 0 {
            fmt.Print(i, " ")
        }
    }
    fmt.Println() // Baris baru setelah output
}
