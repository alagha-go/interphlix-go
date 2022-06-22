package accounts

import (
	"interphlix/lib/variables"
	"net/http"
)


func GetMyAccount(res http.ResponseWriter, req *http.Request) {
	var Response variables.Response
	Response.Action = variables.GetAccount
	res.Header().Set("content-type", "application/json")
	err := ValidateRequest(req)
	if err != nil {
		Response.Failed = true
		Response.Error = err.Error()
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	account, err := GetmyAccount(req)
	if err != nil {
		Response.Failed = true
		Response.Action = variables.GetAccount
		Response.Error = err.Error()
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	Response.Success = true
	Response.Data = account
	res.WriteHeader(200)
	res.Write(variables.JsonMarshal(Response))
}