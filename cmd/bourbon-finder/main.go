package main

import (
	"bourboncommon"
	"bourbonfinder"
	"bourbonweb"
	"flag"
	"fmt"
	"os"
)

func main() {

	var configFileName string
	var startWeb bool

	flag.StringVar(&configFileName, "c", "", "Path to the config JSON file. Required.")
	flag.BoolVar(&startWeb, "w", false, "Start web server.")
	flag.Parse()

	if configFileName == "" {
		usage()
		os.Exit(1)
	}

	config, err := bourboncommon.OpenConfig(configFileName)
	if err != nil {
		fmt.Println("Error: Could not open config file " + configFileName)
		fmt.Println(err)
		os.Exit(1)
	}

	if startWeb {
		bourbonweb.StartWebServer("0.0.0.0:5001", config)
	} else {
		results, _ := bourbonfinder.Search(config)
		m := bourbonfinder.GroupByStore(results)
		bourbonfinder.PrintGroup(m, os.Stdout, true)
	}
}

func usage() {
	fmt.Println("Usage: bourbon-finder <options>")
	flag.PrintDefaults()
}
