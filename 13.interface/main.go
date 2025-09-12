package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	jariJari float64
}

func (l *lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari, 2)
}

func (l *lingkaran) keliling() float64 {
	return 2 * math.Pi * l.jariJari
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

func main() {
	l := lingkaran{4}
	fmt.Printf("Luas Lingkaran : %.2f\n", l.luas())
	fmt.Printf("Keliling Lingkaran :%.2f\n", l.keliling())

	p := persegi{2}
	fmt.Println("Luas Persegi :", p.luas())
	fmt.Println("Keliling Persegi :", p.keliling())
}

/* TODO: Interface
adalah kumpulan definisi method yang tidak memiliki isi(hanya definisi saja), yang dibungkus dengan nama tertentu
interface merupakan tipe data seperti polimorphism dalam OPP
*/
