package main

import (
	"fmt"
)

func main() {
	var temperature float64
	var option int

	fmt.Println("Choose temperature conversion and enter option (1 or 2):\n1: Celsius to Fahrenheit\n2: Fahrenheit to Celsius")
	fmt.Scanln(&option)

	if option != 1 && option != 2 {
		fmt.Println("Invalid option. Please run the program again and select 1 or 2.")
		return
	}

	fmt.Print("Enter temperature value: ")
	fmt.Scanln(&temperature)

	switch option {
	case 1:
		fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", temperature, convertCelsiusToFahrenheit(temperature))
	case 2:
		fmt.Printf("%.2f Fahrenheit is %.2f Celsius\n", temperature, convertFahrenheitToCelsius(temperature))
	}
}

// convertCelsiusToFahrenheit converts Celsius to Fahrenheit
func convertCelsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

// convertFahrenheitToCelsius converts Fahrenheit to Celsius
func convertFahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}
