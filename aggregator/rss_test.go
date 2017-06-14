package aggregator

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rssXml string = `
<?xml version="1.0" encoding="UTF-8"?>
<rdf:RDF>
  <item>
    <title>あああ</title>
    <link>http://foobar.com</link>
  </item>
  <item>
    <title>いいい</title>
    <link>http://hogefuga.com</link>
  </item>
</rdf:RDF>
`

func TestParseRss(t *testing.T) {
	assert := assert.New(t)

	data := ParseRss([]byte(rssXml))

	assert.Equal(2, len(data.Items))
	assert.Equal("あああ", data.Items[0].Title)
	assert.Equal("http://foobar.com", data.Items[0].Link)
	assert.Equal("いいい", data.Items[1].Title)
	assert.Equal("http://hogefuga.com", data.Items[1].Link)
}

func TestRssParserWorker(t *testing.T) {
	wg := &sync.WaitGroup{}
	iq := make(chan []byte)
	oq := make(chan *RssItem)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go RssParserWorker(i+1, wg, iq, oq)
	}

	testmutex := &sync.Mutex{}
	testwg := &sync.WaitGroup{}
	count := 0

	for i := 0; i < 2; i++ {
		testwg.Add(1)
		go func() {
			defer testwg.Done()
			for {
				_, ok := <-oq
				if !ok {
					return
				}
				testmutex.Lock()
				count += 1
				testmutex.Unlock()
			}
		}()
	}

	for i := 0; i < 10; i++ {
		iq <- []byte(rssXml)
	}

	close(iq)
	wg.Wait()

	close(oq)
	testwg.Wait()

	assert.Equal(t, 20, count)
}
