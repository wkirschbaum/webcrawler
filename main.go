package main

import (
	"fmt"

	"github.com/unboxed/webcrawler/crawler"
	"github.com/unboxed/webcrawler/fetcher"
)

func handleFetchedData(ch chan fetcher.FetchedResult) {
	for result := range ch {
		fmt.Printf("got result: %s status: %d\n", result.Url, result.Status)
	}
}

func main() {
	baseUrl := "http://localhost:3000"
	fetchingChan := make(chan fetcher.FetchedResult)
	doneChan := make(chan bool)
	go handleFetchedData(fetchingChan)

	c := crawler.Crawler{
		Fetching:    fetchingChan,
		Done:        doneChan,
		Depth:       5,
		Concurrency: 6,
	}

	go c.Crawl(baseUrl, fetcher.MakeFetcher(baseUrl))

	<-doneChan

	close(doneChan)
	close(fetchingChan)
}
