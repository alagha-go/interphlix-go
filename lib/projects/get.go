package projects

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func GetMyProjects(accountID primitive.ObjectID) ([]byte, int) {
	var projects []Project
	Response := variables.Response{Action: variables.GetProjects}
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Projects")

	opts := options.Find().SetProjection(bson.D{{"api_keys", 0}})

	cursor, err := collection.Find(ctx, bson.M{"account_id": accountID}, opts)
	if err != nil {
		variables.SaveError(err, "projects", "GetMyProjects")
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	err = cursor.All(ctx, &projects)
	if err != nil {
		variables.SaveError(err, "projects", "GetMyProjects")
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}

	Response.Success = true
	Response.Data = projects
	return variables.JsonMarshal(Response), http.StatusOK
}


func (project *Project) GetApiKeys() ([]byte, int) {
	var Project Project
	Response := variables.Response{Action: variables.GetProjects}
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Projects")

	opts := options.FindOne().SetProjection(bson.D{{"api_keys", 1}})

	err := collection.FindOne(ctx, bson.M{"_id": project.ID}, opts).Decode(&Project)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.ProjectNotFound
		return variables.JsonMarshal(Response), http.StatusNotFound
	}
	Response.Success = true
	Response.Data = Project.ApiKeys
	return variables.JsonMarshal(Response), http.StatusOK
}