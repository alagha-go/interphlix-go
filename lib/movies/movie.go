package movies

import (
	"context"
	"errors"
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

	opts := options.FindOne().SetProjection(bson.M{"seasons.episodes": 0})

	err := collection.FindOne(ctx, bson.M{"_id": movie.ID}, opts).Decode(&Movie)
	if err != nil {
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
		return errors.New(variables.MovieNotFound)
	}
	if len(movie.Seasons) > 0 {
		for index := range movie.Seasons[0].Episodes {
			movie.Seasons[0].Episodes[index].Servers = nil
			movie.Seasons[0].Episodes[index].Server = nil
		}
		season.Code = movie.Seasons[0].Code
		season.Episodes = movie.Seasons[0].Episodes
	}
	return nil
}


func (Movie *Movie) GetCode() bool {
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Movies")

	opts := options.FindOne().SetProjection(bson.M{"_id": 1, "code": 1, "title": 1, "type": 1,})

	err := collection.FindOne(ctx, bson.M{"_id": Movie.ID}, opts).Decode(&Movie)
	return err == nil
}