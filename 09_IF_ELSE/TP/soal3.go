package main

import (
	"fmt"
)

func main() {
	var umur int
	var kewarganegaraan string

	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umur)
	fmt.Print("Masukkan kewarganegaraan: ")
	fmt.Scan(&kewarganegaraan)

	if umur >= 17 && kewarganegaraan == "WNI" {
		fmt.Println("Anda dapat mengikuti pemilu")
	} else {
		fmt.Println("Anda tidak bisa mengikuti pemilu karena:")

		if umur < 17 {
			fmt.Println("umur anda dibawah umur")
		} else if kewarganegaraan != "WNI" {
			fmt.Println("kewarganegaraan anda bukan WNI")
		}
	}
}