package genres

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)

func GetGenres() []Genre {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Genres")
	var Genres []Genre
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		variables.SaveError(err, "genres", "GetGenres")
		return Genres
	}
	cursor.All(ctx, &Genres)
	return Genres
}