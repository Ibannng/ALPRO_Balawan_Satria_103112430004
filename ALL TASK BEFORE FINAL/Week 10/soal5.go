package main

import (
	"fmt"
)

func main() {
	var totalBelanja int
	var buatKartu bool

	// Input
	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalBelanja)
	fmt.Print("Apakah ingin membuat kartu membership (true/false): ")
	fmt.Scan(&buatKartu)

	// Inisialisasi variabel
	var diskon bool
	var cashback bool
	var totalBayar int

	// Proses perhitungan
	if totalBelanja >= 100000 {
		diskon = true
	}

	if totalBelanja >= 200000 {
		cashback = true
	}

	if diskon {
		totalBayar = totalBelanja - (totalBelanja * 10 / 100) // Diskon 10%
	} else {
		totalBayar = totalBelanja
	}

	if cashback {
		totalBayar -= 75000 // Potongan cashback
	}

	// Output
	fmt.Printf("Kartu? %t\n", buatKartu)
	fmt.Printf("Diskon? %t\n", diskon)
	fmt.Printf("Cashback? %t\n", cashback)
	fmt.Printf("Rp. %d\n", totalBayar)
}


//Pseudocode
// program AkhirTahun
//     kamus
//         totalBelanja: INTEGER
//         bersediaKartu: BOOLEAN
//         kartu: BOOLEAN
//         diskon: BOOLEAN
//         cashback: BOOLEAN
//         finalBelanja: INTEGER

//     algoritma
//         // Input total belanja
//         OUTPUT "Masukkan total belanja:"
//         INPUT totalBelanja

//         // Input pilihan untuk membuat kartu
//         OUTPUT "Apakah bersedia dibuatkan kartu? (true/false):"
//         INPUT bersediaKartu

//         // Logika penentuan kartu, diskon, dan cashback
//         kartu <- bersediaKartu AND totalBelanja >= 100000
//         diskon <- totalBelanja >= 100000
//         cashback <- kartu AND totalBelanja >= 200000

//         // Inisialisasi nilai finalBelanja
//         finalBelanja <- totalBelanja

//         // Hitung diskon jika berlaku
//         IF diskon THEN
//             finalBelanja <- finalBelanja - (totalBelanja DIV 10) // Diskon 10%
//         ENDIF

//         // Hitung cashback jika berlaku
//         IF cashback THEN
//             finalBelanja <- finalBelanja - 75000 // Cashback Rp 75.000
//         ENDIF

//         // Output hasil
//         OUTPUT "Kartu? ", kartu
//         OUTPUT "Diskon? ", diskon
//         OUTPUT "Cashback? ", cashback
//         OUTPUT "Rp. ", finalBelanja
// end program
