package main

import (
	"fmt"
)

func main() {
	var x, y, sum int

	fmt.Print("Masukkan nilai x (lebih kecil atau sama dengan y): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	for i := x; i <= y; i++ {
		sum += i
	}
	
	fmt.Println("Hasil penjumlahan dari", x, "sampai", y, "adalah:", sum)
}
