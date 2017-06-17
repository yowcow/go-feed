package aggregator

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogging(t *testing.T) {
	assert := assert.New(t)

	f, e := ioutil.TempFile("", "test-log")
	defer os.Remove(f.Name())

	assert.Nil(e)

	logging := NewLogging(f)
	logging.Log("あ", "111")
	logging.Log("い", "222")
	logging.Log("foo", "333")
	logging.Log("bar", "444")
	logging.Close()

	content, _ := ioutil.ReadFile(f.Name())

	lines := strings.Split(string(content), "\n")
	validLog := regexp.MustCompile(`^.{6}\s\d{2}\:\d{2}\:\d{2}\.\d{9}\s.+\s\d{3}$`) // Jun  2 09:52:38.382469767 あ 111

	for _, line := range lines[:len(lines)-1] {
		assert.True(validLog.MatchString(line), line)
	}
}

func TestLoggingWorker(t *testing.T) {
	assert := assert.New(t)

	f, e := ioutil.TempFile("", "test-log")
	defer os.Remove(f.Name())

	assert.Nil(e)

	logging := NewLogging(f)
	wg := &sync.WaitGroup{}
	q := make(chan *RssItem)

	for i := 0; i < 6; i++ {
		wg.Add(1)
		go LoggingWorker(i+1, wg, q, logging)
	}

	for i := 0; i < 20; i++ {
		q <- &RssItem{fmt.Sprintf("Title%d", i), ""}
	}

	close(q)
	wg.Wait()

	content, _ := ioutil.ReadFile(f.Name())
	lines := strings.Split(string(content), "\n")

	assert.Equal(20+1, len(lines))
}
