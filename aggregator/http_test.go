package aggregator

import (
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func startHttpServer(wg *sync.WaitGroup, ch chan struct{}) {
	srvmux := &http.ServeMux{}
	srvmux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Header.Get("User-Agent"))
	})
	server := &http.Server{
		Addr:    ":8888",
		Handler: srvmux,
	}

	defer func() {
		server.Close()
		wg.Done()
	}()

	go func() {
		server.ListenAndServe()
	}()

	for _ = range ch {
	}
}

func TestHttpGet(t *testing.T) {
	assert := assert.New(t)

	wg := &sync.WaitGroup{}
	ch := make(chan struct{})

	wg.Add(1)
	go startHttpServer(wg, ch)

	body, err := HttpGet("http://localhost:8888/")

	close(ch)
	wg.Wait()

	assert.Nil(err)
	assert.Equal("GoClient/0.1", string(body))
}

func TestHttpWorker(t *testing.T) {
	serverwg := &sync.WaitGroup{}
	serverch := make(chan struct{})

	serverwg.Add(1)
	go startHttpServer(serverwg, serverch)

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
		httpqueue.In <- "http://localhost:8888/"
	}

	close(httpqueue.In)
	httpqueue.Wg.Wait()

	close(httpqueue.Out)
	testqueue.Wg.Wait()

	close(serverch)
	serverwg.Wait()

	assert.Equal(t, 10, testqueue.Count)
}
