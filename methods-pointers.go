package main

import "fmt"

type Rectangle struct {
	Length float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}

// TODO: delete * untuk melihat perbedaan method pointers dengan value receivers
func (r *Rectangle) Scale(s float32) {
	r.Length *= s
	r.Width *= s
}

func Scale(r *Rectangle, s float32) {
	r.Length *= s
	r.Width *= s
}

func main() {
	rect := Rectangle{10, 5}

	// TODO: method with pointer receiver
	rect.Scale(2)
	area := rect.Area()

	// TODO: function with parameter pointer
	Scale(&rect, 2)
	area := rect.Area()

	fmt.Println("Area of Rectangle:", area)
}
