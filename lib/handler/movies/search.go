package movies

import (
	"interphlix/lib/movies"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func Search(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Type := mux.Vars(req)["type"]
	query := req.URL.Query().Get("query")
	round, _ := strconv.Atoi(req.URL.Query().Get("round"))
	if round < 0 {
		round = 0
	}
	if Type == "all" {
		data, status := movies.Search(round, query)
		res.WriteHeader(status)
		res.Write(data)
		return
	}

	data, status := movies.Search(round, query, Type)
	res.WriteHeader(status)
	res.Write(data)
}