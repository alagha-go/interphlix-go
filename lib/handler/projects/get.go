package projects

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/projects"
	"interphlix/lib/variables"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetMyProjects(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.GetProjects}
	account, err := accounts.GetmyAccount(req)
	if err != nil {
		Response.Failed = true
		Response.Error = err.Error()
		res.WriteHeader(http.StatusNotFound)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	data, status := projects.GetMyProjects(account.ID)
	res.WriteHeader(status)
	res.Write(data)
}


func GetProjectApiKeys(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.GetApiKeys}
	projectID, err := primitive.ObjectIDFromHex(mux.Vars(req)["projectId"])
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidID
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	project := projects.Project{ID: projectID}
	data, status := project.GetApiKeys()
	res.WriteHeader(status)
	res.Write(data)
}