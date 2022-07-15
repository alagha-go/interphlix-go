package projects

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	RequestsLimit = 50
)

func (project *Project) CreateProject() ([]byte, int) {
	project.RequestsLimit = int64(RequestsLimit)
	var Response variables.Response
	Response.Action = variables.CreateProject
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Projects")

	canCreate, err := CanCreateNewProject(project.AccountID)
	if !canCreate {
		Response.Failed = true
		Response.Error = err.Error()
		return variables.JsonMarshal(Response), http.StatusNotAcceptable
	}

	if project.Name == "" {
		Response.Failed = true
		Response.Error = variables.InvalidName
		return variables.JsonMarshal(Response), http.StatusBadRequest
	}

	if project.Exists() {
		Response.Failed = true
		Response.Error = variables.ProjectExists
		return variables.JsonMarshal(Response), http.StatusConflict
	}

	project.ID = CreateProjectID()
	_, err = collection.InsertOne(ctx, project)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	Response.Success = true
	Response.Data = project
	return variables.JsonMarshal(Response), http.StatusOK
}


func (project *Project) Exists() bool {
	var Project Project
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Projects")

	err := collection.FindOne(ctx, bson.M{"name": project.Name, "account_id": project.AccountID}).Decode(&Project)
	return err == nil
}

func CreateProjectID() primitive.ObjectID {
	var Project Project
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Projects")
	ID := primitive.NewObjectID()
	err := collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&Project)
	if err == nil {
		return CreateProjectID()
	}
	return ID
}