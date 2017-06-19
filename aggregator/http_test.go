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
	httpqueue := HttpQueue{
		Wg:  &sync.WaitGroup{},
		In:  make(chan string),
		Out: make(chan []byte),
	}

	for i := 0; i < 4; i++ {
		httpqueue.Wg.Add(1)
		go HttpWorker(i+1, httpqueue)
	}

	type TestQueue struct {
		Wg    *sync.WaitGroup
		Mutex *sync.Mutex
		Count int
	}
	testqueue := TestQueue{&sync.WaitGroup{}, &sync.Mutex{}, 0}

	for i := 0; i < 2; i++ {
		testqueue.Wg.Add(1)
		go func(q HttpQueue, tq *TestQueue) {
			defer testqueue.Wg.Done()
			for {
				_, ok := <-q.Out
				if !ok {
					return
				}
				tq.Mutex.Lock()
				tq.Count += 1
				tq.Mutex.Unlock()
			}
		}(httpqueue, &testqueue)
	}

	for i := 0; i < 10; i++ {
		httpqueue.In <- "http://www.beaconsco.com/"
	}

	close(httpqueue.In)
	httpqueue.Wg.Wait()

	close(httpqueue.Out)
	testqueue.Wg.Wait()

	assert.Equal(t, 10, testqueue.Count)
}
