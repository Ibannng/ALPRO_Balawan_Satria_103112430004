package main

import (
	"fmt"
)

func main() {
    angkaRahasia := 7 // Angka rahasia yang bisa Anda ubah
    var tebakan int

    fmt.Println("Selamat datang di permainan tebak angka!")

    for {
        fmt.Print("Tebak angka (1-10): ")
        fmt.Scan(&tebakan)

        if tebakan == angkaRahasia {
            fmt.Println("Selamat, tebakan Anda benar!")
            break
        } else {
            fmt.Println("Tebakan Anda salah, coba lagi.")
        }
    }
}
