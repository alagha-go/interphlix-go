package accounts

import (
	"encoding/json"
	"interphlix/lib/accounts"
	"interphlix/lib/requests"
	"interphlix/lib/variables"
	"net/http"
)


func SignUp(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var account accounts.Account
	Response := variables.Response{Action: variables.CreateUserAction}
	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InvalidJson
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	if len(account.Password) < 4 {
		Response.Failed = true
		Response.Error = variables.ShortPassword
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(Response))
		return
	}
	account.EmailVerified = false
	data, status := accounts.CreateAccount(&account)
	if status == 201 {
		cookie, _, _ := requests.GenerateToken(account)
		http.SetCookie(res, cookie)
	}
	res.WriteHeader(status)
	res.Write(data)
}

func GoogleSignUp(res http.ResponseWriter, req *http.Request) {
	url, err := GetUrl()
	if err != nil {
		res.Header().Set("content-type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(variables.JsonMarshal(variables.Response{Action: variables.CreateUserAction, Failed: true, Error: variables.InternalServerError}))
		return
	}
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}