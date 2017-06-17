package aggregator

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func Run(configFile, outputFile string) {
	fmt.Println("-- Using config from", configFile)

	config := Config(configFile)

	if len(outputFile) == 0 {
		fmt.Println("-- Writing output to STDOUT")
		run(config, NewLogging(os.Stdout))
	} else {
		fmt.Println("-- Writing output to", outputFile)
		f, err := os.Create(outputFile)

		if err != nil {
			panic(fmt.Errorf("Failed opening output file to write", f, err))
		}

		run(config, NewLogging(f))
	}
}

func run(urls []string, logging *Logging) {
	defer logging.Close()

	rand.Seed(time.Now().UnixNano())

	for i := range urls {
		j := rand.Intn(i + 1)
		urls[i], urls[j] = urls[j], urls[i]
	}

	httpwg := &sync.WaitGroup{}
	rsswg := &sync.WaitGroup{}
	loggingwg := &sync.WaitGroup{}

	feedurlch := make(chan string)
	httpbodych := make(chan []byte)
	rssitemch := make(chan *RssItem)

	for i := 0; i < 4; i++ {
		httpwg.Add(1)
		go HttpWorker(i+1, httpwg, feedurlch, httpbodych)
	}

	for i := 0; i < 4; i++ {
		rsswg.Add(1)
		go RssParserWorker(i+1, rsswg, httpbodych, rssitemch)
	}

	for i := 0; i < 4; i++ {
		loggingwg.Add(1)
		go LoggingWorker(i+1, loggingwg, rssitemch, logging)
	}

	for _, url := range urls {
		feedurlch <- url
	}

	close(feedurlch)
	httpwg.Wait()

	close(httpbodych)
	rsswg.Wait()

	close(rssitemch)
	loggingwg.Wait()
}
