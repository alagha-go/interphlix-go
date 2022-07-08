package movies

import (
	"context"
	"interphlix/lib/movies/casts"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// get movies by genre
func GetMoviesByGenre(genre string, round int, seed int64) ([]byte, int) {
	var Movies []Movie
	start := 0
	if round > 0 {
		start = (round*MoviesLimit) - MoviesLimit
	}
	end := start+MoviesLimit
	Response := variables.Response{Action: variables.GetMovies, Success: true, Data: []Movie{}}

	switch genre {
	case "trending":
		Movies = LoadTrendingMovies(start, end, seed)
	case "featured":
		Movies = LoadFeaturedMovies(start, end, seed)
	case "popular":
		Movies = LoadPopulareContent(start, end, seed)
	default:
		Movies = LoadMoviesByGenre(genre)
		if seed != 0 {
			Movies = RandomMovies(seed, Movies)
		}
		if start > len(Movies) {
			Movies = []Movie{}
		}else if end > len(Movies) {
			Movies = Movies[start:]
		}else {
			Movies = Movies[start:end]
		}
	}



	Response.Data = Movies

	return variables.JsonMarshal(Response), http.StatusOK
}

// get movies by genre and type
func GetMoviesByGenreAndType(Type, genre string, round int, seed int64) ([]byte, int) {
	var Movies []Movie
	start := 0
	if round > 0 {
		start = (round*MoviesLimit) - MoviesLimit
	}
	end := start+MoviesLimit
	Response := variables.Response{Action: variables.GetMovies, Success: true, Data: []Movie{}}
	switch genre {
	case "trending":
		Movies = LoadTrendingMovies(start, end, seed, Type)
	case "featured":
		Movies = LoadFeaturedMovies(start, end, seed, Type)
	case "popular":
		Movies = LoadPopulareContent(start, end, seed, Type)
	default:
		Movies = LoadMoviesByGenreAndType(Type, genre)
		if seed != 0 {
			Movies = RandomMovies(seed, Movies)
		}
		if start > len(Movies) {
			Movies = []Movie{}
		}else if end > len(Movies) {
			Movies = Movies[start:]
		}else {
			Movies = Movies[start:end]
		}
	}

	Response.Data = Movies

	return variables.JsonMarshal(Response), http.StatusOK
}


// get Movies By Cast
func GetMoviesByCast(id string, round int) ([]byte, int) {
	var Movies []Movie
	var ID *primitive.ObjectID
	var name string
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ID = nil
	}else {
		ID = &Id
	}
	if ID != nil {
		Cast := casts.LoadCastByID(*ID)
		name = Cast.Name
	}else {
		name = id
	}
	start := 0
	if round > 0 {
		start = (round*MoviesLimit) - MoviesLimit
	}
	end := start+MoviesLimit

	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Movies")
	Response := variables.Response{Action: variables.GetMovies}
	opts := options.Find().SetProjection(bson.M{"_id": 1, "image_url": 1, "title": 1, "type": 1,})

	cursor, err := collection.Find(ctx, bson.M{"casts": name}, opts)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	err = cursor.All(ctx, &Movies)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}

	if start > len(Movies) {
		Response.Data = []Movie{}
	}else if end > len(Movies) {
		Response.Data = Movies[start:]
	}else {
		Response.Data = Movies[start:end]
	}

	Response.Success = true
	return variables.JsonMarshal(Response), http.StatusOK
}