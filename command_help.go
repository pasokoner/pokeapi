package main

import "fmt"

func commandHelp(c *Config) error {

	commands := getCommands()

	fmt.Println()
	fmt.Println("Welcome to Pokedex")
	fmt.Println("Usage:")
	fmt.Println()

	for _, value := range commands {
		fmt.Printf("%v: %v\n", value.name, value.description)
		fmt.Println()
	}

	return nil
}
