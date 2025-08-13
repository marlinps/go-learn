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

	fmt.Printf("bilangan desimal: %f\n", decimalNumber) //  %f  digunakan untuk memformat data numerik desimal menjadi string

	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber) // Jumlah digit yang muncul bisa dikontrol menggunakan %.nf , tinggal ganti n dengan angka yang diinginkan.

	// TODO: Tipe Data Boolean (bool)
	var exist bool = true
	fmt.Printf("exist ? %t \n", exist) //  %t  untuk memformat data  bool  menggunakan fungsi  fmt.Printf() .

	// TODO: Tipe Data string
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)
}

/*  TODO: Tipe Data
Dianjurkan untuk tidak sembarangan dalam menentukan tipe data variabel, sebisa mungkin tipe yang dipilih harus disesuaikan dengan nilainya, karena efeknya adalah ke alokasi memori variabel. Pemilihan tipe data yang tepat akan membuat pemakaian memori lebih optimal, tidak berlebihan.
*/
