package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// TODO: Custom error
func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

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

	// TODO: Custom error usage
	var name string
	fmt.Printf("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hello", name)
	} else {
		fmt.Println(err.Error())
	}

	// TODO: panic -> informasi error yang ditampilkan adalah stack trace error
	var name2 string
	fmt.Printf("Type your name2: ")
	fmt.Scanln(&name2)

	if valid, err := validate(name2); valid {
		fmt.Println("Hello", name2)
	} else {
		panic(err.Error())
		fmt.Println("end process")
	}
}
