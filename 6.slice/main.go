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
	fmt.Println("fruits[0][2]:", afruits)
	fmt.Println("fruits[1][4]:", bfruits)
}

/* TODO: notes
Slice adl reference elemen array (Tipe Data Reference)
perbedaan antara array dan slice:
1. Array adl kumpulan nilai atau elemen, Array jumlah elemen didefinisikan
2. Slice adl referensi tiap elemen, jumlah elemen tidak didefinisikan
penjelasan detail hal 52
*/
