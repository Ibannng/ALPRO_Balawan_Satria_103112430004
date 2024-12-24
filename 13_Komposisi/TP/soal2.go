package main

import "fmt"

func main() {
    var num int
    fmt.Print("Masukkan sebuah bilangan: ")
    fmt.Scan(&num)

    sum := 0
    for i := 1; i < num; i++ {
        if num%i == 0 {
            sum += i
        }
    }

    if sum == num {
        fmt.Println("Ya") 
    } else {
        fmt.Println("Tidak") 
    }
}