package projects

import (
	"interphlix/lib/projects"
	"interphlix/lib/variables"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func DeleteProject(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.Delete}
	projectID, err := primitive.ObjectIDFromHex(mux.Vars(req)["projectId"])
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidID
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	project := projects.Project{ID: projectID}
	data, status := project.Delete()
	res.WriteHeader(status)
	res.Write(data)
}


func DeleteApiKey(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.Delete}
	projectID, err := primitive.ObjectIDFromHex(mux.Vars(req)["projectId"])
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidID
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	name := mux.Vars(req)["name"]
	data, status := projects.DeleteApiKey(projectID, name)
	res.WriteHeader(status)
	res.Write(data)
}