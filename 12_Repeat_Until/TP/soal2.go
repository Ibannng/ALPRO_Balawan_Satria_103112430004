package main

import (
	"fmt"
	"strings"
)

func main() {
    var input string

    fmt.Println("Masukkan kata (ketik 'telkom' untuk berhenti):")

    for {
        fmt.Print("Masukkan kata: ")
        fmt.Scan(&input)

        if strings.ToLower(input) == "telkom" {
            fmt.Println("Program selesai.")
            break
        }

        fmt.Printf("Anda mengetik: %s\n", input)
    }
}
