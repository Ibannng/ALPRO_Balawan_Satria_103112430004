package main

import "fmt"

func main() {
    var angka int
    var jumlahDigit int

    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&angka)

    for selesai := false; !selesai; {
        jumlahDigit++
        angka /= 10
        selesai = (angka == 0)
    }

    fmt.Printf("Jumlah digit: %d\n", jumlahDigit)
}
