package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) sayHello(newName string) {
	s.Name = newName
	fmt.Println("Hello", s.Name)

}

// TODO: Method pointer receiver (kelebihan bisa mengubah property objectnya, akan diubah pada referencenya, hemat memory)
func (s *Student) sayHelloPointer(newName string) {
	s.Name = newName
	fmt.Println("Hello", s.Name)
}

func main() {
	s1 := Student{
		Name: "Eko",
		Age:  30,
	}

	s1.sayHello("Budi")
	fmt.Println(s1)

	// Call pointer receiver method to modify Name
	s1.sayHelloPointer("Budi")
	fmt.Println(s1)
}

/* TODO: Method
Method adl sebuah fungsi yang menempel pada type (bisa struct, interface, dll).
Method diakses lewat variable object, keunggulan method adalah bisa mengakses property dari object tersebut hingga level private
menggunakan proses encapsulation.
#
*/
