/*
Jawaban Soal
a. Output atau keluaran dari hasil percabangan dengan inputan 80.1 adalah A dan program yang di minta adalah benar karena nilai A di dapatkan
   Ketika nilai > 80
b. Kesalahan yang pertama adalah variable pada print nilai yang harusnya menggunakan "nmk" bukan "nam" kesalahan selanjutnya adalah dengan mengasih
   range atau batasan nilai agar memunculkan output yang diinginkan. Selanjutnya jika sudah menggunakan if sangat di sarankan selanjutnya menggunakan
   else if
c. Lihat lah code dibawah ini
*/

package main

import "fmt"

func main() {
	var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)
	if nam > 80 {
		nmk = "A"
	} else if nam <= 80 && nam > 72.5 {
		nmk = "AB"
	} else if nam <= 72.5 && nam > 65 {
		nmk = "B"
	} else if nam <= 65 && nam > 57.5 {
		nmk = "BC"
	} else if nam <= 57.5 && nam > 50 {
		nmk = "C"
	} else if nam <= 50 && nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}
	fmt.Println("Nilai mata kuliah: ", nmk)
}