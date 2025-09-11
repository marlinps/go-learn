package main

import (
	. "12.level-akses/library" // menambahkan titi (.) untuk mengimport semua struct dan function di package library tanpa menulis library.

	//TODO: pemanfaatan alias import
	f "fmt"
)

func main() {
	// library.SayHello("Marlin")
	SayHello("Marlin")

	// TODO: untuk penggunaan struct pada package lain harus dipastikan nama struct dan propertynya harus berbentuk Exported(CamelCase) kalau hanya memenuhi salah satu maka akan error
	// student := library.Student{
	// 	Name:  "Eko",
	// 	Grade: 10,
	// }

	student := Student{
		Name:  "Eko",
		Grade: 10,
	}

	//library.SayHello(student.Name)
	SayHello(student.Name)

	f.Println("Pemanfaatan alias import, menulisakan alias di awal import package")
}
