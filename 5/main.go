package main

import (
	"bufio"
	"fmt"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите предложение: ")
	scanner.Scan()

	sentence := scanner.Text()

	fmt.Println(longestWord(sentence))
}
