package main

import (
	"errors"
	"fmt"
	"log"
)

func commandMapB(c *Config) error {
	if c.prevUrl == nil {
		return errors.New("You are on the first page")
	}

	res, err := c.pokeapiClient.ListLocationAreas(c.prevUrl)

	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, location := range res.Results {
		fmt.Printf(" - %s\n", location.Name)
	}

	c.nextUrl = res.Next
	c.prevUrl = res.Previous

	return nil
}
