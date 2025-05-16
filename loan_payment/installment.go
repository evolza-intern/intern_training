package main

import (
	"fmt"
	"math"
)

func main() {
	var principal, annualRate float64
	var months int

	fmt.Println("Loan Calculator")

	fmt.Print("Enter loan amount: ")
	fmt.Scanln(&principal)

	fmt.Print("Enter annual interest rate (e.g., 5 for 5%): ")
	fmt.Scanln(&annualRate)

	fmt.Print("Enter payment duration in months: ")
	fmt.Scanln(&months)

	monthlyPayment := calculateMonthlyPayment(principal, annualRate, months)

	fmt.Printf("\n Monthly Payment: %.2f\n", monthlyPayment)
}

// calculateMonthlyPayment calculates the value of a monthly installment
func calculateMonthlyPayment(principal, annualRate float64, months int) float64 {
	r := annualRate / 12 / 100
	n := float64(months)

	if r == 0 {
		return principal / float64(months)
	}

	payment := principal * r * math.Pow(1+r, n) / (math.Pow(1+r, n) - 1)
	return payment
}
