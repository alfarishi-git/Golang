package main

import "fmt"

func main() {
	var uts, uas, total float64
	var grade string

	fmt.Print("Input nilai UTS : ")
	fmt.Scan(&uts)
	fmt.Print("Input nilai UAS : ")
	fmt.Scan(&uas)

	total = (uts * 0.3) + (uas * 0.7)

	if total >= 81 {
		grade = "A"
	} else if total >= 61 {
		grade = "B"
	} else if total >= 41 {
		grade = "C"
	} else if total >= 21 {
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Println("-------------------")
	fmt.Println("Nilai UTS : ", uts)
	fmt.Println("Nilai UAS : ", uas)
	fmt.Println("Total nilai : ", total)
	fmt.Println("Grade : ", grade)
}
