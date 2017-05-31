package main

import (
	"flag"
	"fmt"
	"github.com/yowcow/go-feed/aggregator"
	"os"
)

var (
	configFile string
	outputFile string
)

func main() {
	flag.StringVar(&configFile, "config", "urls.json", "Path to JSON file containing a list of URLs")
	flag.StringVar(&outputFile, "output", "", "Path to output file")
	flag.Parse()

	if _, err := os.Stat(configFile); err != nil {
		panic(fmt.Errorf("Config file \"%s\" does not exist", configFile))
	}

	aggregator.Run(configFile, outputFile)
}
