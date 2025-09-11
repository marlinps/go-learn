package library

import "fmt"

type Student struct { // TODO: Exported
	Name  string
	Grade int
}

func SayHello(name string) { // TODO: Exported
	fmt.Println("Hello")
	introduce(name)
}

func introduce(name string) { // TODO: Unexported
	fmt.Println("My Name is", name)
}
