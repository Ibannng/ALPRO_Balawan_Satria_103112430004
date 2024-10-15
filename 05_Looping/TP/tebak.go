package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	targetNumber := rand.Intn(100) + 1
	var guess int
	attempts := 0

	fmt.Println("Selamat datang di permainan tebak angka!")
	fmt.Println("Saya telah memilih sebuah angka antara 1 dan 100.")

	for attempts < 5 {
		attempts++
		fmt.Printf("Percobaan ke-%d. Masukkan tebakan Anda: ", attempts)
		fmt.Scan(&guess)

		if guess == targetNumber {
			fmt.Printf("Selamat! Anda berhasil menebak angka %d dalam %d percobaan.\n", targetNumber, attempts)
			return
		} else if guess < targetNumber {
			fmt.Println("Terlalu kecil. Coba lagi.")
		} else {
			fmt.Println("Terlalu besar. Coba lagi.")
		}
	}

	fmt.Printf("Maaf, Anda kehabisan kesempatan. Angka yang benar adalah %d.\n", targetNumber)
}