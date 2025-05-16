package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the input file
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	lineCount := 0
	wordToFind := "go"
	wordCount := 0

	// Read each line
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		// Count the word occurrences in the line (case-insensitive)
		wordCount += strings.Count(strings.ToLower(line), wordToFind)
	}

	// Prepare the result
	result := fmt.Sprintf("Number of lines: %d\nOccurrences of '%s': %d\n", lineCount, wordToFind, wordCount)

	// Write result to output file
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	outputFile.WriteString(result)
}
