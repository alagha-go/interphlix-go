package casts

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// get cast with the name
func LoadCastByName(name string) Cast {
	var Cast Cast
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Casts")

	collection.FindOne(ctx, bson.M{"name": name}).Decode(&Cast)

	return Cast
}


// get cast with the id
func LoadCastByID(ID primitive.ObjectID) Cast {
	var Cast Cast
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Casts")

	collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&Cast)

	return Cast
}