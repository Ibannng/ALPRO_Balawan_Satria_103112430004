package main

import "fmt"

func main() {

	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	for i := 2; i*i <= bilangan; i++ {
		if bilangan%i == 0 {
            fmt.Println(bilangan, "bukan prima")
            return
        } 
	}
	fmt.Println(bilangan, "adalah prima")
}