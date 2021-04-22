package main

import "fmt"

// Soal A
func main() {
	for i := 0; i < 15; i++ {
		if i%2 == 1 {
			fmt.Printf("Angka %d\n", i)
		}
	}
}

// Soal B
// func main() {
// 	for i := 0; i < 5; i++ {
// 		for j := i; j < 5; j++ {
// 			fmt.Print(i, " ")
// 		}
// 		fmt.Println()
// 	}
// }

// Soal C
// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i, " adalah angka genap")
// 		} else {
// 			fmt.Println(i, " adalah angka ganjil")
// 		}
// 	}
// }
