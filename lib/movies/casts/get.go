package casts

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	CastsLimit = 30
)

// get casts
func GetCasts(round int) ([]byte, int) {
	start := 0
	if round > 0 {
		start = (round*CastsLimit) - CastsLimit
	}
	end := start+CastsLimit
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

	if start > len(Casts) {
		Response.Data = []Cast{}
	}else if end > len(Casts) {
		Response.Data = Casts[start:]
	}else {
		Response.Data = Casts[start:end]
	}

	Response.Success = true
	return variables.JsonMarshal(Response), http.StatusOK
}


// get cast by either providing name or ID
func GetCast(name string, ID *primitive.ObjectID) ([]byte, int) {
	Response := variables.Response{Action: variables.GetCast}
	if ID != nil {
		Cast := LoadCastByID(*ID)
		if Cast.KnownForDepartment == "" {
			Cast := GetCastInfo(Cast.Name)
			Cast.ID = ID
			Cast.Update()
		}
		Response.Success = true
		Response.Data = Cast
		return variables.JsonMarshal(Response), http.StatusOK
	}
	Cast := LoadCastByName(name)
	ID = Cast.ID
	if Cast.KnownForDepartment == "" {
		Cast := GetCastInfo(Cast.Name)
		Cast.ID = ID
		Cast.Update()
	}
	Response.Success = true
	Response.Data = Cast
	return variables.JsonMarshal(Response), http.StatusOK
}