package aggregator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	config := Config("urls_test.json")

	assert.True(assert.EqualValues(
		[]string{
			"https://bo.beaconsco.com/app/feed/rss",
			"https://bo.beaconsco.com/app/feed/rss",
		},
		config,
	))
}
