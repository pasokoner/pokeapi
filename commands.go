package main

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List 20 Locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List 20 Locations of the previous page",
			callback:    commandMapB,
		},
	}
}
