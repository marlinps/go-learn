package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// TODO:Custom error
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

	//TODO: panic -> informasi error yang ditampilkan adalah stack trace error, setelah ada panic maka program akan berhenti
	var name2 string
	fmt.Printf("Type your name2: ")
	fmt.Scanln(&name2)

	if valid, err := validate(name2); valid {
		fmt.Println("Hello", name2)
	} else {
		panic(err.Error())
		fmt.Println("end process")
	}

	//TODO: recover -> menangkap panic agar program tidak berhenti
	var name3 string
	fmt.Printf("Type your name3: ")
	fmt.Scanln(&name3)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()
	if valid, err := validate(name3); valid {
		fmt.Println("Hello", name3)
	} else {
		panic(err.Error())
	}
}
