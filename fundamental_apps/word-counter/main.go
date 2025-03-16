package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// put your text in data.txt file
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
	fmt.Println("File read successfully")

	fmt.Println("\nStatistics: ")
	fmt.Println("Words: ", countWords(string(data)))
	fmt.Println("Lines: ", countLines(string(data)))
	fmt.Println("Characters: ", countCharacters(string(data)))
}

func countWords(data string) int {
	return len(strings.Fields(data))
}

func countLines(data string) int {
	return len(strings.Split(data, "\n"))
}

func countCharacters(data string) int {
	return len(data)
}
