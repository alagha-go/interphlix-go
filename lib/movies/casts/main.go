package casts

import (
	"context"
	"interphlix/lib/variables"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func init() {
	LoadCasts()
	go ListenForCasts()
}


func LoadCasts() {
	var documents []interface{}
	ctx := context.Background()
	collection1 := variables.Client.Database("Interphlix").Collection("Casts")
	collection := variables.Local.Database("Interphlix").Collection("Casts")

	cursor, err := collection1.Find(ctx, bson.M{})
	HandleError(err)
	err = cursor.All(ctx, &documents)
	HandleError(err)
	err = collection.Drop(ctx)
	HandleError(err)
	_, err = collection.InsertMany(ctx, documents)
	HandleError(err)
}

func HandleError(err error) {
	if err != nil && err != mongo.ErrEmptySlice {
		log.Panic(err)
	}
}