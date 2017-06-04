package main

import (
	"flag"
	"fmt"
	"github.com/yowcow/go-feed/generator"
	"os"
)

var (
	configFile string
	outputFile string
)

func main() {
	flag.StringVar(&configFile, "config", "feed.json", "Path to JSON file containing a list of items")
	flag.StringVar(&outputFile, "output", "", "Path to output file")
	flag.Parse()

	if _, err := os.Stat(configFile); err != nil {
		panic(fmt.Errorf("Config file \"%s\" does not exist", configFile))
	}

	generator.Run(configFile, outputFile)
}
