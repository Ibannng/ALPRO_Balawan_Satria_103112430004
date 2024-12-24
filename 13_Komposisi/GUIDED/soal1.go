package main

import "fmt"

func main() {
    var n int

    // Meminta input bilangan dari pengguna
    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&n)

    fmt.Printf("Bilangan ganjil dari 1 hingga %d:\n", n)

    // Perulangan untuk mencetak bilangan ganjil
    for i := 1; i <= n; i++ {
        if i%2 != 0 {
            fmt.Print(i, " ")
        }
    }
    fmt.Println() // Membuat baris baru setelah output
}
