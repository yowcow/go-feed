package aggregator

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
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

	assert.True(len(content) > 0)
}
