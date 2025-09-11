package main

import "12.level-akses/library"

func main() {
	library.SayHello("Marlin")

	// TODO: untuk penggunaan struct pada package lain harus dipastikan nama struct dan propertynya harus berbentuk Exported(CamelCase)
	student := library.Student{
		Name:  "Eko",
		Grade: 10,
	}

	library.SayHello(student.Name)
}
