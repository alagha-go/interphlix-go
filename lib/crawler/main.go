package crawler

import (
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

func init() {
	go StartCrawler()
}


func StartCrawler() {
	for {
		var wg sync.WaitGroup
		wg.Add(1)
		go CollectAllMovies(&wg, GetNumberOfPages("https://tinyzonetv.to/movie"))
		wg.Add(1)
		go CollectTvShows(&wg, GetNumberOfPages("https://tinyzonetv.to/tv-show"))
		wg.Wait()
		time.Sleep(48*time.Hour)
	}
}


func GetNumberOfPages(url string) int {
	var err error
	var numberofPages int
	collector := colly.NewCollector()

	collector.OnHTML(".pagination.pagination-lg.justify-content-center", func(element *colly.HTMLElement) {
		element.ForEach(".page-item", func(_ int, element *colly.HTMLElement) {
			title := element.ChildAttr("a", "title")
			href := element.ChildAttr("a", "href")
			if title == "Last" {
				href = strings.ReplaceAll(href, "/movie?page=", "")
				href = strings.ReplaceAll(href, "/tv-show?page=", "")
				numberofPages, err = strconv.Atoi(href)
				HandleError(err)
			}
		})
	})

	collector.Visit(url)

	return numberofPages
}


func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}