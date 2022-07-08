package crawler

import (
	"interphlix/lib/movies"
	"strings"

	"github.com/gocolly/colly"
)

// collect movie content
func CollectMovieContent(Movie *movies.Movie) {
    collector := colly.NewCollector()

    collector.OnHTML(".description", func(element *colly.HTMLElement){
        Movie.Description = element.Text
        Movie.Description = strings.ReplaceAll(Movie.Description, "\n", "")
        Movie.Description = strings.ReplaceAll(Movie.Description, "  ", "")
        Movie.Description = strings.TrimPrefix(Movie.Description, " ")
        Movie.Description = strings.TrimSuffix(Movie.Description, " ")
    })

    
    collector.OnHTML(".elements", func(element *colly.HTMLElement) {
        SetElements(element, Movie)
    })
    
    collector.Visit(Movie.PageUrl)
}

// collect movie content from the given html
func SetElements(element *colly.HTMLElement, Movie *movies.Movie) {
    functions := []func(element *colly.HTMLElement, Movie *movies.Movie){}
    functions = append(functions,  SetReleased)
    functions = append(functions,  SetGenre)
    functions = append(functions,  SetCasts)
    functions = append(functions,  SetDuration)
    functions = append(functions,  SetCountries)
    functions = append(functions,  SetProducers)
    element.ForEach(".row-line", func(index int, element *colly.HTMLElement){
        functions[index](element, Movie)
    })
}
