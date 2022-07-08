package crawler

import (
	"interphlix/lib/movies"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetEpisodes(Season *movies.Season) {
	collector := colly.NewCollector()
	url := "https://tinyzonetv.to/ajax/v2/season/episodes/" + Season.Code

	collector.OnHTML(".nav", func(element *colly.HTMLElement) {
		
	})

	collector.Visit(url)
}

func CollectAllEpisodes(element *colly.HTMLElement, Season *movies.Season) {
	element.ForEach(".nav-item", func(_ int, element *colly.HTMLElement) {
		var Episode movies.Episode
		Episode.ID = primitive.NewObjectID()
		Episode.Name = element.ChildText("a")
		index := strings.Index(Episode.Name, "Eps")
		end := strings.Index(Episode.Name, "\n")
		Episode.Index, _ = strconv.Atoi(Episode.Name[index+4:end])
		index = strings.Index(Episode.Name, "\n")
		Episode.Name = Episode.Name[index+1:]
		Episode.Name = strings.ReplaceAll(Episode.Name, "                    : ", "")
		Episode.Code = element.ChildAttr("a", "data-id")
		Season.Episodes = append(Season.Episodes, Episode)
	})
}


func CheckForNewEpisodes(Season *movies.Season, ID *primitive.ObjectID) {
	Season.SetEpisodes()
	season := movies.Season{Episodes: Season.Episodes}
	GetEpisodes(Season)

	for index := range Season.Episodes {
		if !EpisodeInEpisodes(&Season.Episodes[index], &season) {
			Season.AddEpisode(&Season.Episodes[index], ID)
		}
	}
}

func EpisodeInEpisodes(Episode *movies.Episode, Season *movies.Season) bool {
	for index := range Season.Episodes {
		if Season.Episodes[index].Code == Episode.Code {
			return true
		}
	}
	return false
}