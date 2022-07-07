package casts

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (cast *Cast) Create() {
	if cast.Exists() {
		cast.Update()
		return
	}
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Casts")
	cast.NewID()

	collection.InsertOne(ctx, cast)
}


func (cast *Cast) Update() {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Casts")

	filter := bson.M{"_id": bson.M{"$eq": cast.ID}}
	update := bson.M{"$set": cast}
	collection.UpdateOne(ctx, filter, update)
}



func (cast *Cast) Exists() bool {
	var Cast Cast
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Casts")

	err := collection.FindOne(ctx, bson.M{"name": cast.Name}).Decode(&Cast)
	return err == nil
}

func (cast *Cast) NewID() {
	var Cast Cast
	ID := primitive.NewObjectID()
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Casts")

	err := collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&Cast)
	if err == nil {
		cast.NewID()
	}
	cast.ID = &ID
}