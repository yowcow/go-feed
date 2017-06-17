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

func HttpWorker(id int, wg *sync.WaitGroup, iq chan string, oq chan []byte) {
	defer wg.Done()
	name := fmt.Sprintf("[HTTP Worker %d]", id)
	for {
		url, ok := <-iq
		if !ok {
			fmt.Println(name, "Exiting")
			return
		}
		fmt.Println(name, "Got URL", url)
		body, err := HttpGet(url)
		if err != nil {
			panic(err)
		}
		oq <- body
	}
}
