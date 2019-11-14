package main

import (
	"fmt"
	"flag"
	"os"
	"bourbonfinder"
)

func main() {

	var configFileName string
	flag.StringVar(&configFileName, "config", "", "Path to the config JSON file. Required.")
	flag.StringVar(&configFileName,"c", "", "Path to the config JSON file (shorthand). Required.")
	flag.Parse()

	if configFileName == "" {
		usage()
		os.Exit(1)
	}

	//bourbonfinder.Search("data/config.json")
	bourbonfinder.Search(configFileName)
}

func usage() {
	fmt.Println("Usage: bourbon-finder <options>")
	flag.PrintDefaults()
}
