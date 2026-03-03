package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check for correct number of arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	// Read input file
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Transform the text
	words := strings.Fields(string(content))
	transformedWords := Transform(words)
	finalText := FormatPunctuation(strings.Join(transformedWords, " "))

	// Write to output file
	err = os.WriteFile(outputFileName, []byte(finalText), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
}
