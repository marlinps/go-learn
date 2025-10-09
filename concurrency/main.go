package main

import (
	"fmt"
	"runtime"
	"time"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
	time.Sleep(10 * time.Millisecond) // beri waktu scheduler berpindah goroutine
}

func main() {
	runtime.GOMAXPROCS(2) // mengatur jumlah core yang digunakan

	go print(5, "Hello") // menjalankan fungsi print di goroutine baru
	print(5, "World")    // menjalankan fungsi print di goroutine utama

	var input string
	fmt.Scanln(&input) // menunggu input dari user agar program tidak langsung selesai
}
