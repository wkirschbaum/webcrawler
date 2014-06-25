package main

import (
	"log"
	"os"

	"github.com/unboxed/webcrawler/crawler"
	"github.com/unboxed/webcrawler/fetcher"
)

func openLogFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("Unable to open the log file")
	}
	return f
}

func handleFetchedData(ch chan fetcher.FetchedResult) {
	logFileTheRest := openLogFile("rest.log")
	defer logFileTheRest.Close()
	logFile200 := openLogFile("200.log")
	defer logFile200.Close()
	logFile400 := openLogFile("400.log")
	defer logFile400.Close()
	logFile500 := openLogFile("500.log")
	defer logFile500.Close()
	logFileDoubleEscaped := openLogFile("doublescaped.log")
	defer logFileDoubleEscaped.Close()

	for result := range ch {
		switch {
		case result.HasDoubleEscapedHtml:
			log.SetOutput(logFileDoubleEscaped)
		case result.Status > 199 && result.Status < 299:
			log.SetOutput(logFile200)
		case result.Status > 399 && result.Status < 499:
			log.SetOutput(logFile400)
		case result.Status > 499 && result.Status < 599:
			log.SetOutput(logFile500)
		default:
			log.SetOutput(logFileTheRest)
		}
		log.Printf("got result: %s status: %d on level %d\n", result.Url, result.Status, result.Level)
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
		Depth:       10,
		Concurrency: 6,
	}

	go c.Crawl(baseUrl, fetcher.MakeFetcher(baseUrl))

	<-doneChan

	close(doneChan)
	close(fetchingChan)
}
