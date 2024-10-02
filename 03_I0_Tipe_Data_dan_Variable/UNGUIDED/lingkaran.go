package main

import (
	"fmt"
	"math"
)

func main() {
    var jari float64
    fmt.Print("Jejari = ")
    fmt.Scanln(&jari)

	// Rumus
    volume := (4.0 / 3.0) * math.Pi * math.Pow(jari, 3)
    luas := 4 * math.Pi * math.Pow(jari, 2)

    // Menampilkan hasil 
    fmt.Println("Bola dengan jejari", jari, "memiliki volume", volume, "dan luas kulit", luas)
}
