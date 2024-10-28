package main

import (
	"fmt"
)

func isCumlaude(semester, eprtScore int) bool {
	return semester <= 8 && eprtScore >= 500
}

func main() {
	var semester, eprtScore int

	fmt.Println("Masukkan jumlah semester dan skor EPrT (pisahkan dengan spasi):")
	_, err := fmt.Scanf("%d %d", &semester, &eprtScore)
	if err != nil {
		fmt.Println("Error: Input tidak valid")
		return
	}

	result := isCumlaude(semester, eprtScore)

	fmt.Printf("Masukan: %d %d\n", semester, eprtScore)
	fmt.Printf("Keluaran: %t\n", result)

	if result {
		fmt.Printf("Penjelasan: mahasiswa lulus cumlaude dengan kuliah selama %d semester dan EPrT %d\n", semester, eprtScore)
	} else {
		if semester > 8 {
			fmt.Printf("Penjelasan: tidak cumlaude karena kuliah hingga %d semester\n", semester)
		} else {
			fmt.Printf("Penjelasan: tidak cumlaude karena nilai EPrT %d kurang dari 500\n", eprtScore)
		}
	}
}