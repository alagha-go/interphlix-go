package projects

import (
	"encoding/json"
	"interphlix/lib/projects"
	"interphlix/lib/requests"
	"interphlix/lib/variables"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CreateProject(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var project projects.Project
	Response := variables.Response{Action: variables.CreateProject}
	err := json.NewDecoder(req.Body).Decode(&project)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidJson
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	account, err := requests.GetmyAccount(req)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.UserNotFound
		res.WriteHeader(http.StatusNotFound)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	if !account.EmailVerified {
		Response.Failed = true
		Response.Error = variables.EmailNotVerified
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	project.AccountID = account.ID
	data, status := project.CreateProject()
	res.WriteHeader(status)
	res.Write(data)
}


func GenerateApiKey(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.GenerateApiKey}
	projectID, err := primitive.ObjectIDFromHex(mux.Vars(req)["projectId"])
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidID
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	project := projects.Project{ID: projectID}
	name := req.URL.Query().Get("name")
	data, status := project.GenerateKey(name)
	res.WriteHeader(status)
	res.Write(data)
}