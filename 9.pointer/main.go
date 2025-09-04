package main

import "fmt"

// TODO: variable pointer hanya bisa menampung nilai pointer/alamat memory

func main() {
	numberA := 4
	var numberB *int = &numberA // pointer ke , ampersand(&) mengambil alamat memory variable, asterisk(*) untuk mengambil value dari alamat memory

	fmt.Println("numberA (value):", numberA)
	fmt.Println("numberB (address):", numberB)
	fmt.Println("numberB (value):", *numberB) // dereference pointer

	numberA = 8
	fmt.Println("numberA (value):", numberA)
	fmt.Println("numberB (value):", *numberB) // dereference pointer
}

/* TODO: pointer
- pointer adalah variable yang menyimpan alamat memory dari variable lain
- operator & untuk mengambil alamat memory dari variable
- operator * untuk mengambil value dari alamat memory (dereference)
- pointer bisa digunakan untuk mengubah value dari variable lain
- pointer bisa digunakan untuk menghemat memory karena tidak perlu menyalin value dari variable lain, cukup menyalin alamat memorynya saja
- pointer bisa digunakan untuk menghemat waktu karena tidak perlu menyalin value dari variable lain, cukup menyalin alamat memorynya saja
- pointer bisa digunakan untuk menghemat resource karena tidak perlu menyalin value dari variable lain, cukup menyalin alamat memorynya saja
- pointer bisa digunakan untuk menghemat space karena tidak perlu menyalin value dari variable lain, cukup menyalin alamat memorynya saja
*/
