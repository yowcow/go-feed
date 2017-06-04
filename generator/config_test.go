package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfig(t *testing.T) {
	assert := assert.New(t)

	config := Config("feed.json")

	assert.Equal(3, len(config))
	assert.Equal("ああ", config[0].Title)
	assert.Equal("いい", config[1].Title)
	assert.Equal("うう", config[2].Title)
}
