package movies

import (
	"fmt"
	"interphlix/lib/movies/genres"
	"interphlix/lib/variables"
	"math/rand"
	"net/http"
	"time"
)

var (
	MoviesLimit = 20
)


func GetRecommendationMovies() ([]byte, int) {
	Response := variables.Response{Action: variables.GetRecommend}
	var recommendation Recommendation
	recommendation.Seed = time.Now().UnixNano()

	Genres := genres.GetGenres()

	TrendingMovies := LoadTrendingMovies(0, MoviesLimit, recommendation.Seed)
	FeaturedMovies := LoadFeaturedMovies(0, MoviesLimit, recommendation.Seed)
	PopularMovies := LoadPoPularMovies(0, MoviesLimit, recommendation.Seed)
	PopularTvShows := LoadPoPularTvShows(0, MoviesLimit, recommendation.Seed)

	Categories := []Category{
		{Title: "Trending", Path: "/movies/trending", Movies: TrendingMovies},
		{Title: "Featured", Path: "/movies/featured", Movies: FeaturedMovies},
		{Title: "Popular Movies", Path: "/movies/popular", Movies: PopularMovies},
		{Title: "Popular Tvs", Path: "/tv-shows/popular", Movies: PopularTvShows},
	}

	recommendation.Categories = append(recommendation.Categories, Categories...)

	for index := range Genres {
		category := Category{Title: Genres[index].Title, Path: fmt.Sprintf("/movies/%s", Genres[index].Title), Movies: RandomMovies(recommendation.Seed, LoadMoviesByGenre(Genres[index].Title))}
		if len(category.Movies) > MoviesLimit {
			category.Movies = category.Movies[:MoviesLimit]
		}
		recommendation.Categories = append(recommendation.Categories, category)
	}

	Response.Success = true
	Response.Data = recommendation
	return variables.JsonMarshal(Response), http.StatusOK
}

// randomly shuffle movies and return
func RandomMovies(seed int64, Movies []Movie) []Movie {
	rand.Seed(seed)
	rand.Shuffle(len(Movies), func(i, j int) { Movies[i], Movies[j] = Movies[j], Movies[i] })
	return Movies
}