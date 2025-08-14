package main

import "fmt"

func main() {
	// TODO: inisialisasi array dengan nilai default
	var values [5]int
	values[0] = 1
	values[1] = 2
	values[2] = 3
	values[3] = 4
	values[4] = 5

	for i, value := range values {
		fmt.Println("angka"+"["+fmt.Sprint(i)+"]:", value)
	}

	// TODO: inisisasi arrat dengan nilai awal
	// Horizontal
	var fruits = [5]string{"apple", "mango", "watermelon", "orange", "banana"}
	for _, fruit := range fruits {
		fmt.Println("Fruit:", fruit)
	}

	fmt.Println("Vertikal")

	// vertikal
	var fruits2 = [5]string{
		"apple",
		"mango",
		"watermelon",
		"orange",
		"banana",
	}

	for _, fruit2 := range fruits2 {
		fmt.Println("Fruit2:", fruit2)
	}

	// TODO: inisialisasi arrayy tanpa jumlah elemen [...]
	var numbers = [...]int{1, 2, 3, 4, 5}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))
}
