package movies

import (
	"encoding/json"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UploadMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.UploadMovie}
	var movie movies.Movie
	err := json.NewDecoder(req.Body).Decode(&movie)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidJson
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	data, status := movie.Upload()
	res.WriteHeader(status)
	res.Write(data)
}

func UploadSeason(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.AddSeason}
	MovieID, err := primitive.ObjectIDFromHex(mux.Vars(req)["movieId"])
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidID
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	var season movies.Season
	err = json.NewDecoder(req.Body).Decode(&season)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidJson
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	movie := movies.Movie{ID: MovieID}
	data, status := movie.AddSeason(&season)
	res.WriteHeader(status)
	res.Write(data)
}


func UploadEpisode(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.AddEpisode}
	MovieID, err := primitive.ObjectIDFromHex(mux.Vars(req)["movieId"])
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidID
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	SeasonID, err := primitive.ObjectIDFromHex(mux.Vars(req)["seasonId"])
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidID
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	var episode movies.Episode
	err = json.NewDecoder(req.Body).Decode(&episode)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidJson
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	season := movies.Season{ID: SeasonID}
	data, status := season.AddEpisode(&episode, &MovieID)
	res.WriteHeader(status)
	res.Write(data)
}