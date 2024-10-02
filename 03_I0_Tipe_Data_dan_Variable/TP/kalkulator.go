package main

import (
	"fmt"
	"os"
)

func main() {
    var num1, num2 float64
    var operator string

    fmt.Println("Kalkulator Sederhana")
    fmt.Print("Masukkan angka pertama: ")
    fmt.Scanln(&num1)
    fmt.Print("Masukkan operator (+, -, *, /): ")
    fmt.Scanln(&operator)
    fmt.Print("Masukkan angka kedua: ")
    fmt.Scanln(&num2)

    var result float64

    switch operator {
    case "+":
        result = num1 + num2
    case "-":
        result = num1 - num2
    case "*":
        result = num1 * num2
    case "/":
        if num2 != 0 {
            result = num1 / num2
        } else {
            fmt.Println("Error: Pembagian dengan nol")
            os.Exit(1)
        }
    default:
        fmt.Println("Operator tidak valid")
        os.Exit(1)
    }

    fmt.Printf("Hasil: %.2f\n", result)
}