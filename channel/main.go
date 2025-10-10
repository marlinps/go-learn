package main

func main() {
	ch := make(chan int, 3) // membuat channel dengan buffer size 3

	go func() {
		ch <- 42 // mengirim data ke channel
	}()

	go func() {
		ch <- 27 // mengirim data ke channel
	}()

	go func() {
		ch <- 30 // mengirim data ke channel
	}()

	// menerima data dari channel
	value1 := <-ch
	println(value1)

	// menerima data dari channel
	value2 := <-ch
	println(value2)

	// menerima data dari channel
	value3 := <-ch
	println(value3)

	close(ch) // menutup channel setelah selesai digunakan
}

/* Channel adalah sebuah tipe data yang digunakan untuk mengirim dan menerima data antar goroutine. Channel memungkinkan komunikasi dan sinkronisasi antara goroutine, sehingga memudahkan pengelolaan konkuren dalam program Go.
 */
