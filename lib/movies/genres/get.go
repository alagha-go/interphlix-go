package genres

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)

func GetGenres() []Genre {
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Genres")
	var Genres []Genre
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		variables.SaveError(err, "genres", "GetGenres")
		return Genres
	}
	cursor.All(ctx, &Genres)
	return Genres
}

func GetGenre(title string) *Genre {
	var Genre *Genre
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Genres")
	collection.FindOne(ctx, bson.M{"title": title}).Decode(&Genre)
	return Genre
}


func (Genre *Genre) HasType(Type string) bool {
	for index := range Genre.Types {
		if Genre.Types[index] == Type {
			return true
		}
	}
	return false
}