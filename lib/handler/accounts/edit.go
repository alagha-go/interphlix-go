package accounts

import (
	"encoding/json"
	"interphlix/lib/accounts"
	"interphlix/lib/variables"
	"net/http"
)


func ChangePassword(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var account accounts.Account
	Response := variables.Response{Action: variables.ChangePassword}
	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidJson
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	if account.Email == "" {
		dbaccount, err := GetmyAccount(req)
		if err != nil {
			Response.Failed = true
			Response.Error = variables.UserNotFound
			res.WriteHeader(http.StatusBadRequest)
			res.Write(variables.JsonMarshal(Response))
			return
		}
		account.Email = dbaccount.Email
	}
	if len(account.NewPassword) < 4 {
		Response.Failed = true
		Response.Error = variables.ShortPassword
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	data, status := accounts.ChangePassword(account.Email, account.Password, account.NewPassword)
	res.WriteHeader(status)
	res.Write(data)
}


func UpdateAccount(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var account accounts.Account
	Response := variables.Response{Action: variables.UpdateAccount}
	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidJson
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	Account, err := GetmyAccount(req)
	if err != nil {
		Response.Failed = true
		Response.Error = err.Error()
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	account.ID = Account.ID
	data, status := account.Update()
	res.WriteHeader(status)
	res.Write(data)
}