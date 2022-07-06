package movies

import (
	"context"
	"errors"
	"fmt"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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


func (movie *Movie) SetSeasons()error {
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Movies")
	var Movie Movie

	opts := options.FindOne().SetProjection(bson.D{{"seasons.episodes", 0}})

	err := collection.FindOne(ctx, bson.M{"_id": movie.ID}, opts).Decode(&Movie)
	if err != nil {
		fmt.Println(err)
		return errors.New(variables.MovieNotFound)
	}
	movie.Seasons = Movie.Seasons
	return nil
}