package nyaa

import (
	"github.com/mmcdole/gofeed"
	t "github.com/quantumsheep/go-nyaa/v2/types"
)

type SearchOptions struct {
	Provider string
	Query    string
	Category string
	SortBy   string
	OrderBy  string
	Filter   string
}

var (
	fp = gofeed.NewParser()
)

func Search(opts SearchOptions) ([]t.Torrent, error) {
	url, err := buildURL(opts)
	if err != nil {
		return nil, err
	}

	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}

	res := convertRSS(feed)

	return res, nil
}
