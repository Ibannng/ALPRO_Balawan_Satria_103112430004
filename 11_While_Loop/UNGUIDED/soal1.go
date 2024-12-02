package main

import "fmt"

func main() {
	// Username dan password yang benar
	const correctUsername = "Admin"
	const correctPassword = "Admin"

	// Variabel untuk input pengguna dan jumlah percobaan gagal
	var username, password string
	var attempts int

	fmt.Println("Login System")

	for {
		// Meminta input username dan password
		fmt.Print("Masukkan username: ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&password)

		// Memeriksa kecocokan username dan password
		if username == correctUsername && password == correctPassword {
			// Jika benar, keluar dari perulangan
			fmt.Printf("Login berhasil setelah %d percobaan gagal.\n", attempts)
			break
		} else {
			// Jika salah, tambahkan percobaan gagal
			attempts++
			fmt.Println("Username atau password salah, coba lagi.")
		}
	}
}
