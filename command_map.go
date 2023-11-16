package main

import (
	"fmt"
	"log"
)

func commandMap(c *Config) error {
	res, err := c.pokeapiClient.ListLocationAreas(c.nextUrl)

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
