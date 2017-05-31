package aggregator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Config(filename string) []string {
	c := []string{}

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(fmt.Errorf("Cannot read from file", err))
	}

	err = json.Unmarshal(content, &c)

	if err != nil {
		panic(fmt.Errorf("Cannot parse JSON from file", err))
	}

	return c
}
