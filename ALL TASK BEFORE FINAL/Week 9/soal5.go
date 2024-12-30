package main

import "fmt"

func main() {
	var gol1, gol2, gol3, gol4 int

	fmt.Println("Masukkan jumlah gol untuk 4 tim:")
	fmt.Scan(&gol1, &gol2, &gol3, &gol4)

	max := gol1
	min := gol1

	if gol2 > max {
		max = gol2
	}
	if gol3 > max {
		max = gol3
	}
	if gol4 > max {
		max = gol4
	}

	if gol2 < min {
		min = gol2
	}
	if gol3 < min {
		min = gol3
	}
	if gol4 < min {
		min = gol4
	}

	fmt.Printf("Gol terbanyak: %d\n", max)
	fmt.Printf("Gol tersedikit: %d\n", min)
}

//Pseudocode:
// program LigaBola
//     kamus
//         gol1, gol2, gol3, gol4: integer
//         max, min: integer

//     algoritma
//         print "Masukkan jumlah gol untuk 4 tim:"
//         read gol1, gol2, gol3, gol4

//         max = gol1
//         min = gol1

//         if gol2 > max then
//             max = gol2
//         end if

//         if gol3 > max then
//             max = gol3
//         end if

//         if gol4 > max then
//             max = gol4
//         end if

//         if gol2 < min then
//             min = gol2
//         end if

//         if gol3 < min then
//             min = gol3
//         end if

//         if gol4 < min then
//             min = gol4
//         end if

//         print "Gol terbanyak: ", max
//         print "Gol tersedikit: ", min
// end program