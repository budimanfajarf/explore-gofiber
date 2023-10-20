package utils

import "strings"

func CalculateOffset(page, limit int) int {
	return (page - 1) * limit
}

func SnakeCaseToWords(input string) string {
	words := strings.Split(input, "_")
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, " ")
}
