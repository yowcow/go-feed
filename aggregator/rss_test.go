package aggregator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRss(t *testing.T) {
	assert := assert.New(t)

	rss := `
	<?xml version="1.0" encoding="UTF-8"?>
	<rdf:RDF>
	  <item>
	    <title>あああ</title>
	    <link>http://foobar.com</link>
	  </item>
	  <item>
	    <title>いいい</title>
	    <link>http://hogefuga.com</link>
	  </item>
	</rdf:RDF>
	`

	data := ParseRss([]byte(rss))

	assert.Equal(2, len(data.Items))
	assert.Equal("あああ", data.Items[0].Title)
	assert.Equal("http://foobar.com", data.Items[0].Link)
	assert.Equal("いいい", data.Items[1].Title)
	assert.Equal("http://hogefuga.com", data.Items[1].Link)
}
