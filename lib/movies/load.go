package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// load featured movies
func LoadFeaturedMovies() []Movie {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	opts := options.Find().SetProjection(bson.D{{"_id", 1}, {"image_url", 1}, {"title", 1}, {"type", 1},})

	cursor, err := collection.Find(ctx, bson.M{"featured": true}, opts)
	if err != nil {
		variables.SaveError(err, "movies", "LoadFeaturedMovies")
		return Movies
	}
	cursor.All(ctx, &Movies)

	return Movies
}

// loading trending Movies
func LoadTrendingMovies() []Movie {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	opts := options.Find().SetProjection(bson.D{{"_id", 1}, {"image_url", 1}, {"title", 1}, {"type", 1},})

	cursor, err := collection.Find(ctx, bson.M{"trending": true}, opts)
	if err != nil {
		variables.SaveError(err, "movies", "LoadTrendingMovies")
		return Movies
	}
	cursor.All(ctx, &Movies)

	return Movies
}


// get popular movies
func LoadPoPularMovies() []Movie {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	opts := options.Find().SetProjection(bson.D{{"_id", 1}, {"image_url", 1}, {"title", 1}, {"type", 1},})

	cursor, err := collection.Find(ctx, bson.M{"type": "Movie", "popular": true}, opts)
	if err != nil {
		variables.SaveError(err, "movies", "LoadPopularMovies")
		return Movies
	}
	cursor.All(ctx, &Movies)

	return Movies
}

// load popular tvshows
func LoadPoPularTvShows() []Movie {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	opts := options.Find().SetProjection(bson.D{{"_id", 1}, {"image_url", 1}, {"title", 1}, {"type", 1},})

	cursor, err := collection.Find(ctx, bson.M{"type": "Tv-Show", "popular": true}, opts)
	if err != nil {
		variables.SaveError(err, "movies", "LoadPopularTvShows")
		return Movies
	}
	cursor.All(ctx, &Movies)

	return Movies
}

// load movies by genre
func LoadMoviesByGenre(genre string) []Movie {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	opts := options.Find().SetProjection(bson.D{{"_id", 1}, {"image_url", 1}, {"title", 1}, {"type", 1},})

	cursor, err := collection.Find(ctx, bson.M{"type": "Tv-Show", "genre": genre}, opts)
	if err != nil {
		variables.SaveError(err, "movies", "LoadMoviesByGenre")
		return Movies
	}
	cursor.All(ctx, &Movies)

	return Movies
}