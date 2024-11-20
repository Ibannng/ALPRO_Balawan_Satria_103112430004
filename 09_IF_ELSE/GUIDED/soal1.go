package main

import (
	"fmt"
)

func main() {
	var umur int
	var KK bool

	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umur)
	fmt.Print("apakah mempunyai KK (true/false): ")
	fmt.Scan(&KK)

	if umur >= 17 && KK == true {
		fmt.Println("anda dapat membuat ktp")
	} else {
		fmt.Println("anda tidak dapat membuat ktp:")

		if umur < 17 {
			fmt.Println("umur anda dibawah umur")
		} else if KK != true {
			fmt.Println("Anda tidak memiliki kk")
		}
	}
}