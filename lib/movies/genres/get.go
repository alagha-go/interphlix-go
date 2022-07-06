package genres

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetGenresByType(Type ...string) []Genre {
	filter := bson.M{}
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Genres")
	var Genres []Genre
	opts := options.Find().SetProjection(bson.M{"types": 0})
	if len(Type) > 0 {
		filter = bson.M{"types": Type[0]}
	}
	cursor, err := collection.Find(ctx, filter, opts)
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