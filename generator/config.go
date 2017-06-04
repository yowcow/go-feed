package generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type RssItem struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func Config(file string) []*RssItem {
	content, err := ioutil.ReadFile(file)

	if err != nil {
		panic(fmt.Errorf("Cannot read from file", err))
	}

	config := []*RssItem{}
	err = json.Unmarshal(content, &config)

	if err != nil {
		panic(fmt.Errorf("Cannot parse JSON from file", err))
	}

	return config
}
