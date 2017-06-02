package aggregator

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
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

	lines := strings.Split(string(content), "\n")
	validLog := regexp.MustCompile(`^.{6}\s\d{2}\:\d{2}\:\d{2}\.\d{9}\s.+\s\d{3}$`) // Jun  2 09:52:38.382469767 あ 111

	for _, line := range lines[:len(lines)-1] {
		assert.True(validLog.MatchString(line), line)
	}
}
