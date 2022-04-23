package nyaa

import (
	"github.com/mmcdole/gofeed"
	t "github.com/quantumsheep/go-nyaa/v2/types"
)

func convertRSS(feed *gofeed.Feed) []t.Torrent {
	var res []t.Torrent

	for _, item := range feed.Items {
		res = append(
			res,
			t.Torrent{
				Name:        item.Title,
				Link:        item.Link,
				Date:        item.Published,
				Description: item.Description,
				GUID:        item.GUID,
				Comments:    item.Extensions["nyaa"]["comments"][0].Value,
				IsTrusted:   item.Extensions["nyaa"]["trusted"][0].Value,
				IsRemake:    item.Extensions["nyaa"]["remake"][0].Value,
				Size:        item.Extensions["nyaa"]["size"][0].Value,
				Seeders:     item.Extensions["nyaa"]["seeders"][0].Value,
				Leechers:    item.Extensions["nyaa"]["leechers"][0].Value,
				Downloads:   item.Extensions["nyaa"]["downloads"][0].Value,
				Category:    item.Extensions["nyaa"]["category"][0].Value,
				CategoryID:  item.Extensions["nyaa"]["categoryId"][0].Value,
				InfoHash:    item.Extensions["nyaa"]["infoHash"][0].Value,
			},
		)
	}
	return res
}
