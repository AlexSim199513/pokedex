package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	initCommands()
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{}

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
			if err := cmd.callback(cfg); err != nil {
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
	callback    func(*Config) error
}

var commands map[string]CliCommand

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"

	if cfg.Next != "" {
		url = cfg.Next
	}

	locations, err := getLocationAreas(url)
	if err != nil {
		return err
	}

	cfg.Next = locations.Next
	if locations.Previous != nil {
		cfg.Previous = *locations.Previous
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page.")
		return nil
	}

	locations, err := getLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = locations.Next
	if locations.Previous != nil {
		cfg.Previous = *locations.Previous
	}

	if locations.Previous == nil {
		cfg.Previous = ""
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
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

		"map": {
			name:        "map",
			description: "Returns a list of locations",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Returns the previous section of locations",
			callback:    commandMapB,
		},
	}
}

type Config struct {
	Next     string
	Previous string
}

type locationResponse struct {
	Count    int        `json:"count"`
	Results  []location `json:"results"`
	Next     string     `json:"next"`
	Previous *string    `json:"previous"`
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getLocationAreas(url string) (locationResponse, error) {

	var locations locationResponse

	res, err := http.Get(url)

	if err != nil {
		return locations, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return locations, fmt.Errorf("Status Code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locations, err
	}

	err = json.Unmarshal(body, &locations)
	if err != nil {
		return locations, err
	}

	return locations, nil
}
