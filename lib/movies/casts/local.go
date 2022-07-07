package casts

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// add cast to the local database
func (cast *Cast) AddToLocal() {
	if cast.Exists() {
		cast.UpdateLocal()
		return
	}
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Casts")

	collection.InsertOne(ctx, cast)
}


// update cast in the local database
func (cast *Cast) UpdateLocal() {
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Casts")

	filter := bson.M{"_id": bson.M{"$eq": cast.ID}}
	update := bson.M{"$set": cast}

	collection.UpdateOne(ctx, filter, update)
}

// get cast with the name
func LoadCastByName(name string) Cast {
	var Cast Cast
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Casts")

	collection.FindOne(ctx, bson.M{"name": name}).Decode(&Cast)

	return Cast
}


// get cast with the id
func LoadCastByID(ID primitive.ObjectID) Cast {
	var Cast Cast
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Casts")

	collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&Cast)

	return Cast
}