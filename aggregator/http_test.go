package aggregator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpGet(t *testing.T) {
	assert := assert.New(t)

	body, err := HttpGet("http://www.beaconsco.com")

	assert.Nil(err)
	assert.True(len(body) > 0)
}
