package movies

import (
	"interphlix/lib/variables"
	"net/http"
)

// get movies by genre
func GetMoviesByGenre(genre string, round int, seed int64) ([]byte, int) {
	var Movies []Movie
	start := 0
	if round > 0 {
		start = (round*MoviesLimit) - MoviesLimit
	}
	end := start+MoviesLimit
	Response := variables.Response{Action: variables.GetMovies, Success: true, Data: []Movie{}}

	switch genre {
	case "trending":
		Movies = LoadTrendingMovies(start, end, seed)
	case "featured":
		Movies = LoadFeaturedMovies(start, end, seed)
	case "popular":
		Movies = LoadPopulareContent(start, end, seed)
	default:
		Movies = LoadMoviesByGenre(genre)
		if seed != 0 {
			Movies = RandomMovies(seed, Movies)
		}
		if start > len(Movies) {
			Movies = []Movie{}
		}else if end > len(Movies) {
			Movies = Movies[start:]
		}else {
			Movies = Movies[start:end]
		}
	}



	Response.Data = Movies

	return variables.JsonMarshal(Response), http.StatusOK
}

// get movies by genre and type
func GetMoviesByGenreAndType(Type, genre string, round int, seed int64) ([]byte, int) {
	var Movies []Movie
	start := 0
	if round > 0 {
		start = (round*MoviesLimit) - MoviesLimit
	}
	end := start+MoviesLimit
	Response := variables.Response{Action: variables.GetMovies, Success: true, Data: []Movie{}}
	switch genre {
	case "trending":
		Movies = LoadTrendingMovies(start, end, seed, Type)
	case "featured":
		Movies = LoadFeaturedMovies(start, end, seed, Type)
	case "popular":
		Movies = LoadPopulareContent(start, end, seed, Type)
	default:
		Movies = LoadMoviesByGenreAndType(Type, genre)
		if seed != 0 {
			Movies = RandomMovies(seed, Movies)
		}
		if start > len(Movies) {
			Movies = []Movie{}
		}else if end > len(Movies) {
			Movies = Movies[start:]
		}else {
			Movies = Movies[start:end]
		}
	}

	Response.Data = Movies

	return variables.JsonMarshal(Response), http.StatusOK
}