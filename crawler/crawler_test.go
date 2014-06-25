package crawler

import (
	"errors"
	"fmt"
	"testing"

	"github.com/unboxed/webcrawler/fetcher"
)

func handleFetchedData(ch chan fetcher.FetchedResult) {
	for result := range ch {
		fmt.Printf("got result: %s\n", result.Url)
	}
}

func TestCrawlAllUrlsOnOneDepth(t *testing.T) {
	fetchingChan := make(chan fetcher.FetchedResult)
	doneChan := make(chan bool)
	go handleFetchedData(fetchingChan)

	crawler := Crawler{
		Fetching:    fetchingChan,
		Done:        doneChan,
		Depth:       5,
		Concurrency: 6,
	}
	go crawler.Crawl("http://golang.org/", fakeFetcherPopulated)

	<-doneChan

	close(doneChan)
	close(fetchingChan)
}

func TestCrawlAllUrlsOnFiveDepth(t *testing.T) {
	fetchingChan := make(chan fetcher.FetchedResult)
	doneChan := make(chan bool)
	go handleFetchedData(fetchingChan)

	crawler := Crawler{
		Fetching:    fetchingChan,
		Done:        doneChan,
		Depth:       5,
		Concurrency: 6,
	}
	go crawler.Crawl("http://golang.org/", fakeFetcherPopulated)

	<-doneChan

	close(doneChan)
	close(fetchingChan)
}

type fakeFetcher map[string]*fetcher.FetchedResult

func (f fakeFetcher) Fetch(url string) (*fetcher.FetchedResult, error) {
	if res, ok := f[url]; ok {
		return res, nil
	}
	return nil, errors.New("URL does not exist")
}

var fakeFetcherPopulated = fakeFetcher{
	"http://golang.org/": &fetcher.FetchedResult{
		"http://golang.org/",
		"The go programming langauge",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
		200,
	},
	"http://golang.org/pkg/": &fetcher.FetchedResult{
		"http://golang.org/pkg/",
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
		200,
	},
	"http://golang.org/pkg/fmt/": &fetcher.FetchedResult{
		"http://golang.org/pkg/fmt/",
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
		200,
	},
	"http://golang.org/pkg/os/": &fetcher.FetchedResult{
		"http://golang.org/pkg/os/",
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
		200,
	},
}
