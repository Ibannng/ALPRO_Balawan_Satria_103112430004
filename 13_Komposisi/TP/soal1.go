package main

import "fmt"

func main() {
    var bilangan int
    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&bilangan)

    for i := 2; i <= bilangan; i++ { 
        bilanganPrima := true
        for j := 2; j*j <= i; j++ { 
            if i%j == 0 {
                bilanganPrima = false 
                break
            }
        }
        if bilanganPrima { 
            fmt.Print(i, " ")
        }
    }
}