package main

import "fmt"

func main() {
	// TODO: Tipe Data Numerik
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// TODO: Tipe Data Numerik Desimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber) // %f  digunakan untuk memformat data numerik desimal menjadi string

	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber) // Jumlah digit yang muncul bisa dikontrol menggunakan %.nf , tinggal ganti n dengan angka yang diinginkan.

	// TODO: Tipe Data Boolean (bool)
	var exist bool = true
	fmt.Printf("exist ? %t \n", exist) //  %t  untuk memformat data  bool  menggunakan fungsi  fmt.Printf() .

	// TODO: Tipe Data string
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	// TODO: Penggunaan Konstanta
	const phi float32 = 3.14
	fmt.Println("Phi: ", phi)

	/* TODO: Nilai nil & Zero Value

	nil bukan merupakan tipe data, melainkan sebuah nilai. Variabel yg isi nilainya nil berarti memiiki nilai kosong

	Semua tipe data yang sudah dibahas di atas memiliki zero value (nilai default tipe data). Artinya meskipun variabel dideklarasikan
	dengan tanpa nilai awal, tetap akan ada nilai default-nya.

	Zero value dari  string  adalah  ""  (string kosong).
	Zero value dari  bool  adalah  false .
	Zero value dari tipe numerik non-desimal adalah  0 .
	Zero value dari tipe numerik desimal adalah  0.0 .

	Zero value berbeda dengan nil. Nil adl nilai kosong, benar-benar kosong. Ada beberapa tipe data yang bisa di-set nilainya dengan  nil , di antaranya:
	- pointer
	- tipe data fungsi
	- slice
	- map
	- channel
	- interface kosong atau  interface{}
	*/
}

/*  TODO: Tipe Data
Dianjurkan untuk tidak sembarangan dalam menentukan tipe data variabel, sebisa mungkin tipe yang dipilih harus disesuaikan dengan nilainya, karena efeknya adalah ke alokasi memori variabel. Pemilihan tipe data yang tepat akan membuat pemakaian memori lebih optimal, tidak berlebihan.
*/
