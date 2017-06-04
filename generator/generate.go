package generator

import (
	"fmt"
	"os"
)

func Run(configFile, outputFile string) {
	config := Config(configFile)

	if len(outputFile) == 0 {
		fmt.Println("-- Writing output to STDOUT")

		run(config, os.Stdout)
	} else {
		fmt.Println("-- Writing output to", outputFile)
		file, err := os.Create(outputFile)

		if err != nil {
			panic(fmt.Errorf("Failed opening output file to write", file, err))
		}

		run(config, file)
	}
}

func run(config []*RssItem, file *os.File) {
	rssXml, _ := GenRss(config)

	file.WriteString(rssXml)
}
