package main

import "fmt"

func main() {
	var angka int
	for selesai := false; !selesai; {
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&angka)
		selesai = (angka > 0)
	}
	fmt.Println("angka yang anda berikan adalah bilangan positif")
}