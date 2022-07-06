package movies

import (
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	ID, err := primitive.ObjectIDFromHex(mux.Vars(req)["id"])
	if err != nil {
		Response := variables.Response{Action: variables.GetMovie, Failed: true, Error: variables.InvalidID}
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	data, status := movies.GetMovie(ID)
	res.WriteHeader(status)
	res.Write(data)
}

func GetSeasons(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.GetSeasons}
	ID, err := primitive.ObjectIDFromHex(mux.Vars(req)["id"])
	if err != nil {
		Response := variables.Response{Action: variables.GetMovie, Failed: true, Error: variables.InvalidID}
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	movie := movies.Movie{ID: ID}
	err = movie.SetSeasons()
	if err != nil {
		Response.Failed = true
		Response.Error = err.Error()
		res.WriteHeader(http.StatusNotFound)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	Response.Success = true
	Response.Data = movie.Seasons
	res.WriteHeader(http.StatusOK)
	res.Write(variables.JsonMarshal(Response))
}