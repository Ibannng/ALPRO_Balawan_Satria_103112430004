package main

import "fmt"

func main() {
	var x float64
	var hasil float64

	fmt. Println("Masukkan nilai x: ")
	fmt. Scanln(&x)

	hasil = (2/ (x+5)) + 5	//rumus

	fmt.Println("Hasil dari",x, "adalah ",  hasil)
}