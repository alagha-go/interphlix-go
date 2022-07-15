package casts

import (
	"context"
	"interphlix/lib/variables"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IndexCasts() {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Casts")

	model := mongo.IndexModel{
		Keys: bson.D{{"name", "text"}, {"also_known_as", "text"}},
		Options: options.Index().SetName("Casts"),
	}

	_, err := collection.Indexes().CreateOne(ctx, model)
	if err != nil {
		log.Panic(err)
	}
}