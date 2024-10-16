package main

import (
	"fmt"
)

func main() {
	var a, b int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan b: ")
	fmt.Scan(&b)

	// Memastikan a <= b
	if a <= b {
		fmt.Printf("Baris bilangan dari %d hingga %d: ", a, b)
		for i := a; i <= b; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	} else {
		fmt.Println("Nilai a harus kurang dari atau sama dengan b.")
	}
}
