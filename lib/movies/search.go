package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Search(round int, querry string, Type ...string) ([]byte, int) {
	var Movies []Movie
	Response := variables.Response{Action: variables.SearchMovie}
	filter := bson.M{"$text": bson.M{"$search": querry}}
	start := 0
	if round > 0 {
		start = (round*MoviesLimit) - MoviesLimit
	}
	end := start+MoviesLimit

	if len(Type) > 0 {
		filter = bson.M{"type": Type[0], "$text": bson.M{"$search": querry}}
	}

	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Movies")

	sort := bson.M{"score": bson.M{"$meta": "textScore"}}
	projection := bson.M{"score": bson.M{"$meta": "textScore"}, "_id": 1, "image_url": 1, "type": 1, "title": 1}
	opts := options.Find().SetSort(sort).SetProjection(projection)

	cursor, err := collection.Find(ctx, filter, opts)
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

	return variables.JsonMarshal(Response), http.StatusOK
}