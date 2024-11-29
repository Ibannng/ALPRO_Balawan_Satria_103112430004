package main

import "fmt"

func main() {
    const password = "golang123"
    var input string
    var attempts int

    fmt.Println("Sistem Login Sederhana")
    for attempts < 3 {
        fmt.Print("Masukkan password: ")
        fmt.Scan(&input)

        if input == password {
            fmt.Println("Login berhasil!")
            return
        }

        attempts++
        fmt.Printf("Password salah! Kesempatan tersisa: %d\n", 3-attempts)
    }

    fmt.Println("Login ditolak.")
}
