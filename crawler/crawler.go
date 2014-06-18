package crawler

import (
	"fmt"

	"github.com/wkirschbaum/webcrawler/fetcher"
)

type Crawler struct {
	Found   map[string]int
	Fetched chan fetcher.FetchedResult
}

func (crawler *Crawler) Crawl(url string, depth int, fetcher fetcher.Fetcher) {
	if depth <= 0 {
		return
	}

	if _, ok := crawler.Found[url]; ok {
		return
	}

	result, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	crawler.Fetched <- result
	crawler.Found[url] = result.Status

	fmt.Printf("crawling %d urls from %s\n", len(result.Urls), url)
	for _, u := range result.Urls {
		go crawler.Crawl(u, depth-1, fetcher)
	}
	return
}
