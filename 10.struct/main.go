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
}

/* TODO: struct
- struct adalah tipe data yang digunakan untuk mengelompokkan beberapa variable menjadi satu kesatuan
- struct bisa berisi variable dengan tipe data yang berbeda
- struct bisa diakses menggunakan dot notation (.)
*/
