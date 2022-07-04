package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func (movie *Movie) Upload() ([]byte, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	Response := variables.Response{Action: variables.UploadMovie}
	if movie.Exists() {
		Response.Failed = true
		Response.Error = variables.MovieExists
		return variables.JsonMarshal(Response), http.StatusConflict
	}
	movie.NewID()
	_, err := collection.InsertOne(ctx, movie)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	Response.Success = true
	Response.Data = movie
	return variables.JsonMarshal(Response), http.StatusOK
}



/// check if the movie exists in our database
func (movie *Movie) Exists() bool {
	var dbMovie Movie
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	err := collection.FindOne(ctx, bson.M{"code": movie.Code}).Decode(&dbMovie)
	return err == nil
}

/// generate new MovieID 
func (movie *Movie) NewID() {
	var dbMovie Movie
	ID := primitive.NewObjectID()
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	err := collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&dbMovie)
	if err == nil {
		movie.NewID()
	}
	movie.ID = ID
}