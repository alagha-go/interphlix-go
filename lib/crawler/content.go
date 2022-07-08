package crawler

import (
	"interphlix/lib/movies"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)


// collect the movie release date from the html
func SetReleased(element *colly.HTMLElement, Movie *movies.Movie) {
    layout := "2006-01-02"
    released := element.Text
    released = strings.ReplaceAll(released, "Released: ", "")
    released = strings.ReplaceAll(released, "  ", "")
    released = strings.ReplaceAll(released, "\n", "")
    released = strings.TrimPrefix(released, " ")
    released = strings.TrimSuffix(released, " ")
    Released, _ := time.Parse(layout, released)
    Movie.Released = &Released
}

// collect the movie genres from the html
func SetGenre(element *colly.HTMLElement, Movie *movies.Movie) {
    element.ForEach("a", func(index int, element *colly.HTMLElement){
        Movie.Genres = append(Movie.Genres, element.Text)
    })
}

// collect the movie casts from the html
func SetCasts(element *colly.HTMLElement, Movie *movies.Movie) {
    element.ForEach("a", func(index int, element *colly.HTMLElement){
        Movie.Casts = append(Movie.Casts, element.Text)
    })
}

// collect the movie duration from the html
func SetDuration(element *colly.HTMLElement, Movie *movies.Movie) {
    duration := element.Text
    duration = strings.ReplaceAll(duration, "Duration: ", "")
    duration = strings.ReplaceAll(duration, "  ", "")
    duration = strings.ReplaceAll(duration, "\n", "")
    duration = strings.ReplaceAll(duration, "min", "")
    duration = strings.TrimPrefix(duration, " ")
    duration = strings.TrimSuffix(duration, " ")
    if strings.Contains(duration, "N/A") {
        Movie.Duration = 0
    }else {
        minutes, _ := strconv.Atoi(duration)
        Movie.Duration = time.Duration(minutes*int(time.Minute))
    }
}

// collect the movie countries from the html
func SetCountries(element *colly.HTMLElement, Movie *movies.Movie) {
    element.ForEach("a", func(index int, element *colly.HTMLElement){
        Movie.Countries = append(Movie.Countries, element.Text)
    })
}


// collect the movie production companies from the html
func SetProducers(element *colly.HTMLElement, Movie *movies.Movie) {
    production := element.Text
    production = strings.ReplaceAll(production, "Production: ", "")
    production = strings.ReplaceAll(production, "  ", "")
    production = strings.ReplaceAll(production, "\n", "")
    production = strings.TrimPrefix(production, " ")
    production = strings.TrimSuffix(production, " ")
    if production == "N/A" {
        return
    }
    Movie.Producers = strings.Split(production, ",")
}