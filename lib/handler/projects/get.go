package projects

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/projects"
	"interphlix/lib/variables"
	"net/http"
)


func GetMyProjects(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	Response := variables.Response{Action: variables.GenerateApiKey}
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