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

func TestLogger(t *testing.T) {
	assert := assert.New(t)

	f, e := ioutil.TempFile("", "test-log")
	defer os.Remove(f.Name())

	assert.Nil(e)

	logger := NewLogger(f)
	logger.Log("あ", "111")
	logger.Log("い", "222")
	logger.Log("foo", "333")
	logger.Log("bar", "444")
	logger.Close()

	content, _ := ioutil.ReadFile(f.Name())

	lines := strings.Split(string(content), "\n")
	validLog := regexp.MustCompile(`^.{6}\s\d{2}\:\d{2}\:\d{2}\.\d{9}\s.+\s\d{3}$`) // Jun  2 09:52:38.382469767 あ 111

	for _, line := range lines[:len(lines)-1] {
		assert.True(validLog.MatchString(line), line)
	}
}

func TestLoggerWorker(t *testing.T) {
	assert := assert.New(t)

	f, e := ioutil.TempFile("", "test-log")
	defer os.Remove(f.Name())

	assert.Nil(e)

	logqueue := LoggerQueue{
		Wg:  &sync.WaitGroup{},
		In:  make(chan RssItem),
		Out: NewLogger(f),
	}

	for i := 0; i < 6; i++ {
		logqueue.Wg.Add(1)
		go LoggerWorker(i+1, logqueue)
	}

	for i := 0; i < 20; i++ {
		logqueue.In <- RssItem{fmt.Sprintf("Title%d", i), ""}
	}

	close(logqueue.In)
	logqueue.Wg.Wait()

	content, _ := ioutil.ReadFile(f.Name())
	lines := strings.Split(string(content), "\n")

	assert.Equal(20+1, len(lines))
}
