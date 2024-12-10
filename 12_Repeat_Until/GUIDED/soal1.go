package main

import "fmt"

func main() {
	var kata string
	var perulangan int

// Meminta input dari pengguna
	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)

	fmt.Print("Masukkan jumlah perulangan: ")
	fmt.Scan(&perulangan)
	counter := 0
	for done := false; !done; {
		fmt.Print(kata, " ")
		counter++
		done = (counter >= perulangan)
	}
}

