package main

import (
	"bourbonfinder"
	"log"
)

func main() {
	log.Println("Bourbon Finder")
	bourbonfinder.Search("data/config.json")
}
