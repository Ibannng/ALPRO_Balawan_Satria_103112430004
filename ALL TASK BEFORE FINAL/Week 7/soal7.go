package main

import "fmt"

func sumOddNumbers(a, b int) int {
	sum := 0

	if a%2 == 0 {
		a++
	}

	for i := a; i <= b; i += 2 {
		sum += i
	}

	return sum
}

func main() {
	var a, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	result := sumOddNumbers(a, b)
	fmt.Printf("Hasil penjumlahan bilangan ganjil dari %d sampai %d adalah: %d\n", a, b, result)
}

//Pseuodocode:
// program SumGanjil
//     kamus
//         a, b, sum: integer

//     algoritma
//         function sumOddNumbers(a, b: integer) -> integer
//             sum = 0
//             if a mod 2 == 0 then
//                 a = a + 1  
//             end if
//             for i = a to b step 2 do
//                 sum = sum + i 
//             end for
//             return sum
//         end function

//         print "Masukkan nilai a:"
//         read a
//         print "Masukkan nilai b:"
//         read b

//         result = sumOddNumbers(a, b)
//         print "Hasil penjumlahan bilangan ganjil dari ", a, " sampai ", b, " adalah: ", result
// end program