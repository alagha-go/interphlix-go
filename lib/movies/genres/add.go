package genres

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)

// add genre to the local databse
func (Genre *Genre) AddToLocal() {
	if Genre.Exists() {
		Genre.UpdateLocal()
		return
	}
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Genres")

	collection.InsertOne(ctx, Genre)
}

// update genre to the local database
func (Genre *Genre) UpdateLocal() {
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Genres")
	filter := bson.M{"_id": bson.M{"$eq": Genre.ID}}
	update := bson.M{"$set": Genre}
	collection.UpdateOne(ctx, filter, update)
}

// check if a genre exists
func (genre *Genre) Exists() bool {
	Genre := GetGenre(genre.Title)
	return Genre.Title != ""
}