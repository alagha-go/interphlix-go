package projects

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (project *Project) GenerateKey(name string) ([]byte, int) {
	var Project Project
	var Response variables.Response
	Key := primitive.NewObjectID().Hex()
	Response.Action = variables.GenerateApiKey
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Projects")

	if name == "" {
		Response.Failed = true
		Response.Error = variables.InvalidName
		return variables.JsonMarshal(Response), http.StatusBadRequest
	}

	err := collection.FindOne(ctx, bson.M{"_id": project.ID}).Decode(&Project)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.ProjectNotFound
		return variables.JsonMarshal(Response), http.StatusNotFound
	}

	for index := range Project.ApiKeys {
		if Project.ApiKeys[index].Name == name {
			Response.Failed = true
			Response.Error = variables.ApiKeyExists
			return variables.JsonMarshal(Response), http.StatusConflict
		}
	}

	apiKey := ApiKey{Name: name, Key: Hasher([]byte(Key))}

	filter := bson.M{"_id": bson.M{"$eq": project.ID}}
	update := bson.M{"$addToSet": bson.M{"api_keys": apiKey}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	Response.Success = true
	Response.Data = ApiKey{Name: name, Key: Key}
	return variables.JsonMarshal(Response), http.StatusCreated
}