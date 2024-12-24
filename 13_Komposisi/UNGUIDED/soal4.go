package main

import "fmt"

func main() {

	var bunga, pita string
	var jumlah int

	jumlah = 0

	for{
		fmt.Print("Bunga ", jumlah + 1, " : ")
		fmt.Scan(&bunga)

		if bunga == "selesai" || bunga =="SELESAI" {
		break
		}

		pita = pita + bunga + " - "
		jumlah++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
	
}