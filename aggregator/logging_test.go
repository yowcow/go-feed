package aggregator

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
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

	assert.True(len(content) > 0)
}
