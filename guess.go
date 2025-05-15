package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random := rand.Intn(10) + 1
	attempts := 0

	for {
		var guess int
		fmt.Println("Guess the correct number between 1 and 10 (inclusive)")
		fmt.Scanln(&guess)
		attempts++

		if guess < 1 || guess > 10 {
			fmt.Println("Please enter a number between 1 and 10.")
			continue
		}

		if guess == random {
			fmt.Printf("Congratulations! You guessed it right in %d attempts.", attempts)
			break
		} else if guess < random {
			fmt.Println("Too low! Try a higher number : )")
		} else {
			fmt.Println("Too high! Try a lower number  : )")
		}
	}
}
