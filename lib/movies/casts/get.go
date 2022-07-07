package casts

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// get casts
func GetCasts(round int) ([]byte, int) {
	Response := variables.Response{Action: variables.GetCasts}
	var Casts []Cast
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Casts")

	opts := options.Find().SetProjection(bson.M{"_id": 1, "name": 1})

	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}

	err = cursor.All(ctx, &Casts)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	Response.Success = true
	Response.Data = Casts
	return variables.JsonMarshal(Response), http.StatusOK
}