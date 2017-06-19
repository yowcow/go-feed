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
	Items []RssItem `xml:"item"`
}

func ParseRss(data []byte) Rss {
	rss := Rss{}
	xml.Unmarshal(data, &rss)
	return rss
}

type RssQueue struct {
	Wg  *sync.WaitGroup
	In  chan []byte
	Out chan RssItem
}

func RssWorker(id int, q RssQueue) {
	name := fmt.Sprintf("[RSS Worker %d]", id)
	defer func() {
		fmt.Println(name, "Exiting")
		q.Wg.Done()
	}()
	for data := range q.In {
		fmt.Println(name, "Got XML to parse")
		rss := ParseRss(data)
		for _, item := range rss.Items {
			q.Out <- item
		}
	}
}
