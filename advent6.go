package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const file = "inputs.txt"

func main() {
	// Open the file
	file, _ := os.Open(file)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	input := []rune(scanner.Text())

	log.Println(calculateResult(input, 4))
	log.Println(calculateResult(input, 14))
}

func calculateResult(input []rune, windowSize int) int {
	var window []rune

	for i, val := range input {
		if len(window) == windowSize {
			window = window[1:]
		}
		window = append(window, val)

		if len(window) == windowSize {
			for j, val := range window {
				if strings.Count(string(window), string(val)) != 1 {
					break
				}
				if j == windowSize-1 {
					log.Println(string(window))
					return i + windowSize - j
				}
			}
		}
	}
	return -1
}
