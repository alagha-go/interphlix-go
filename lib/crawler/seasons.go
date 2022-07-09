package crawler

import (
	"interphlix/lib/movies"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		Season.ID = primitive.NewObjectID()
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
		exists, i := SeasonExistInSeasons(&Movie.Seasons[index], &movie)
		if !exists {
			Movie.AddSeason(&Movie.Seasons[index])
			CheckForNewEpisodes(&Movie.Seasons[index], &Movie.ID)
			continue
		}
		CheckForNewEpisodes(&movie.Seasons[i], &Movie.ID)
	}
}

func SeasonExistInSeasons(Season *movies.Season, Movie *movies.Movie) (bool, int) {
	for index := range Movie.Seasons {
		if Movie.Seasons[index].Code == Season.Code {
			return true, index
		}
	}
	return false, 0
}