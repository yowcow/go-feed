package aggregator

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func HttpGet(url string) ([]byte, error) {
	client := new(http.Client)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "GoClient/0.1")
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

type HttpQueue struct {
	Wg  *sync.WaitGroup
	In  chan string
	Out chan []byte
}

func HttpWorker(id int, q HttpQueue) {
	name := fmt.Sprintf("[HTTP Worker %d]", id)
	defer func() {
		fmt.Println(name, "Exiting")
		q.Wg.Done()
	}()
	for url := range q.In {
		fmt.Println(name, "Got URL", url)
		body, err := HttpGet(url)
		if err != nil {
			panic(err)
		}
		q.Out <- body
	}
}
