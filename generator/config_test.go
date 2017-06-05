package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	var config []*RssItem

	config = Config("feed.json")

	assert.Equal(3, len(config))

	assert.Equal("ああ", config[0].Title)
	assert.Equal("http://hoge", config[0].Link)

	assert.Equal("いい", config[1].Title)
	assert.Equal("http://fuga", config[1].Link)

	assert.Equal("うう", config[2].Title)
	assert.Equal("http://foo", config[2].Link)
}
