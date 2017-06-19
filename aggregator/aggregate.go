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
		run(config, NewLogger(os.Stdout))
	} else {
		fmt.Println("-- Writing output to", outputFile)
		f, err := os.Create(outputFile)

		if err != nil {
			panic(fmt.Errorf("Failed opening output file to write", f, err))
		}

		run(config, NewLogger(f))
	}
}

func run(urls []string, logger *Logger) {
	defer logger.Close()

	rand.Seed(time.Now().UnixNano())

	for i := range urls {
		j := rand.Intn(i + 1)
		urls[i], urls[j] = urls[j], urls[i]
	}

	httpqueue := HttpQueue{
		Wg:  &sync.WaitGroup{},
		In:  make(chan string),
		Out: make(chan []byte),
	}
	rssqueue := RssQueue{
		Wg:  &sync.WaitGroup{},
		In:  httpqueue.Out,
		Out: make(chan RssItem),
	}
	loggerqueue := LoggerQueue{
		Wg:  &sync.WaitGroup{},
		In:  rssqueue.Out,
		Out: logger,
	}

	for i := 0; i < 4; i++ {
		httpqueue.Wg.Add(1)
		go HttpWorker(i+1, httpqueue)
	}

	for i := 0; i < 4; i++ {
		rssqueue.Wg.Add(1)
		go RssWorker(i+1, rssqueue)
	}

	for i := 0; i < 4; i++ {
		loggerqueue.Wg.Add(1)
		go LoggerWorker(i+1, loggerqueue)
	}

	for _, url := range urls {
		httpqueue.In <- url
	}

	close(httpqueue.In)
	httpqueue.Wg.Wait()

	close(rssqueue.In)
	rssqueue.Wg.Wait()

	close(loggerqueue.In)
	loggerqueue.Wg.Wait()
}
