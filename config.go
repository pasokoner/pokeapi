package main

import "github.com/fhero2030/pokeapi/internal/pokeapi"

type Config struct {
	pokeapiClient pokeapi.Client
	nextUrl       *string
	prevUrl       *string
}
