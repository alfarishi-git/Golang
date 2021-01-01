// SOAL 1
package main

import "fmt"

func main(){

	var a, b, c int

	a = 12
	b = 8
	c = a + b

	fmt.Println(c)
}

/*// SOAL 2
package main

import "fmt"

func main(){

	var a, b, c, x, y, z int

	a = 50
	b = 20
	c = 30

	x = a + b - c
	y = a * b
	z = a / b

	fmt.Println("Diketahui:\na : 50\nb : 20\nc : 30")
	fmt.Println("Hitunglah: \nx : a + b - c\ny : a * b\nz : a / b")
	fmt.Println("---------------------------")
	fmt.Println("Hasil Dari Operasi X adalah ", x)
	fmt.Println("Hasil Dari Operasi Y adalah ", y)
	fmt.Println("Hasil Dari Operasi Z adalah ", z)
}*/

/*// SOAL 3
package main

import "fmt"

func main(){

	for i := 1; i <= 10; i++ {
		fmt.Printf("i ke %d\n", i)
	}
}*/

/*// SOAL 4
package main

import "fmt"

func main(){

	i := 9

	if i > 10 {
		fmt.Println("Output Jika i lebih dari 10")
		fmt.Println("---------------------------")
		fmt.Println("You're Big")
	} else {
		fmt.Println("Output Jika i kurang dari 10")
		fmt.Println("---------------------------")
		fmt.Println("Sorry, You're Small")
	}
}*/

/*// SOAL 5
package main

import "fmt"

func main(){

	var grade int

	grade = 70

	switch grade {
	case 100:
		fmt.Println("Output Jika Nilai 100")
		fmt.Println("---------------------")
		fmt.Println("Excellent")
		break
	case 90:
		fmt.Println("Output Jika Nilai 90")
		fmt.Println("--------------------")
		fmt.Println("Good")
		break
	case 80:
		fmt.Println("Output Jika Nilai 80")
		fmt.Println("--------------------")
		fmt.Println("Better")
		break
	default:
		fmt.Println("Output Jika Nilai Tidak Ada Dalam Kriteria")
		fmt.Println("---------------------")
		fmt.Println("Sorry, You're Fail")
	}
}*/