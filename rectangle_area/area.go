package main

import "fmt"

func main() {
	var length, width int

	fmt.Println("Enter length of rectangle.")
	fmt.Scanln(&length)

	fmt.Println("Enter width of rectangle.")
	fmt.Scanln(&width)

	fmt.Println("The area of the rectangle is", calculateArea(length, width))
}

// calculateArea returns the area of a rectangle given its length and width
func calculateArea(length, width int) int {
	return length * width
}
