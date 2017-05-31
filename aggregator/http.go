package aggregator

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) ([]byte, error) {
	fmt.Println("Making GET request to", url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Request failed", url, err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Reading response body failed", url, err)
		return nil, err
	}

	fmt.Println("Done making request to", url)

	return body, nil
}
