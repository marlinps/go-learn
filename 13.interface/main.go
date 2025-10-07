package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

// TODO: embedded interface
type hitung interface {
	hitung2d
	hitung3d
}
type lingkaran struct {
	jariJari float64
}

func (l *lingkaran) keliling() float64 {
	return 2 * math.Pi * l.jariJari
}

func (l *lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari, 2)
}

type persegi struct {
	sisi float64
}

func (p *persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p *persegi) keliling() float64 {
	return 4 * p.sisi
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return 6 * math.Pow(k.sisi, 2)
}

func (k *kubus) keliling() float64 {
	return 12 * k.sisi
}

func main() {
	// TODO: pemanggilan langsung
	l := lingkaran{4} // TODO: Go otomatis kasih &, jadi bisa dipanggil walaupun keliling() pointer receiver.
	fmt.Printf("Luas Lingkaran : %.2f\n", l.luas())
	fmt.Printf("Keliling Lingkaran :%.2f\n", l.keliling())

	p := persegi{2}
	fmt.Println("Luas Persegi :", p.luas())
	fmt.Println("Keliling Persegi :", p.keliling())

	// TODO: pemanggilan via interface
	var bangunDatar hitung2d // inisialisasi variable object bertipe interface
	fmt.Println("Implementasi variable object bertipe interface")
	bangunDatar = &lingkaran{4} //TODO:Tapi saat masuk ke interface, Go tidak otomatis kasih &, karena method set dicek ketat.
	fmt.Printf("Luas Bangun Datar Lingkaran: %.2f\n", bangunDatar.luas())
	fmt.Printf("Keliling Bangun Datar Lingkaran: %.2f\n", bangunDatar.keliling())

	bangunDatar = &persegi{4}
	fmt.Printf("Luas Bangun Datar Persegi: %.2f\n", bangunDatar.luas())
	fmt.Printf("Keliling Bangun Datar Persegi: %.2f\n", bangunDatar.keliling())

	// TODO: embedded interface
	var bangunRuang hitung = &kubus{4}
	fmt.Println("Implementasi Embedded Interface")
	fmt.Printf("Luas Bangun Ruang Kubus: %.2f\n", bangunRuang.luas())
	fmt.Printf("Keliling Bangun Ruang Kubus: %.2f\n", bangunRuang.keliling())
	fmt.Printf("Volume Bangun Ruang Kubus: %.2f\n", bangunRuang.volume())

}

/* TODO: Interface
adalah kumpulan definisi method yang tidak memiliki isi(hanya definisi saja), yang dibungkus dengan nama tertentu
interface merupakan tipe data seperti polimorphism dalam OPP
*/
