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

func (movie *Movie) SetSeasons() error {
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Movies")
	var Movie Movie

	opts := options.FindOne().SetProjection(bson.M{"seasons.episodes": 0, "seasons.code": 0})

	err := collection.FindOne(ctx, bson.M{"_id": movie.ID}, opts).Decode(&Movie)
	if err != nil {
		fmt.Println(err)
		return errors.New(variables.MovieNotFound)
	}
	movie.Seasons = Movie.Seasons
	return nil
}

func (season *Season) SetEpisodes() error {
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Movies")
	var movie Movie

	opts := options.FindOne().SetProjection(bson.M{"seasons": bson.M{"$elemMatch": bson.M{"_id": season.ID}}})

	err := collection.FindOne(ctx, bson.M{"seasons._id": season.ID}, opts).Decode(&movie)
	if err != nil {
		println(err.Error())
		return errors.New(variables.MovieNotFound)
	}
	if len(movie.Seasons) > 0 {
		for index := range movie.Seasons[0].Episodes {
			movie.Seasons[0].Episodes[index].Servers = nil
			movie.Seasons[0].Episodes[index].Server = nil
			movie.Seasons[0].Episodes[index].Code = ""
		}
		season.Episodes = movie.Seasons[0].Episodes
	}
	return nil
}