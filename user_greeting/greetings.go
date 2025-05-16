package main

import "fmt"

func main() {
	var userName string

	fmt.Println("Enter your name.")
	fmt.Scanln(&userName)

	fmt.Println("Hello,", userName)
}
