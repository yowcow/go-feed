package aggregator

import (
	"fmt"
	"math/rand"
	"os"
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

func worker(id int, logging *Logging, jobChan chan string, resChan chan int) {
	for url := range jobChan {
		fmt.Printf("Worker (%d): Start working on %s\n", id, url)

		rssXml, err := HttpGet(url)

		if err == nil {
			fmt.Printf("Worker (%d): Successfully got response\n", id)

			rssData := ParseRss(rssXml)

			fmt.Printf("Worker (%d): Successfully parsed RSS\n", id)

			count := 0

			for _, item := range rssData.Items {
				logging.Log(item.Link, item.Title)
				count += 1
			}

			fmt.Printf("Worker (%d): Successfully logged %d items\n", id, count)

			resChan <- count
		} else {
			fmt.Printf("Worker (%d): Failed getting response from %s (%s)\n", id, url, err)

			resChan <- -1
		}
	}
}

func run(urls []string, logging *Logging) {
	defer logging.Close()

	jobChan := make(chan string, len(urls))
	resChan := make(chan int, len(urls))

	for w := 1; w <= 4; w++ {
		go worker(w, logging, jobChan, resChan)
	}

	rand.Seed(time.Now().UnixNano())

	for i := range urls {
		j := rand.Intn(i + 1)
		urls[i], urls[j] = urls[j], urls[i]
	}

	for _, url := range urls {
		jobChan <- url
	}

	close(jobChan)

	for i := 0; i < len(urls); i++ {
		fmt.Printf("Main (%d): Logged item count %d\n", i, <-resChan)
	}
}
