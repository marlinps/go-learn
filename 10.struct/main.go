package main

import "fmt"

type Student struct {
	Name  string
	Grade int
}

func main() {
	var studentA Student
	studentA.Name = "Budi"
	studentA.Grade = 10

	studentB := Student{
		Name:  "Siti",
		Grade: 9,
	}

	fmt.Println("Student A:", studentA)
	fmt.Println("Student B:", studentB)

	// TODO: struct variable object pointer
	var s2 *Student = &studentA                                 // pointer to variable pointer object
	fmt.Println("Alamat memory Student A (using pointer):", s2) // print alamat memory
	fmt.Println("Student A (value using pointer):", *s2)        // print value using pointer
	fmt.Println("Student A Name (using pointer):", s2.Name)     // print value using pointer, llangsung akses field tanpa dereference *s2.Name

	// TODO: keisitimewaan property dalam object pointer tanpa perlu deference nilai aslinya tetapi langsung akses fieldnya
	s2.Name = "Fadhil"
	s2.Grade = 9

	fmt.Println("Student A (after modification using pointer):", studentA) // print value asli yang sudah di modifikasi menggunakan pointer

}

/* TODO: struct
- struct adalah tipe data yang digunakan untuk mengelompokkan beberapa variable menjadi satu kesatuan
- struct bisa berisi variable dengan tipe data yang berbeda
- struct bisa diakses menggunakan dot notation (.)
*/
