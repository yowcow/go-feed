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

func TestRssWorker(t *testing.T) {
	rssqueue := RssQueue{
		Wg:  &sync.WaitGroup{},
		In:  make(chan []byte),
		Out: make(chan RssItem),
	}

	for i := 0; i < 4; i++ {
		rssqueue.Wg.Add(1)
		go RssWorker(i+1, rssqueue)
	}

	type TestQueue struct {
		Wg    *sync.WaitGroup
		Mutex *sync.Mutex
		Count int
	}
	testqueue := TestQueue{&sync.WaitGroup{}, &sync.Mutex{}, 0}

	for i := 0; i < 2; i++ {
		testqueue.Wg.Add(1)
		go func(q RssQueue, tq *TestQueue) {
			defer tq.Wg.Done()
			for {
				_, ok := <-q.Out
				if !ok {
					return
				}
				tq.Mutex.Lock()
				tq.Count += 1
				tq.Mutex.Unlock()
			}
		}(rssqueue, &testqueue)
	}

	for i := 0; i < 10; i++ {
		rssqueue.In <- []byte(rssXml)
	}

	close(rssqueue.In)
	rssqueue.Wg.Wait()

	close(rssqueue.Out)
	testqueue.Wg.Wait()

	assert.Equal(t, 20, testqueue.Count)
}
