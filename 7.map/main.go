package main

import "fmt"

func main() {

	var myMap = make(map[string]int)
	myMap["One"] = 1
	myMap["Two"] = 2
	myMap["Three"] = 3
	myMap["Four"] = 4

	fmt.Println(myMap)

	// map literal
	data := map[string]int{} // inisialisasi map kosong
	data["January"] = 1

	// map vertical literal, wajib (,) di akhir
	data2 := map[string]int{
		"January":  1,
		"February": 2,
		"March":    3,
		"April":    4,
	}

	fmt.Println(data)
	fmt.Println(data2)

	// map horizontal literal
	data3 := map[string]int{"January": 1, "February": 2, "March": 3, "April": 4}
	fmt.Println(data3)

	// new map (data pointer ke map)
	var data4 = new(map[string]int)
	*data4 = make(map[string]int)
	(*data4)["January"] = 1
	fmt.Println("alamat pointer data4:", data4)
	fmt.Println("value data4:", *data4)

	// print map
	for key, val := range data2 {
		fmt.Println(key, " \t:", val)
	}

	// akses map
	value, isExist := myMap["Two"]
	if isExist {
		fmt.Println("Key found:", value)
	} else {
		fmt.Println("Key not found")
	}

	// delete map
	delete(myMap, "Two")
	fmt.Println("After delete:", myMap)
}

/* TODO:
default value map is nil (harus di inisialisasasi terlebih dahulu dengan make() atau map literal agar tidak nil dan error)
*/
