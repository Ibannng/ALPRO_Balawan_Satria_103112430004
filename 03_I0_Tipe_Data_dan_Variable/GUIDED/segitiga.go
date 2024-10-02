package main

import "fmt"

func main() { 

	var alas int
	var tinggi int

	fmt.Print("Masukkan alas: ") 
	fmt.Scan(&alas)   
	fmt.Print("Masukkan tinggi: ") 
	fmt.Scan(&tinggi)                   

	fx := (alas / 2) * tinggi       //rumus segitiga      
	fmt.Printf("Hasil volume segitiga adalah: %d\n", fx) 
}
