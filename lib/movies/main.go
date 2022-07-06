package movies

import (
	"context"
	"interphlix/lib/variables"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	LoadMovies()
	go ListenForMoviesCollection()
	IndexMovies()
}


func LoadMovies() {
	var documents []interface{}
	ctx := context.Background()
	collection1 := variables.Client.Database("Interphlix").Collection("Movies")
	collection := variables.Local.Database("Interphlix").Collection("Movies")

	cursor, err := collection1.Find(ctx, bson.M{})
	HandleError(err)
	err = cursor.All(ctx, &documents)
	HandleError(err)
	collection.Drop(ctx)
	_, err = collection.InsertMany(ctx, documents)
	HandleError(err)
}

func HandleError(err error) {
	if err != nil && err != mongo.ErrEmptySlice{
		log.Panic(err)
	}
}