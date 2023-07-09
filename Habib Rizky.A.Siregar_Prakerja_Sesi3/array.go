package main

import "fmt"

func main() {
	input := "prakerja golang"
	filtered := filterLetter(input, 'a')

	fmt.Println(filtered)
}

func filterLetter(input string, letter rune) string {
	count := 0
	filtered := make([]rune, len(input))

	for _, char := range input {
		if char != letter || count == 1 {
			filtered = append(filtered, char)
		}
		if char == letter && count == 0 {
			count++
		}
	}

	return string(filtered)
}
