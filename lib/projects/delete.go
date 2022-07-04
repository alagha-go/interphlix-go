package projects

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func (project *Project) Delete() ([]byte, int) {
	var Response variables.Response
	Response.Action = variables.Delete
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Projects")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": project.ID})
	if err != nil {
		Response.Failed = true
		if err == mongo.ErrNilDocument {
			Response.Error = variables.ProjectNotFound
			return variables.JsonMarshal(Response), http.StatusNotFound
		}
		variables.SaveError(err, "projects", "project.Remove")
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	return project.Remove()
}



func DeleteApiKey(ProjectID primitive.ObjectID, name string) ([]byte, int) {
	var Response variables.Response
	Response.Action = variables.Delete
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Projects")
	filter := bson.M{"_id": bson.M{"$eq": ProjectID}}
	update := bson.M{"$pull": bson.M{"api_keys": bson.D{{"name", name}}}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		Response.Failed = true
		if err == mongo.ErrNilDocument {
			Response.Error = variables.ProjectNotFound
			return variables.JsonMarshal(Response), http.StatusNotFound
		}
		variables.SaveError(err, "projects", "DeleteApiKey")
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	Response.Success = true
	Response.Data = "api key deleted"
	return variables.JsonMarshal(Response), http.StatusOK
}