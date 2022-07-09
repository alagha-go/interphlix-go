package movies

import (
	"interphlix/lib/crawler"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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


func GetMoviesByCast(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	id := mux.Vars(req)["id"]
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 1
	}
	data, status := movies.GetMoviesByCast(id, round)
	res.WriteHeader(status)
	res.Write(data)
}


func CheckForNewEpisodes(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.CheckMovie}
	ID, err := primitive.ObjectIDFromHex(mux.Vars(req)["id"])
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidID
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	movie := movies.Movie{ID: ID}
	exists := movie.GetCode()
	if !exists {
		Response.Failed = true
		Response.Error = variables.MovieNotFound
		res.WriteHeader(http.StatusNotFound)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	crawler.CheckForNewSeasons(&movie)
	Response.Success = true
	Response.Data = "done"
	res.WriteHeader(http.StatusOK)
	res.Write(variables.JsonMarshal(Response))
}