package main

import "fmt"

func main() {
    var fahrenheit float64

    fmt.Println("Konversi Suhu Fahrenheit ke Kelvin")
    fmt.Print("Masukkan suhu dalam Fahrenheit: ")
    fmt.Scanln(&fahrenheit)

    kelvin := (fahrenheit - 32) * 5 / 9 + 273.15

    fmt.Printf("%.2f°F = %.2f K\n", fahrenheit, kelvin)
}