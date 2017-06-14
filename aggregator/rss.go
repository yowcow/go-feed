package aggregator

import (
	"encoding/xml"
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
	for {
		data, ok := <-iq
		if !ok {
			return
		}
		rss := ParseRss(data)
		for _, item := range rss.Items {
			oq <- item
		}
	}
}
