package main

import "fmt"

func Sum(numbers ...int) (total int, avg float64) { // TODO: Predefine return values dimana nilai balik bisa didefinisikan di awal
	total = 0
	for _, number := range numbers {
		total += number
	}

	avg = float64(total) / float64(len(numbers))
	return
}

func main() {
	total, avg := Sum(1, 2, 3, 4, 5)
	fmt.Println("Total:", total)
	fmt.Println("Average:", avg)
}
