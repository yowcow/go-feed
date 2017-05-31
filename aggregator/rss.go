package aggregator

import (
	"encoding/xml"
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
