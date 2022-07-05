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
	if IndexExists("Movies") {
		return
	}
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Movies")

	model := mongo.IndexModel{
		Keys: bson.D{{"title", "text"}, {"casts", "text"}, {"genre", "text"}},
		Options: options.Index().SetName("Movies"),
	}

	_, err := collection.Indexes().CreateOne(ctx, model)
	if err != nil {
		log.Panic(err)
	}
}

func IndexExists(name string) bool {
	var Indexes []mongo.IndexModel
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Movies")

	cursor, _ := collection.Indexes().List(ctx)
	cursor.All(ctx, &Indexes)

	for index := range Indexes {
		if *Indexes[index].Options.Name == name {
			return true
		}
	}
	return false
}