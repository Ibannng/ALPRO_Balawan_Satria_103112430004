package main

import "fmt"

func main() { 

	var nilai int

	fmt.Print("Masukkan nilai: ") 
	fmt.Scan(&nilai)                    

	fx := (2/(nilai)+5) + 5 //persamaan f(x)
	fmt.Printf("Hasil f(x) adalah: %d\n", fx) 
}
