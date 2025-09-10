package main

import "fmt"

type Student struct {
	Name  string
	Grade int
}

// TODO: embedded struct (menempelkan sebuah struct ke property struct lainnya)
type Course struct {
	NameCourse string
	Student    // embedded struct
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
	fmt.Println("Student A Name (using pointer):", s2.Name)     // print value using pointer, langsung akses field tanpa dereference *s2.Name

	// TODO: keistimewaan property dalam object pointer tanpa perlu deference nilai aslinya tetapi langsung akses fieldnya
	s2.Name = "Fadhil"
	s2.Grade = 9
	fmt.Println("Student A (after modification using pointer):", studentA) // print value asli yang sudah di modifikasi menggunakan pointer

	// TODO: Implement embedded struct
	// 1. cara pertama
	courseA := Course{
		NameCourse: "Programming",
		Student:    studentA,
	}
	fmt.Println("CourseA:", courseA)

	// 2. cara kedua
	var courseB Course
	courseB.NameCourse = "Calculus"
	courseB.Name = "Andi" // akses langsung field struct Student yang di embedded
	courseB.Grade = 8
	fmt.Println("CourseB:", courseB)

	// 3. cara ketiga
	coursecC := Course{
		NameCourse: "Machine Learning",
		Student: Student{
			Name:  "Charles",
			Grade: 11,
		},
	}
	fmt.Println("CourseC:", coursecC)
	fmt.Println("CourseC Student Name:", coursecC.Name)                                        // akses langsung field struct Student yang di embedded
	fmt.Println("CourseC student Name (using parent struct(Student)):", coursecC.Student.Name) // akses field struct Student yang di embedded menggunakan parent struct (Student) hasilnya sama saja

	/* TODO: Notes tentang embedded struct
	   - jika ada field yang sama pada parent struct dan embedded struct, maka untuk mengakses field embedded struct harus menggunakan parent struct secara eksplisit atau jelas
	*/

}

/* TODO: struct
- struct adalah tipe data yang digunakan untuk mengelompokkan beberapa variable menjadi satu kesatuan
- struct bisa berisi variable dengan tipe data yang berbeda
- struct bisa diakses menggunakan dot notation (.)
*/
