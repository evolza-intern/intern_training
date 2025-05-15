package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter text to analyze")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	words := countWords(text)
	withSpaces, withoutSpaces := countCharacters(text)

	fmt.Printf("\n📊 Analysis Results:\n")
	fmt.Printf("Total words: %d\n", words)
	fmt.Printf("Total characters (with spaces): %d\n", withSpaces)
	fmt.Printf("Total characters (without spaces): %d\n", withoutSpaces)

	// Ask for a specific letter to count
	fmt.Print("\nEnter a letter to count its frequency: ")
	letterInput, _ := reader.ReadString('\n')
	letterInput = strings.TrimSpace(letterInput)

	if len(letterInput) == 1 {
		letter := rune(letterInput[0])
		letterCount := countLetter(text, letter)
		fmt.Printf("Occurrences of '%c': %d\n", letter, letterCount)
	} else {
		fmt.Println("Please enter only a single character.")
	}
}

// countWords returns the number of words in a string
func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

// countCharacters returns total characters including and excluding spaces
func countCharacters(text string) (withSpaces, withoutSpaces int) {
	for _, ch := range text {
		withSpaces++
		if !unicode.IsSpace(ch) {
			withoutSpaces++
		}
	}
	return
}

// countLetter returns how many times a given letter appears in the text (case-insensitive)
func countLetter(text string, letter rune) int {
	count := 0
	for _, ch := range strings.ToLower(text) {
		if ch == unicode.ToLower(letter) {
			count++
		}
	}
	return count
}
