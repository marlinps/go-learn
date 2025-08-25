package main

import "fmt"

func main() {
	// TODO: inisialisasi slice
	slice := []int{1, 2, 3, 4, 5}   // slice
	array := [5]int{6, 7, 8, 9, 10} // array
	array2 := [3]int{11, 12, 13}    // array

	_, _, _ = slice, array, array2

	// TODO:
	fruits := []string{"apple", "grape", "banana", "melon"}

	afruits := fruits[0:2]
	bfruits := fruits[1:4]

	fmt.Println("fruits:", fruits)
	fmt.Println("afruits[0][2]:", afruits)
	fmt.Println("bfruits[1][4]:", bfruits)

	bfruits[0] = "pinnaple" // TODO: "grape" diganti "pinnaple" maka elemen slice lama berubah semua

	fmt.Println("fruits:", fruits)
	fmt.Println("afruits[0][2]:", afruits)
	fmt.Println("bfruits[1][4]:", bfruits)

	// TODO: fungsi len()
	println("len slice afruits[0][2]:", len(afruits))
	println("len slice bfruits[1][4]:", len(bfruits))

	// TODO: fungsi cap()  -> menghitung kapasitas maksimal sice
	println("cap slice afruits[0][2]:", cap(afruits))
	println("cap slice bfruits[1][4]:", cap(bfruits))
}

/* TODO: notes
Slice adl reference elemen array (Tipe Data Reference).
perbedaan antara array dan slice:
1. Array adl kumpulan nilai atau elemen, Array jumlah elemen didefinisikan
2. Slice adl referensi tiap elemen, jumlah elemen tidak didefinisikan

penjelasan detaill hal 52
*/
