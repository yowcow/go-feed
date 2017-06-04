package generator

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	assert := assert.New(t)

	items := []*RssItem{
		&RssItem{"ほげ", "http://hoge"},
		&RssItem{"ふが", "http://fuga"},
	}

	rssXml, e := GenRss(items)

	assert.Nil(e)

	type RssItem struct {
		Title string `xml:"title"`
		Link  string `xml:"link"`
	}

	type RssChannel struct {
		Items []*RssItem `xml:"item"`
	}

	type Rss struct {
		Channel *RssChannel `xml:"channel"`
	}

	rssData := &Rss{}
	xml.Unmarshal([]byte(rssXml), rssData)

	rssItems := rssData.Channel.Items

	assert.Equal("ほげ", rssItems[0].Title)
	assert.Equal("http://hoge", rssItems[0].Link)

	assert.Equal("ふが", rssItems[1].Title)
	assert.Equal("http://fuga", rssItems[1].Link)
}
