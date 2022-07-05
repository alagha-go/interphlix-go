package movies

import (
	"interphlix/lib/movies"
	"net/http"
)

func GetRecommendationMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	data, status := movies.GetRecommendationMovies()
	res.WriteHeader(status)
	res.Write(data)
}