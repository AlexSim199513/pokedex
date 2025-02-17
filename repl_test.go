package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{input: " Hello World ", expected: []string{"hello", "world"}},
		{input: "Charmander Bulbasaur PIKACHU", expected: []string{"charmander", "bulbasaur", "pikachu"}},
		{input: "ChaRiZard", expected: []string{"charizard"}},
		{input: "Squirtle Wartortle Blastoise", expected: []string{"squirtle", "wartortle", "blastoise"}},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("unexpected slice length for input '%v': got '%d', expected %d", c.input, len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("unexpected word at index %d for input '%v': got '%v', expected '%v'", i, c.input, actual[i], c.expected[i])
			}
		}
	}
}
