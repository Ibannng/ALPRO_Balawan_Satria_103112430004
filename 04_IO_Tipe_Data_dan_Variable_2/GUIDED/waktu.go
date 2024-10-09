package main

import "fmt"

func main() {
	var input, detik, menit, jam int

	fmt.Print("Masukkan detik: ")
	fmt.Scan(&input)
	
	jam = input / 3600
	menit = (input % 3600) / 60
	detik = input % 60

	fmt.Println(jam, "jam", menit, "menit", detik, "detik")

}