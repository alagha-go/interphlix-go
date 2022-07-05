package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/// upload movie to the database
func (movie *Movie) Upload() ([]byte, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	Response := variables.Response{Action: variables.UploadMovie}
	if movie.Exists() {
		Response.Failed = true
		Response.Error = variables.MovieExists
		return variables.JsonMarshal(Response), http.StatusConflict
	}
	if strings.Contains(movie.ID.Hex(), "000000") {
		movie.NewID()
	}
	for sindex := range movie.Seasons {
		for eindex := range movie.Seasons[sindex].Episodes {
			if strings.Contains(movie.Seasons[sindex].Episodes[eindex].ID.Hex(), "000000") {
				movie.Seasons[sindex].Episodes[eindex].ID = primitive.NewObjectID()
			}
		}
		if strings.Contains(movie.Seasons[sindex].ID.Hex(), "000000") {
			movie.Seasons[sindex].ID = primitive.NewObjectID()
		}
	}
	_, err := collection.InsertOne(ctx, movie)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	Response.Success = true
	Response.Data = movie
	return variables.JsonMarshal(Response), http.StatusCreated
}


/// addseason to a movie
func (movie *Movie) AddSeason(season *Season) ([]byte, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	Response := variables.Response{Action: variables.AddSeason}

	if strings.Contains(season.ID.Hex(), "000000") {
		season.ID = primitive.NewObjectID()
	}

	for index := range season.Episodes {
		if strings.Contains(season.Episodes[index].ID.Hex(), "000000") {
			season.Episodes[index].ID = primitive.NewObjectID()
		}
	}

	filter := bson.M{"_id": bson.M{"$eq": movie.ID}}
	update := bson.M{"$addToSet": bson.M{"seasons": season}}

	cursor, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	if cursor.ModifiedCount == 0 && cursor.MatchedCount == 0{
		Response.Failed = true
		Response.Error = variables.MovieNotFound
		return variables.JsonMarshal(Response), http.StatusNotFound
	}
	Response.Success = true
	Response.Data = season
	return variables.JsonMarshal(Response), http.StatusCreated
}


/// add episode to a season
func(season *Season) AddEpisode(episode *Episode, ID *primitive.ObjectID) ([]byte, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	Response := variables.Response{Action: variables.AddEpisode}

	if strings.Contains(episode.ID.Hex(), "000000") {
		episode.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": bson.M{"$eq": ID}, "seasons._id": bson.M{"$eq": season.ID}}
	update := bson.M{"$addToSet": bson.M{"seasons.$.episodes": episode}}

	cursor, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	if cursor.ModifiedCount == 0 && cursor.MatchedCount == 0{
		Response.Failed = true
		Response.Error = variables.SeasonNotFound
		return variables.JsonMarshal(Response), http.StatusNotFound
	}
	Response.Success = true
	Response.Data = episode
	return variables.JsonMarshal(Response), http.StatusCreated
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