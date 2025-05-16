package main

import (
	"fmt"
	"unicode"
)

func main() {

	var password string
	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	hasUpper := false
	hasLower := false
	hasNumber := false
	isLong := len(password) >= 8

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		}
		if unicode.IsLower(char) {
			hasLower = true
		}
		if unicode.IsDigit(char) {
			hasNumber = true
		}
	}

	if isLong && hasUpper && hasLower && hasNumber {
		fmt.Println("Strong password!")
	} else {
		fmt.Println(" Weak password. Make sure it has:")
		if !isLong {
			fmt.Println("- At least 8 characters")
		}
		if !hasUpper {
			fmt.Println("- At least one uppercase letter")
		}
		if !hasLower {
			fmt.Println("- At least one lowercase letter")
		}
		if !hasNumber {
			fmt.Println("- At least one number")
		}
	}
}
