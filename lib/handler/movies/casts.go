package movies

import (
	"interphlix/lib/movies/casts"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetCasts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 0
	}
	data, status := casts.GetCasts(round)
	res.WriteHeader(status)
	res.Write(data)
}


func GetCast(res http.ResponseWriter, req *http.Request) {
	var ID *primitive.ObjectID
	res.Header().Set("content-type", "application/json")
	name := mux.Vars(req)["id"]
	Id, err := primitive.ObjectIDFromHex(name)
	if err != nil {
		ID = nil
	}else {
		ID = &Id
	}

	data, status := casts.GetCast(name, ID)
	res.WriteHeader(status)
	res.Write(data)
}