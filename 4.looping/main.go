package main

import "fmt"

func main() {
	// TODO: for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// TODO: for = while (dengan argumen hanya kondisi)
	j := 0
	for j < 5 {
		println("Angka", j)
		j++
	}

	// TODO: for tanpa argumen
	k := 0
	for {
		fmt.Println("Angka", k)
		k++

		if k == 5 {
			break
		}
	}

	// TODO: for - range

	// TODO: break & continue
	for l := 0; l <= 10; l++ {
		if l%2 == 1 {
			continue
		}

		if l > 8 {
			break
		}
		fmt.Println("angka", l)
	}

	// TODO: Perulangan bersarang
	for m := 0; m < 5; m++ {
		for n := m; n < 5; n++ {
			fmt.Print(n, " ")
		}

		fmt.Println()
	}

	// TODO: Pemanfaatan Label Dalam Perulangan
	// TODO: conteh kata outerloop (maka break dipanggil dengan target adalah perulangan yang dilabeli outerloop)
outerLoop:
	for o := 0; o < 5; o++ {
		for p := 0; p < 5; p++ {
			if o == 3 {
				break outerLoop // break
			}
			fmt.Print("matriks [", o, "][", p, "]", "\n")
		}
	}

}
