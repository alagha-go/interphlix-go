package casts

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func Search(round int, query string) ([]byte, int) {
	var Casts []Cast
	allowDiskUse := true
	start := 0
	if round > 0 {
		start = (round*CastsLimit) - CastsLimit
	}
	end := start+CastsLimit

	Response := variables.Response{Action: variables.SearchCast}
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Casts")

	projection := bson.M{"score": bson.M{"$meta": "textScore"}, "_id": 1, "name": 1, "profile_path": 1}
	sort := bson.M{"score": bson.M{"$meta": "textScore"}}
	opts := options.Find().SetProjection(projection).SetSort(sort)
	opts.AllowDiskUse = &allowDiskUse

	cursor, err := collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}}, opts)
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

	if start > len(Casts) {
		Response.Data = []Cast{}
	}else if end > len(Casts) {
		Response.Data = Casts[start:]
	}else {
		Response.Data = Casts[start:end]
	}

	return variables.JsonMarshal(Response), http.StatusOK
}