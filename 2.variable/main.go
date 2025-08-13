package main

import "fmt"

func main() {
	// TODO: Deklarasi Variable Berserta Var dan Tipe Data
	var firstName string = "john"

	var lastName string
	lastName = "wick"

	fmt.Printf("Halo %s %s!\n", firstName, lastName)
	fmt.Println("Halo", firstName, lastName+"!")

	// TODO: Deklarasi Variable Tanpa Tipe Data (type inference)
	var firstName2 = "alfa"
	lastName2 := "edison" //  hanya di assign sekali diawal deklarasi

	fmt.Printf("Halo, %s %s!\n", firstName2, lastName2)

	// TODO: Deklarasi Multi Variable
	var first, second, third int
	first, second, third = 1, 2, 3

	fourth, fifth, sixth := 4, 5, 6

	fmt.Println(first, second, third)
	fmt.Println(fourth, fifth, sixth)

	// Deklarasi multi variable dengan tipe data yang berbeda
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println(one, isFriday, twoPointTwo, say)

	// TODO: Variable Underscore _ (bahwa variable tidak digunakan)
	_ = "belajar golang"

	// TODO: Deklarasi Variable menggunakan Keyword new (variable pointer)
	name := new(string)

	fmt.Println(name)  // TODO: menampilkan alamat memori stringnya tersebut
	fmt.Println(*name) // TODO: asterisk (*) u/ menampilkan nilai aslinya
	fmt.Println(&name) // TODO: Jadi ini alamat memori tempat pointer disimpan (bukan alamat stringnya).

	/*	TODO: Deklarasi Variable menggunakan Keywaord make
		biasa digunukan u/ beberapa jenis variable saja seperti channel, slice, map
	*/

}
