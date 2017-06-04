package generator

import (
	"github.com/gorilla/feeds"
	"time"
)

const FeedTitle string = "Test Feed"
const FeedLink string = "http://test-feed"

func GenRss(items []*RssItem) (string, error) {
	now := time.Now()

	feed := &feeds.Feed{
		Title:   FeedTitle,
		Link:    &feeds.Link{Href: FeedLink},
		Created: now,
	}

	feed.Items = make([]*feeds.Item, len(items))

	for i, item := range items {
		feed.Items[i] = &feeds.Item{
			Title:   item.Title,
			Link:    &feeds.Link{Href: item.Link},
			Created: now,
		}
	}

	return feed.ToRss()
}
