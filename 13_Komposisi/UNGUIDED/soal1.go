package main

import "fmt"

func main () {

	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	hitung := 0

	for i := 1; i <= bilangan; i++ {
        if i%2 != 0 {
            hitung++
        }
    }

	fmt.Printf("Terdapat %d bilangan ganjil \n", hitung)
}