package main

import (
	"fmt"
	"strconv"
)

func main() {
	// TODO: Error handling
	var input string
	fmt.Printf("Type some something: ")
	fmt.Scanln(&input)

	if number, err := strconv.Atoi(input); err != nil {
		fmt.Println(input, "is not a number")
		fmt.Println("Error:", err)
	} else {
		fmt.Println("You typed number:", number)
	}

}
