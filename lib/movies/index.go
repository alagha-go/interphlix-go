package movies

import (
	"context"
	"interphlix/lib/variables"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IndexMovies() {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	model := mongo.IndexModel{
		Keys: bson.D{{"title", "text"}, {"casts", "text"}, {"genre", "text"}},
		Options: options.Index().SetName("Movies"),
	}

	_, err := collection.Indexes().CreateOne(ctx, model)
	if err != nil {
		log.Panic(err)
	}
}