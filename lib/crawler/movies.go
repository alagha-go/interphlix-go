package crawler

import (
	"fmt"
	"interphlix/lib/movies"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

var (
	MovieExists int
	TvShowExists int
)

// collect all movies from all pages
func CollectAllMovies(wg *sync.WaitGroup, pagesLenght int) {
	defer wg.Done()
	for index:=1; index<pagesLenght+1; index++ {
		CollectPageMovies(index, "Movie")
		if MovieExists > 10 {
			break
		}
	}
}

// collect all tvshows from all pages
func CollectTvShows(wg *sync.WaitGroup, pagesLenght int) {
	defer wg.Done()
	for index:=1; index<pagesLenght+1; index++ {
		CollectPageMovies(index, "Tv-Show")
		if TvShowExists > 10 {
			break
		}
	}
}

// collect all movies or Tvshows from a page
func CollectPageMovies(page int, Type string) {
	url := fmt.Sprintf("https://tinyzonetv.to/%s?page=%d", strings.ToLower(Type), page)
	collector := colly.NewCollector()

	collector.OnHTML(".film_list-wrap", func(element *colly.HTMLElement) {
		CollectMovies(element, Type)
	})

	collector.Visit(url)
}


// collect all movies or tvshows from the html content
func CollectMovies(element *colly.HTMLElement, Type string) {
	element.ForEach(".flw-item", func(pos int, element *colly.HTMLElement) {
		var Movie movies.Movie
        Movie.Title = element.ChildAttr("a", "title")
        Movie.ImageUrl = element.ChildAttr("img", "data-src")
        Movie.PageUrl = "https://tinyzonetv.to" + element.ChildAttr("a", "href")
		index := strings.Index(Movie.PageUrl, "free-")
    	Movie.Code = Movie.PageUrl[index+5:]
		Movie.Type = Type
		if !Movie.Exists() {
			CollectMovieContent(&Movie)
			if Type == "Movie" {
				url := "https://tinyzonetv.to/ajax/movie/episodes/"+ Movie.Code
				setter := Setter{Url: url}
				setter.SetServers()
				Movie.Server = setter.Server
				Movie.Servers = setter.Servers
				Movie.Available = setter.Available
			}
			Movie.Upload()
		}else {
			if Type == "Movie" {
				MovieExists++
			}else {
				TvShowExists++
			}
		}
	})
}