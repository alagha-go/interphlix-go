package projects

import (
	"encoding/json"
	"interphlix/lib/handler/accounts"
	"interphlix/lib/projects"
	"interphlix/lib/variables"
	"net/http"
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
	account, err := accounts.GetmyAccount(req)
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