package aggregator

import (
	"encoding/xml"
	"fmt"
	"sync"
)

type RssItem struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

type Rss struct {
	Items []*RssItem `xml:"item"`
}

func ParseRss(data []byte) *Rss {
	rss := Rss{}
	xml.Unmarshal(data, &rss)
	return &rss
}

func RssParserWorker(id int, wg *sync.WaitGroup, iq chan []byte, oq chan *RssItem) {
	defer wg.Done()
	name := fmt.Sprintf("[RSS Parser Worker %d]", id)
	for {
		data, ok := <-iq
		if !ok {
			fmt.Println(name, "Exiting")
			return
		}
		fmt.Println(name, "Got XML to parse")
		rss := ParseRss(data)
		for _, item := range rss.Items {
			oq <- item
		}
	}
}
