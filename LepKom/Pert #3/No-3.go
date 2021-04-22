package main

import "fmt"

func main() {
	var i int

	for j := 1; j <= 10; j++ {
		fmt.Print("Input nilai i : ")
		fmt.Scan(&i)

		if i <= 10 {
			if i%2 == 0 {
				fmt.Println(i, "adalah angka genap")
				fmt.Println("---------------------")
			} else {
				fmt.Println(i, "adalah angka ganjil")
				fmt.Println("---------------------")
			}
		} else {
			fmt.Println("Pertanyaan selesai, karena angka i yang diinput sudah melebihi 10. Terima kasih.")
			break
		}
	}
}
