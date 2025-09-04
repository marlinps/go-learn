package main

import "fmt"

func main() {
	getFullName := func(firstName, lastName string) string {
		return firstName + " " + lastName
	}

	fmt.Println(getFullName("Rayyan", "Alvaro"))
}

/* Closure function/Anonymous function
- fungsi yang tidak memiliki nama
- biasanya digunakan sebagai argumen fungsi lain atau sebagai nilai balik dari fungsi lain
- bisa mengakses variabel di luar scope-nya (lexical scoping)
- sering digunakan untuk membuat fungsi yang hanya digunakan sekali saja
*/
