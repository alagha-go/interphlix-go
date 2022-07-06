package movies

import (
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetMovie(ID primitive.ObjectID) ([]byte, int) {
	Response := variables.Response{Action: variables.GetMovie}
	movie, err := LoadMovie(ID)
	if err != nil {
		Response.Failed = true
		Response.Error = err.Error()
		return variables.JsonMarshal(Response), http.StatusNotFound
	}
	Response.Success = true
	Response.Data = movie
	return variables.JsonMarshal(Response), http.StatusOK
}