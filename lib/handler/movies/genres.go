package movies

import (
	"interphlix/lib/movies/genres"
	"interphlix/lib/variables"
	"net/http"
)


func GetGenres(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.GetGenres, Success:  true}
	Type := req.URL.Query().Get("type")
	if Type != "" && Type != "all" {
		Response.Data = genres.GetGenresByType(Type)
	}else {
		Response.Data = genres.GetGenresByType()
	}
	res.WriteHeader(http.StatusOK)
	res.Write(variables.JsonMarshal(Response))
}