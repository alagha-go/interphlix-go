package movies

import (
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func GetMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Type := mux.Vars(req)["type"]
	genre := mux.Vars(req)["genre"]
	seed, err := strconv.ParseInt(req.URL.Query().Get("seed"), 10, 64)
	if err != nil {
		seed = 0
	}
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil && seed != 0{
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Response{Action: variables.GetMovies, Failed: true, Error: variables.NoRound}))
		return
	}
	if Type == "all" {
		data, status := movies.GetMoviesByGenre(genre, round, seed)
		res.WriteHeader(status)
		res.Write(data)
		return
	}

	data, status := movies.GetMoviesByGenreAndType(Type, genre, round, seed)
	res.WriteHeader(status)
	res.Write(data)
}