package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Printf("Jumlah deret angka dari 1 hingga %d adalah: %d\n", n, sum)
}