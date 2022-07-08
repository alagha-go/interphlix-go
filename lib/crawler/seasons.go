package crawler

import (
	"interphlix/lib/movies"

	"github.com/gocolly/colly"
)


func GetSeasons(Movie *movies.Movie) {
	collector := colly.NewCollector()
	url := "https://tinyzonetv.to/ajax/v2/tv/seasons/" + Movie.Code

	collector.OnHTML(".dropdown-menu.dropdown-menu-new", func(element *colly.HTMLElement) {
		CollectAllSeasons(element, Movie)
	})

	collector.Visit(url)
}


func CollectAllSeasons(element *colly.HTMLElement, Movie *movies.Movie) {
	element.ForEach("a", func(index int, element *colly.HTMLElement) {
		var Season movies.Season
		Season.Index = index
		Season.Code = element.Attr("data-id")
		Season.Name = element.Text
		Movie.Seasons = append(Movie.Seasons, Season)
	})
}


func CheckForNewSeasons(Movie *movies.Movie) {
	Movie.SetSeasons()
	movie := movies.Movie{Seasons: Movie.Seasons}
	Movie.Seasons = []movies.Season{}
	GetSeasons(Movie)

	for index := range Movie.Seasons {
		if !SeasonExistInSeasons(&Movie.Seasons[index], &movie) {
			Movie.AddSeason(&Movie.Seasons[index])
		}
	}
}

func SeasonExistInSeasons(Season *movies.Season, Movie *movies.Movie) bool {
	for index := range Movie.Seasons {
		if Movie.Seasons[index].Code == Season.Code {
			return true
		}
	}
	return false
}