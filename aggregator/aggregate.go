package aggregator

import (
	"fmt"
	"os"
	"sync"
)

func Run(configFile, outputFile string) {
	fmt.Println("Using config from", configFile)
	config := Config(configFile)

	if len(outputFile) == 0 {
		fmt.Println("Write output to STDOUT")
		run(config, NewLogger(os.Stdout))
	} else {
		fmt.Println("Write output to", outputFile)
		f, err := os.Create(outputFile)

		if err != nil {
			panic(fmt.Errorf("Failed opening output file to write", f, err))
		}

		run(config, NewLogger(f))
	}
}

func run(urls []string, logger *Log) {
	wg := &sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go func(u string, w *sync.WaitGroup) {
			fmt.Println("Begin working on", u)
			rssXml, err := HttpGet(u)
			if err == nil {
				rssData := ParseRss(rssXml)
				for _, item := range rssData.Items {
					logger.Log(item.Link, item.Title)
				}
				fmt.Println("Done working on", u)
			} else {
				fmt.Println("Failed working on", u)
			}
			w.Done()
		}(url, wg)
	}

	wg.Wait()
}
