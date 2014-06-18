package main

import (
	"fmt"

	"github.com/wkirschbaum/webcrawler/crawler"
	"github.com/wkirschbaum/webcrawler/fetcher"
)

func main() {

	result := make(chan fetcher.FetchedResult)

	fchr := fetcher.MakeFetcher("http://app01-uat-rails3.c45383.blueboxgrid.com")
	c := &crawler.Crawler{Found: make(map[string]int), Fetched: result}
	go c.Crawl("http://app01-uat-rails3.c45383.blueboxgrid.com/", 5, fchr)

	for {
		fetched := <-result
		fmt.Printf("found %s : %d", fetched.Url, fetched.Status)
	}
}
