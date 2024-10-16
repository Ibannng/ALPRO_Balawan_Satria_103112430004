package main

import (
	"fmt"
	"math"
)

func main() {

	var ax, ay, bx, by, cx, cy, SisiTerpanjang, SisiAB, SisiBC, SisiCA float64

	fmt.Print("Masukan koordinat titik A (x, y): ")
	fmt.Scan(&ax, &ay)

	fmt.Print("Masukan koordinat titik B (x, y): ")
	fmt.Scan(&bx, &by)

	fmt.Print("Masukan koordinat titik C (x, y): ")
	fmt.Scan(&cx, &cy)

	SisiAB = math.Sqrt(math.Pow((bx-ax), 2) + math.Pow((by-ay), 2))
	SisiBC = math.Sqrt(math.Pow((cx-bx), 2) + math.Pow((cy-by), 2))
	SisiCA = math.Sqrt(math.Pow((ax-cx), 2) + math.Pow((ay-cy), 2))

	// Menggunakan math.Max untuk menentukan sisi terpanjang
	SisiTerpanjang = math.Max(SisiAB, math.Max(SisiBC, SisiCA))

	fmt.Printf("Sisi terpanjang adalah: %.2f\n", SisiTerpanjang)

}