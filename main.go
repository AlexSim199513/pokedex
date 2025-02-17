package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Hello, World!\n")
}

func CleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowercased := strings.ToLower(trimmed)
	cleaned := strings.Fields(lowercased)
	return cleaned
}
