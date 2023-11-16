package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/fhero2030/pokeapi/internal/pokeapi"
)

func main() {

	cfg := Config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		nextUrl:       nil,
		prevUrl:       nil,
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Poke CLI > ")
		if scanner.Scan() {
			commands := getCommands()
			text := scanner.Text()

			c, ok := commands[text]

			if !ok {
				fmt.Println("Command does not exist")
				continue
			}

			err := c.callback(&cfg)

			if err != nil {
				fmt.Println(err)
			}

		} else if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			break
		}
	}

}
