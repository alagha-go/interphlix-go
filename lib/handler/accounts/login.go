package accounts

import (
	"encoding/json"
	"interphlix/lib/accounts"
	"interphlix/lib/variables"
	"net/http"
)

//// user logged in with google an then redirected
func Redirect(res http.ResponseWriter, req *http.Request) {
	var account accounts.Account
	code := req.URL.Query().Get("code")
	if code == "" {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Response{Action: variables.CreateUserAction, Error: variables.NoCode, Failed: true}))
		return
	}
	token, err := GetToken(code)
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Response{Action: variables.CreateUserAction, Error: variables.CouldNotGetToken, Failed: true}))
		return
	}
	account, err = GetUserInfo(token.AccessToken)
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Response{Action: variables.CreateUserAction, Error: variables.CouldNotGetUserInfoFromGoogle, Failed: true}))
		return
	}
	account.SetToken(token)
	account.SignUpMethod = "google"
	data, status := accounts.CreateAccount(&account)
	json.Unmarshal(data, &account)
	cookie, status1, err := GenerateToken(account)
	if err != nil {
		res.WriteHeader(status1)
		res.Write(variables.JsonMarshal(variables.Response{Action: variables.CreateUserAction}))
	}
	http.SetCookie(res, cookie)
	res.WriteHeader(status)
	res.Write(data)
}


func Login(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var account accounts.Account
	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		response := variables.Response{Action: variables.CreateUserAction, Failed: true, Error: variables.InvalidJson}
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(response))
		return
	}
	data, status := account.Login()
	cookie, status1, err := GenerateToken(account)
	if err != nil {
		res.WriteHeader(status1)
		res.Write(variables.JsonMarshal(variables.Response{Action: variables.Login}))
	}
	http.SetCookie(res, cookie)
	res.WriteHeader(status)
	res.Write(data)
}