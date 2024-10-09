package main

import "fmt"

func main() {

	var r float64

	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scan(&r)

	var luas float64 = 3.14 * r * r
	var keliling float64 = 2 * 3.14 * r

	fmt.Println("Hasil dari luas lingkaran dengan jari-jari:", r, "adalah", luas, "dan kelilingnya", keliling)
}