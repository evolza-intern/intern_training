package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var length int
	fmt.Print("Enter desired password length: ")
	fmt.Scanln(&length)

	if length <= 0 {
		fmt.Println("Please enter a positive number.")
		return
	}

	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*"

	password := ""
	for range length {
		index := rand.Intn(len(characters))
		password += string(characters[index])
	}

	fmt.Println("Your random password is:", password)
}
