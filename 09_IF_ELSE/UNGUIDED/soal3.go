package main

import "fmt"

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scanln(&b)

	if b <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1.")
	}

	fmt.Print("Faktor: ")
	faktor := 0
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			faktor++
		}
	}
	fmt.Println()

	prima := b % 2

	if prima == 0 {
		fmt.Print("Prima : False")
	} else {
		fmt.Print("Prima : True")
	}
}