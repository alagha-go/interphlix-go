package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)

// add to movie to the local database
func (movie *Movie) AddToLocal() {
	if movie.Exists() {
		movie.UpdateLocal()
		return
	}
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Movies")

	collection.InsertOne(ctx, movie)
}


// update movie to the local database
func (movie *Movie) UpdateLocal() {
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Movies")

	filter := bson.M{"_id": bson.M{"$eq": movie.ID}}
	update := bson.M{"$set": movie}
	collection.UpdateOne(ctx, filter, update)
}