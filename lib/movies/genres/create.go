package genres

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (genre *Genre) Create() {
	Genre := GetGenre(genre.Title)
	if Genre.Title != "" {
		Genre.Update(genre.Types[0])
		return
	}
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Genres")
	genre.NewID()
	collection.InsertOne(ctx, genre)
}

func (Genre *Genre) Update(Type string) {
	if Genre.HasType(Type) {
		return
	}
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Genres")
	filter := bson.M{"title": bson.M{"$eq": Genre.Title}}
	update := bson.M{"$addToSet": bson.M{"types": Type}}
	collection.UpdateOne(ctx, filter, update)
}

func (genre *Genre) NewID() {
	var Genre Genre
	ID := primitive.NewObjectID()
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Genres")

	err := collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&Genre)
	if err == nil {
		genre.NewID()
	}
	genre.ID = ID
}