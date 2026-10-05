package main

import (
	"fmt"
	"strings"
)

func longestWord(sentence string) string {
	words := strings.Fields(sentence)

	longest := ""

	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}

	return longest
}

func main() {
	var sentence string

	fmt.Print("Введите предложение: ")
	fmt.Scanln(&sentence)

	sentence = strings.TrimSpace(sentence)
	fmt.Println(longestWord(sentence))
}