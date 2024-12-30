package main

import "fmt"

func main() {
    var n int
    
    fmt.Print("Masukkan sebuah bilangan positif: ")
    fmt.Scan(&n)
    
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            if j == i || j == (n-i+1) {
                fmt.Printf("%d ", i)
            } else {
                fmt.Print("  ")
            }
        }
        fmt.Println()
    }
}

//Pseudocode
// program PolaBilanganX
//     kamus
//         n, i, j: integer

//     algoritma
//         print "Masukkan sebuah bilangan positif:"
//         read n

//         for i = 1 to n do
//             for j = 1 to n do
//                 if j = i or j = (n - i + 1) then
//                     print i dengan spasi
//                 else
//                     print spasi kosong
//                 end if
//             end for
//             print newline
//         end for
// end program