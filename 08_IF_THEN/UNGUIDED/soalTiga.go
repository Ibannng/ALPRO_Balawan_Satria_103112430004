package main

import (
	"fmt"
)

func main() {
	var x, y int

	// Meminta pengguna memasukkan dua bilangan bulat positif
	fmt.Print("Masukkan bilangan pertama (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan kedua (y): ")
	fmt.Scan(&y)

	// Mengecek apakah x adalah faktor dari y
	if y%x == 0 {
		fmt.Printf("Apakah %d adalah faktor dari %d? true\n", x, y)
	} else {
		fmt.Printf("Apakah %d adalah faktor dari %d? false\n", x, y)
	}

	// Mengecek apakah y adalah faktor dari x
	if x%y == 0 {
		fmt.Printf("Apakah %d adalah faktor dari %d? true\n", y, x)
	} else {
		fmt.Printf("Apakah %d adalah faktor dari %d? false\n", y, x)
	}
}
