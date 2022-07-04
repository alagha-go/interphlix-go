package projects

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func (project *Project) Remove() ([]byte, int) {
	var Response variables.Response
	Response.Action = variables.Delete
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Projects")

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
	Response.Success = true
	Response.Data = "project deleted"
	return variables.JsonMarshal(Response), http.StatusOK
}