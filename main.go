package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	initCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		input := scanner.Text()

		cleanedInput := CleanInput(input)

		if len(cleanedInput) == 0 {
			fmt.Println("No command found")
			continue
		}

		command := cleanedInput[0]

		if cmd, ok := commands[command]; ok {
			if err := cmd.callback(); err != nil {
				fmt.Println("Error: ", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func CleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowercased := strings.ToLower(trimmed)
	cleaned := strings.Fields(lowercased)
	return cleaned
}

type CliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]CliCommand

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func initCommands() {
	commands = map[string]CliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Show avaiable commands",
			callback:    help,
		},
	}
}
