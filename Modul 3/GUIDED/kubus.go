package main

import "fmt"

func main() { 

	var sisi int

	fmt.Print("Masukkan sisi: ") 
	fmt.Scan(&sisi)                    

	fx := sisi * sisi * sisi         //rumus kubus        
	fmt.Printf("Hasil volume kubus adalah: %d\n", fx) 
}
