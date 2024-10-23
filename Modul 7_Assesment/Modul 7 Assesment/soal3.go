package main

import (
	"fmt"
)

func main() {
	var qirat int

	fmt.Print("Masukkan jumlah uang dalam satuan qirat: ")
	fmt.Scan(&qirat)

	fals := qirat / 6
	sisaQirat := qirat % 6

	dirham := fals / 10
	sisaFals := fals % 10

	dinar := dirham / 10
	sisaDirham := dirham % 10

	fmt.Println("Dirham: ", dirham)
	fmt.Println("Dinar: ", dinar)
	fmt.Println("Qirat: ", fals)
	fmt.Println("Sisa Qirat: ", sisaFals)
	fmt.Println("Sisa Dirham: ", sisaDirham)
	fmt.Println("Sisa Dinar: ", sisaQirat)
}
