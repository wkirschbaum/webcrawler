package main

import (
	"fmt"

	"github.com/wkirschbaum/webcrawler/crawler"
	"github.com/wkirschbaum/webcrawler/fetcher"
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
	go crawler.Crawl(baseUrl, 5, fetcher.MakeFetcher(baseUrl), fetchingChan, doneChan)

	<-doneChan

	close(doneChan)
	close(fetchingChan)
}
