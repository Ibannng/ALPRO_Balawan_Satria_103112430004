package main

import "fmt"

func main() {

	var kantong1, kantong2 float64

	for{

		fmt.Print("Masukkan berat kantong 1: ")
        fmt.Scan(&kantong1)
        fmt.Print("Masukkan berat kantong 2: ")
        fmt.Scan(&kantong2)

		if kantong1+kantong2 > 150 {
			fmt.Println("Proses selesai.")
			break
		}
		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := kantong1 - kantong2
		if selisih < 0 {
			selisih = -selisih 
		}

		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}

}