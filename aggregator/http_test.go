package aggregator

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestHttpGet(t *testing.T) {
	assert := assert.New(t)

	body, err := HttpGet("http://www.beaconsco.com")

	assert.Nil(err)
	assert.True(len(body) > 0)
}

func TestHttpWorker(t *testing.T) {
	wg := &sync.WaitGroup{}
	iq := make(chan string)
	oq := make(chan []byte)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go HttpWorker(i+1, wg, iq, oq)
	}

	testwg := &sync.WaitGroup{}
	testmutex := &sync.Mutex{}
	count := 0

	for i := 0; i < 2; i++ {
		testwg.Add(1)
		go func(w *sync.WaitGroup, m *sync.Mutex, c *int) {
			defer w.Done()
			for {
				_, ok := <-oq
				if !ok {
					return
				}
				m.Lock()
				*c += 1
				m.Unlock()
			}
		}(testwg, testmutex, &count)
	}

	for i := 0; i < 10; i++ {
		iq <- "http://www.beaconsco.com/"
	}

	close(iq)
	wg.Wait()

	close(oq)
	testwg.Wait()

	assert.Equal(t, 10, count)
}
