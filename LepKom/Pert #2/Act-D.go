package main

import "fmt"

var nilai1, nilai2 int

func main() {
	defer fmt.Println("---- SELESAI ----")
	fmt.Print("Masukkan nilai 1: ")
	fmt.Scan(&nilai1)
	fmt.Print("Masukkan nilai 2: ")
	fmt.Scan(&nilai2)

	hasil := nilai1 % nilai2

	fmt.Printf("Hasil dari Nilai1 / Nilai 2: %d\n", hasil)
}
